package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/arivero007/Go-BackEnd-Tutorial/bd"
)

func GetBanner(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	profile, err := bd.GetUser(ID)

	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	openFile, err := os.Open("uploads/banners/" + profile.Avatar)

	if err != nil {
		http.Error(w, "Banner no encontrado", http.StatusNotFound)
		return
	}

	_, err = io.Copy(w, openFile)

	if err != nil {
		http.Error(w, "Banner no se ha podido copiar", http.StatusConflict)
		return
	}

}
