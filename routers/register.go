package routers

import (
	"encoding/json"
	"net/http"

	"github.com/arivero007/Go-BackEnd-Tutorial/bd"
	"github.com/arivero007/Go-BackEnd-Tutorial/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
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
	_, encontrado, _ := bd.CheckUserExist(t.Email)
	if encontrado {
		http.Error(w, "Usuario ya existe", 400)
		return
	}
	_, status, err := bd.InsertUser(t)
	if err != nil {
		http.Error(w, "Ocurrio un error en el registro"+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se a logrado finalizar el registro"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
