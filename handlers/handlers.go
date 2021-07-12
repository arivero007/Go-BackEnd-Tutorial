package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/arivero007/Go-BackEnd-Tutorial/middlew"
	"github.com/arivero007/Go-BackEnd-Tutorial/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.ChequeoBD(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/profile", middlew.ChequeoBD(middlew.ValidateJWT(routers.Profile))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
