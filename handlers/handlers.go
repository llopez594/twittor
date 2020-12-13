package handlers

import (
	"github.com/llopez594/twittor/routers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/llopez594/twittor/middlew"
	"github.com/rs/cors"
)

/*Manejadores setio mi puerto, el handler y pongo a escuchar mi Servidor*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")



	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))


}