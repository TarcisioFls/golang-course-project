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
