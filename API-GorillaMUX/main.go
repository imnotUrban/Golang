package main

import (
	"API_MUX_GORM/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*
* Conexión al servidor
 */

func main() {
	mux := mux.NewRouter()
	prefijo := "/api/"

	mux.HandleFunc(prefijo+"ejemplo", handlers.Ejemplo_get).Methods("GET")
	mux.HandleFunc(prefijo+"ejemplo/{id:[0-9]+}", handlers.Ejemplo_get_con_parametros).Methods("GET")
	mux.HandleFunc(prefijo+"query-string", handlers.Ejemplo_get_query_string).Methods("GET")
	mux.HandleFunc(prefijo+"ejemplo", handlers.Ejemplo_post).Methods("POST")
	mux.HandleFunc(prefijo+"ejemplo/{id:[0-9]+}", handlers.Ejemplo_put).Methods("PUT")
	mux.HandleFunc(prefijo+"ejemplo/{id:[0-9]+}", handlers.Ejemplo_delete).Methods("DELETE")
	mux.HandleFunc(prefijo+"upload", handlers.Upload_file).Methods("POST")
	mux.HandleFunc(prefijo+"view", handlers.View_file).Methods("GET")

	handler := cors.AllowAll().Handler(mux)
	// log.Fatal(http.ListenAndServe(":8084", mux))  -> Se usa cuando no ponemos tenemos mux
	log.Fatal(http.ListenAndServe(":8084", handler))
}
