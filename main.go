package main

import (
	"log"

	"github.com/llopez594/twittor/bd"
	"github.com/llopez594/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection()==0 {
		log.Fatal("sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
