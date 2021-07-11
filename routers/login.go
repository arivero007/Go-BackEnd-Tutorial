package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/arivero007/Go-BackEnd-Tutorial/bd"
	"github.com/arivero007/Go-BackEnd-Tutorial/jwtoken"
	"github.com/arivero007/Go-BackEnd-Tutorial/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o contraseña invalidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Contraseña demasido corta, debe tener más de 6 caracteres", 400)
		return
	}

	u, exist := bd.TryLogin(t.Email, t.Password)

	if !exist {
		http.Error(w, "El usuario no existe", 400)
	}

	jwtKey, err := jwtoken.GenerateJWT(u)
	if err != nil {
		http.Error(w, "Error al generar el token"+err.Error(), 400)
		return
	}

	resp := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
