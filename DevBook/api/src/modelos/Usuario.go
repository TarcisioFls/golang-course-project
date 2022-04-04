package modelos

import (
	"errors"
	"strings"
	"time"
)

//Usuário modelo da tabela no banco de dados
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

//Preparar vai chamar os métodos para validar e formatar o usuário recebido
func (usuario *Usuario) Preparar() error {
	if erro := usuario.validar(); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar() error {

	if usuario.Nome == "" {
		return errors.New("O campo nome é obrigatório e não pode ser branco")
	}

	if usuario.Nick == "" {
		return errors.New("O campo nick é obrigatório e não pode ser branco")
	}

	if usuario.Email == "" {
		return errors.New("O campo email é obrigatório e não pode ser branco")
	}

	if usuario.Senha == "" {
		return errors.New("O campo senha é obrigatório e não pode ser branco")
	}

	return nil

}

func (usuario *Usuario) formatar() {

	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

}
