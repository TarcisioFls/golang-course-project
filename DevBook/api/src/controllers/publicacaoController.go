package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//CriarPublicacao é responsável por salvar uma publicação no banco de dados
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {

	usuarioId, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)

		return
	}

	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)

		return
	}

	var publicacao modelos.Publicacao
	if erro = json.Unmarshal(corpoRequest, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)

		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)

		return
	}

	defer db.Close()

	publicacao.AutorId = usuarioId

	if erro = publicacao.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)

		return
	}

	repositorio := repositorios.NovoRepositorioDePublicacao(db)
	publicacao.Id, erro = repositorio.Criar(publicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)

		return
	}

	respostas.JSON(w, http.StatusCreated, publicacao)

}

//BuscarPublicacoes é resposável por retornar todas as publicações de um determinado usuário
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {

}

//BuscarPublicacao é resposável por retornar uma publicação em particular
func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)

		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)

		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacao(db)
	publicacao, erro := repositorio.BuscarPorId(publicacaoId)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)

		return
	}

	respostas.JSON(w, http.StatusOK, publicacao)

}

//AtualizarPublicacao é responsável por atualizar uma publicação em particular
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {

}

//DeletarPublicacao é responsável por deletar uma publicação em particular
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {

}
