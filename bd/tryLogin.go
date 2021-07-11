package bd

import (
	"github.com/arivero007/Go-BackEnd-Tutorial/models"
	"golang.org/x/crypto/bcrypt"
)

func TryLogin(email string, pass string) (models.User, bool) {
	u, found, _ := CheckUserExist(email)
	if !found {
		return u, false
	}
	passwordData := []byte(pass)
	passwordDB := []byte(u.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordData)
	if err != nil {
		return u, false
	}
	return u, true
}
