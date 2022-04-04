package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
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

	if erro = usuario.Preparar(); erro != nil {
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
	w.Write([]byte("Buscando um usuário"))
}

//AtualizarUsuario Atualiza um usuário
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário"))
}

//DeletarUsuário deletando um usuário
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário"))
}
