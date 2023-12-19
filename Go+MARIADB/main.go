package main

import (
	"go_mysql_driver/handlers"
)

func main() {

	// newClient := models.Client{Nombre: "Pepe", Correo: "pepe@gmail.com", Telefono: "000000000"}
	// handlers.Edit(newClient, 1)
	// handlers.Delete(1)
	handlers.Execute()
}
