package main

import (
	"encoding/json"
	"io"
	"log"
	"strings"
)

func main() {
	const jsonStream = `
    {"ProductoID": "AXW123", "Precio": 25.50}
    {"ProductoID": "NLBR17", "Precio": 357.58}
    {"ProductoID": "KNUB82", "Precio": 150}
    `
	// Es necesario crear nuestro type Decode, para esto se llama la funcion NewDecoder
	// esta recibe por parámetro un streaming
	// Se crea una variable jsonStream y se usa el método NewReader del pkg strings
	// NewReader genera un streaming para la cadena de texto que recibe.
	myStreaming := strings.NewReader(jsonStream)
	myDecoder := json.NewDecoder(myStreaming)
	type MyData struct {
		ProductoID string
		Precio     float64
	}
	// el streaming se comporta de forma que cada línea en el texto jsonStrem es transmitida por separado
	// recorremos toda la data transmitida en el streaming hasta que se alcanza el final del texto
	for {
		// se crea la variable sobre la que se va a escribir los datos
		var data MyData
		// se invoca el método decode
		// Decode se encarga de leer la data transmitida por el streaming y transformarla de json a nuestra interfaz
		if err := myDecoder.Decode(&data); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		// Se imprimen los datos recibidos
		log.Println("datos recibidos:", data.ProductoID, data.Precio)
	}

}
