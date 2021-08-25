package routers

import (
	"net/http"

	"github.com/arivero007/Go-BackEnd-Tutorial/bd"
	"github.com/arivero007/Go-BackEnd-Tutorial/models"
)

func RemoveRelation(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID

	status, err := bd.DeleteRelation(t)

	if err != nil {
		http.Error(w, "Error al eliminar relación"+err.Error(), http.StatusConflict)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado eliminar relación"+err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
