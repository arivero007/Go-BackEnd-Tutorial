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
	router.HandleFunc("/updateProfile", middlew.ChequeoBD(middlew.ValidateJWT(routers.UpdateProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidateJWT(routers.RecordTweet))).Methods("POST")
	router.HandleFunc("/getTweets", middlew.ChequeoBD(middlew.ValidateJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/deleteTweet", middlew.ChequeoBD(middlew.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/setAvatar", middlew.ChequeoBD(middlew.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/getAvatar", middlew.ChequeoBD(middlew.ValidateJWT(routers.GetAvatar))).Methods("GET")
	router.HandleFunc("/setBanner", middlew.ChequeoBD(middlew.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getBanner", middlew.ChequeoBD(middlew.ValidateJWT(routers.GetBanner))).Methods("GET")

	router.HandleFunc("/createRelation", middlew.ChequeoBD(middlew.ValidateJWT(routers.CreateRelation))).Methods("POST")
	router.HandleFunc("/deleteRelation", middlew.ChequeoBD(middlew.ValidateJWT(routers.RemoveRelation))).Methods("DELETE")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
