package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

//CriarUsuario Criando o usuário
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	coporRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)

		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(coporRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)

		return
	}

	if erro = usuario.Preparar("cadastro"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)

		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)

		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuario(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)

		return
	}

	respostas.JSON(w, http.StatusCreated, usuario)
}

//BuscarUsuario retornar todos os usuários
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)

		return
	}
	defer db.Close()

	respositorio := repositorios.NovoRepositorioDeUsuario(db)
	usuarios, erro := respositorio.Buscar(nomeOuNick)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}

//BuscarUsuário retonar um usuário
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

	usuarioID, erro := strconv.ParseInt(mux.Vars(r)["usuarioId"], 10, 64)
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

	repositorio := repositorios.NovoRepositorioDeUsuario(db)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)

		return
	}

	usuario, erro := repositorio.GetUsuario(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
	}
	respostas.JSON(w, http.StatusOK, usuario)
}

//AtualizarUsuario Atualiza um usuário
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {

	usuarioID, erro := strconv.ParseUint(mux.Vars(r)["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)

		return
	}

	usuarioIDNoToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioID != usuarioIDNoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar um usuário diferente do seu"))
		return
	}

	coporRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)

		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(coporRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)

		return
	}

	if erro = usuario.Preparar("edicao"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)

		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)

		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuario(db)
	if erro = repositorio.UpdateUsuario(usuarioID, usuario); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)

		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

//DeletarUsuário deletando um usuário
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {

	usuarioID, erro := strconv.ParseUint(mux.Vars(r)["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)

		return
	}

	usuarioIDNoToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioID != usuarioIDNoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível deletar um usuário diferente do seu"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)

		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuario(db)
	if erro = repositorio.DeleteUsuario(usuarioID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)

		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}
