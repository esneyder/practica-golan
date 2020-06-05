package main

import (
	"log"

	"github.com/esneyder/practica-golan/bd"
	"github.com/esneyder/practica-golan/handlers"
)

func main() {
	if bd.CheckConnect() == 0 {
		log.Fatal("Sin conexion")
	}
	handlers.Managers()
}
