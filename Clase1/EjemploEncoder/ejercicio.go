package main

import (
	"encoding/json"
	"os"
)

// Es necesario crear nuestro type Encode
// para esto se llama la funcion NewEncoder
// esta recibe por parámetro un streaming
// hacemos uso de uno de los stream standard ofrecidos por el pkg os Stdout
// stdout genera un stream a un archivo que se imprime en consola.

// se prepara la información que quiere enviar en formato json al streaming
type MyData struct {
	ProductoID string
	Precio     float64
}

func main() {
	myEncoder := json.NewEncoder(os.Stdout)

	data := MyData{
		ProductoID: "XASD",
		Precio:     25.50,
	}

	// se invoca el método Encode.
	// internamente este método hace una especie de marshall y lo escribe el stream
	myEncoder.Encode(data)
}
