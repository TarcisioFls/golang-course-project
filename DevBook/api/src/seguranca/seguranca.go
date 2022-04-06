package seguranca

import "golang.org/x/crypto/bcrypt"

//Hash recebe uma string e retorna um hash
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

//VerificarSenha compara o valor da senhaString com o a senhaHash
func VerificarSenha(senhaHash string, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senhaString))
}
