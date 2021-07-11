package jwtoken

import (
	"time"

	"github.com/arivero007/Go-BackEnd-Tutorial/models"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(t models.User) (string, error) {
	myKey := []byte("gotutorial")

	payload := jwt.MapClaims{
		"email":     t.Email,
		"name":      t.Name,
		"firstName": t.FirstName,
		"bornDate":  t.BornDate,
		"biography": t.Biography,
		"location":  t.Location,
		"website":   t.Website,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
