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
		INSERT INTO publicacao (titulo, conteudo, autor_id) 
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
		FROM publicacao p INNER JOIN usuarios u ON p.autor_id = u.id
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
