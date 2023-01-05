package main

import (
	"github.com/gin-gonic/gin"
)

// Definimos una pseudobase de datos donde consultaremos la información.
var empleados = map[string]string{
	"644": "Empleado A",
	"755": "Empleado B",
	"777": "Empleado C",
}

func main() {
	server := gin.Default()
	server.GET("/", PaginaPrincipal)
	server.GET("/empleados/:id", BuscarEmpleado)
	server.Run(":8085")
}

// Este handler se encargará de responder a /.
func PaginaPrincipal(ctxt *gin.Context) {
	ctxt.String(200, "¡Bienvenido a la Empresa Gophers!")
}

// Este handler verificará si la id que pasa el cliente existe en nuestra base de datos.
func BuscarEmpleado(ctxt *gin.Context) {
	empleado, ok := empleados[ctxt.Param("id")]
	if ok {
		ctxt.String(200, "Información del empleado %s, nombre: %s", ctxt.Param("id"), empleado)
	} else {
		ctxt.String(404, "Información del empleado ¡No existe!")
	}
}

// valorDeNuestraKey := ctx.Query("key") Para acceder a una query
