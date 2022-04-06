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
func (repositorio usuarios) UpdateUsuario(usuarioID int64, usuario modelos.Usuario) error {

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
func (repositorio usuarios) DeleteUsuario(usuarioID int64) error {

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
