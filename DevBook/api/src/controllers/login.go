package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

	corpoRequisição, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)

		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequisição, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)

		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)

		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuario(db)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)

		return
	}

	usuarioSalvoNoBanco, erro := repositorio.Login(usuario.Email)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)

		return
	}

	if erro = seguranca.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)

		return
	}

	respostas.JSON(w, http.StatusAccepted, "Logado")
}
