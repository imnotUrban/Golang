package main

import (
	"API_MUX_GORM/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
* Conexión al servidor
 */

func main() {
	mux := mux.NewRouter()
	prefijo := "/api/"
	mux.HandleFunc(prefijo+"ejemplo", handlers.Ejemplo_get)

	log.Fatal(http.ListenAndServe(":8084", mux))
}
