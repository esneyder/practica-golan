package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/esneyder/practica-golan/middlew"
	"github.com/esneyder/practica-golan/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Managers setear handlre y escucha del servidor*/
func Managers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckBD(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
