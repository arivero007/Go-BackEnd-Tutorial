package routers

import (
	"io"
	"net/http"
	"os"

	"strings"

	"github.com/arivero007/Go-BackEnd-Tutorial/bd"
	"github.com/arivero007/Go-BackEnd-Tutorial/models"
)

func UploadAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar")

	if err != nil {
		http.Error(w, "Falta el formulario de avatar "+err.Error(), http.StatusBadRequest)
		return
	}

	var extension = strings.Split(handler.Filename, ".")[1]
	var identifier string = IDUser + "." + extension
	var savePath string = "uploads/avatars/" + identifier

	f, err := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "Error al subir la imagen"+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "Error al guardar la imagen"+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Avatar = identifier
	status, err = bd.UpdateUser(user, IDUser)

	if err != nil || !status {
		http.Error(w, "Error al guardar avatar en bd"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
