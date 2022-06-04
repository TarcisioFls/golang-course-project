package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type usuarios struct {
	db *sql.DB
}

//NovoRepositorioDeUsuario cria um repositório de usuários
func NovoRepositorioDeUsuario(db *sql.DB) *usuarios {
	return &usuarios{db}
}

//Criar Insere um usuário no banco de dados
func (repositorio usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(`
	INSERT INTO  usuarios (nome, nick, email, senha)
	VALUES (?, ?, ?, ?)
	`)

	if erro != nil {

		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {

		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {

		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

//Buscar traz todos os usuários que atendem ao filtro de nome ou nick
func (repositorio usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) //%nomeOuNick%

	linhas, erro := repositorio.db.Query(
		`SELECT id, nome, nick, email, criadoEm 
	FROM usuarios 	
	WHERE nome LIKE ? OR nick LIKE ?`,
		nomeOuNick, nomeOuNick,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {

		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

//GetUsuario método para retornar o usuário que possuí o id passado como paramentro
func (repositorio usuarios) GetUsuario(usuarioID int64) (modelos.Usuario, error) {
	var usuario modelos.Usuario
	linha, erro := repositorio.db.Query(`
		SELECT id, nome, nick, email, criadoEm
		FROM usuarios
		WHERE id = ?
	`, usuarioID)
	if erro != nil {
		return usuario, erro
	}

	defer linha.Close()

	if linha.Next() {
		if erro = linha.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return usuario, erro
		}
	}

	return usuario, nil
}

//UpdateUsuario atualiza os dados do usuário informando
func (repositorio usuarios) UpdateUsuario(usuarioID uint64, usuario modelos.Usuario) error {

	statement, erro := repositorio.db.Prepare(`
		UPDATE usuarios set nome = ?, nick = ?, email = ? 
		WHERE id = ?
	`)
	if erro != nil {

		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuarioID); erro != nil {

		return erro
	}

	return nil

}

//DeletarUsuario deleta as informações do usuário com o id informado
func (repositorio usuarios) DeleteUsuario(usuarioID uint64) error {

	statement, erro := repositorio.db.Prepare(`
		DELETE FROM usuarios WHERE id = ?
	`)
	if erro != nil {

		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(usuarioID); erro != nil {
		return erro
	}

	return nil

}

func (repositorio usuarios) Login(email string) (modelos.Usuario, error) {
	var usuario modelos.Usuario
	linha, erro := repositorio.db.Query(`
		SELECT id, senha
		FROM usuarios
		WHERE email = ?
	`, email)
	if erro != nil {
		return usuario, erro
	}

	defer linha.Close()

	if linha.Next() {
		if erro = linha.Scan(
			&usuario.ID,
			&usuario.Senha,
		); erro != nil {
			return usuario, erro
		}
	}

	return usuario, nil
}

//Seguir permite que um usúario siga o outro
func (repositorio usuarios) Seguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"INSERT IGNORE INTO seguidores (usuario_id, seguidor_id) VALUES (?,?)",
	)

	if erro != nil {

		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {

		return erro
	}

	return nil
}

func (repositorio usuarios) PararDeSeguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(`
		DELETE FROM seguidores WHERE usuario_id = ? AND seguidor_id = ?
	`)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

//BuscarSeguidores permite retornar os usuários que seguem o usuário pasado por parametro
func (repositorio usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
		SELECT u.id, u.nome, u.nick, u.email, u.criadoEm
		FROM usuarios u
		INNER JOIN seguidores s ON u.id = s.seguidor_Id
		WHERE s.usuario_id = ?
	`, usuarioID)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

//BuscarSeguindo permite retornar os seguidores do usuário selecionado
func (repositorio usuarios) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
		SELECT u.id, u.nome, u.nick, u.email, u.criadoEm
		FROM usuarios u
		INNER JOIN seguidores s ON u.id = s.usuario_Id
		WHERE s.seguidor_id = ?
	`, usuarioID)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

//BuscarSenha a senha do usuário informado
func (repositorio usuarios) BuscarSenha(usuarioId uint64) (string, error) {
	linha, erro := repositorio.db.Query(`
		SELECT senha 
		FROM usuarios 
		WHERE id = ?
	`, usuarioId)

	if erro != nil {
		return "", erro
	}

	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil
}

//AtualizarSenha altera a senha do usuário informado
func (repositorio usuarios) AtualizarSenha(usuarioId uint64, senha string) error {
	statement, erro := repositorio.db.Prepare(`
		UPDATE usuarios 
		SET senha = ?
		WHERE id = ?
	`)

	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(senha, usuarioId); erro != nil {
		return erro
	}

	return nil
}
