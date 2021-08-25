package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/arivero007/Go-BackEnd-Tutorial/bd"
	"github.com/arivero007/Go-BackEnd-Tutorial/models"
)

func RecordTweet(w http.ResponseWriter, r *http.Request) {
	var msg models.Tweet

	err := json.NewDecoder(r.Body).Decode(&msg)

	if err != nil {
		http.Error(w, "Error al decodificar body"+err.Error(), 400)
		return
	}

	register := models.RecTweet{
		UserID:  IDUser,
		Message: msg.Message,
		Date:    time.Now(),
	}

	_, status, err := bd.InsertTweet(register)

	if err != nil {
		http.Error(w, "Error al insertar tweet"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
