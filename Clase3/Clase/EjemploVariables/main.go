package main

import (
	"log"

	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {

		log.Fatal("Error al intentar cargar archivo .env")

	}

	usuario := os.Getenv("MY_USER")

	password := os.Getenv("MY_PASS")

	println("\nUsuario sacado de variables de Entorno: ", usuario)

	println("\nPassword sacado de variables de Entorno: ", password)

}
