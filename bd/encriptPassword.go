package bd

import (
	"golang.org/x/crypto/bcrypt"
)

func EncriptPassword(pass string) (string, error) {
	cost := 6
	data, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(data), err
}
