package modelos

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar(etapa string) error {

	if usuario.Nome == "" {
		return errors.New("O campo nome é obrigatório e não pode ser branco")
	}

	if usuario.Nick == "" {
		return errors.New("O campo nick é obrigatório e não pode ser branco")
	}

	if usuario.Email == "" {
		return errors.New("O campo email é obrigatório e não pode ser branco")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O e-mail inserido é inválido")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("O campo senha é obrigatório e não pode ser branco")
	}

	return nil

}

func (usuario *Usuario) formatar() {

	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

}
