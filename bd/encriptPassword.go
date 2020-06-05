package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptPassword methods from password users*/
func EncriptPassword(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}
