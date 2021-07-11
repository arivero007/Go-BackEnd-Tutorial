package routers

import (
	"errors"
	"strings"

	"github.com/arivero007/Go-BackEnd-Tutorial/bd"
	"github.com/arivero007/Go-BackEnd-Tutorial/models"
	"github.com/golang-jwt/jwt"
)

var Email string
var IDUser string

func TokenProcess(tk string) (*models.Claims, bool, string, error) {
	myKey := []byte("gotutorial")
	claims := &models.Claims{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, "", errors.New("formato de token invalido")
	}
	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err == nil {
		_, found, _ := bd.CheckUserExist(claims.Email)
		if found {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, found, IDUser, nil
	}

	if !tkn.Valid {
		return claims, false, "", errors.New("token invalido")
	}
	return claims, false, "", err
}
