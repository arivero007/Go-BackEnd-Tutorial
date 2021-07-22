package routers

import (
	"encoding/json"
	"net/http"

	"github.com/arivero007/Go-BackEnd-Tutorial/bd"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Error al enviar el ID", http.StatusBadRequest)
		return
	}
	profile, err := bd.GetUser(ID)
	if err != nil {
		http.Error(w, "Ocurrio un error al buscar el registro"+err.Error(), 400)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}
