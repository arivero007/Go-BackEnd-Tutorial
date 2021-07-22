package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/arivero007/Go-BackEnd-Tutorial/bd"
)

func ReadTweets(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parámetro page", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro page con un valor mayor a 0", http.StatusBadRequest)
		return
	}

	pageValue := int64(page)
	result, status := bd.ReadTweets(ID, pageValue)
	if !status {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
