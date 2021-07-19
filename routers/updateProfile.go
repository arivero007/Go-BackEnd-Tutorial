package routers

import (
	"encoding/json"
	"net/http"

	"github.com/arivero007/Go-BackEnd-Tutorial/bd"
	"github.com/arivero007/Go-BackEnd-Tutorial/models"
)

func UpdateProfile(w http.ResponseWriter, r *http.Request) {

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Incorrect data"+err.Error(), 400)
		return
	}

	var status bool
	status, err = bd.UpdateUser(t, IDUser)

	if err != nil {
		http.Error(w, "Error al intentar modificar usuario"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado modificar el registro de ususario", 400)
		return
	}

	w.WriteHeader(http.StatusAccepted)

}
