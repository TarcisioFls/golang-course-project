package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

//Publicacoes representa um respositorio de publicações
type publicacoes struct {
	db *sql.DB
}

//NovoRepositorioDePublicacao cria um repositório de publicações
func NovoRepositorioDePublicacao(db *sql.DB) *publicacoes {

	return &publicacoes{db}
}

//Criar insere uma publicação no banco de dados
func (repositorio publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statament, erro := repositorio.db.Prepare(`
		INSERT INTO publicacoes (titulo, conteudo, autor_id) 
		VALUES(?,?,?)
	`)

	if erro != nil {
		return 0, erro
	}

	defer statament.Close()

	resultado, erro := statament.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorId)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

//BuscarPorId retornar uma única publicação do banco de dados
func (repositorio publicacoes) BuscarPorId(publicacaoId uint64) (modelos.Publicacao, error) {
	var publicacao modelos.Publicacao
	linha, erro := repositorio.db.Query(`
		SELECT p.*, u.nick
		FROM publicacoes p INNER JOIN usuarios u ON p.autor_id = u.id
		WHERE p.id = ?
	`, publicacaoId)
	if erro != nil {
		return publicacao, erro
	}

	defer linha.Close()

	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.Id,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorId,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return publicacao, erro
		}
	}

	return publicacao, nil
}

//Buscar retornar todas as públicações do usuário e dos seus seguidores
func (repositorio publicacoes) Buscar(usuarioId uint64) ([]modelos.Publicacao, error) {

	linhas, erro := repositorio.db.Query(`
		SELECT DISTINCT p.*, u.nick
		FROM publicacoes p 
		INNER JOIN usuarios u ON p.autor_id = u.id
		LEFT JOIN seguidores s on p.autor_id = s.usuario_id
		WHERE p.autor_id = ? OR s.seguidor_id = ?
		ORDER BY 1 DESC
	`, usuarioId, usuarioId)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var publicacoes []modelos.Publicacao
	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(
			&publicacao.Id,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorId,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}
