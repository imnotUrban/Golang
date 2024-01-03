package main

import (
	"API_MUX_GORM/handlers"
	"API_MUX_GORM/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*
* Conexión al servidor
 */

func main() {

	//Migrar db

	// models.Migraciones()

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
	mux.HandleFunc(prefijo+"categorias", handlers.Categoria_get).Methods("GET")
	mux.HandleFunc(prefijo+"categorias/{id:[0-9]+}", handlers.Categoria_get_params).Methods("GET")
	mux.HandleFunc(prefijo+"categorias", handlers.Categoria_post).Methods("POST")
	mux.HandleFunc(prefijo+"categorias/{id:[0-9]+}", handlers.Categoria_put).Methods("PUT")
	mux.HandleFunc(prefijo+"categorias/{id:[0-9]+}", handlers.Categoria_delete).Methods("DELETE")
	mux.HandleFunc(prefijo+"productos", handlers.Producto_post).Methods("POST")
	mux.HandleFunc(prefijo+"productos", handlers.Productos_get).Methods("GET")

	mux.HandleFunc(prefijo+"seguridad/registro", handlers.Seguridad_registro).Methods("POST")
	mux.HandleFunc(prefijo+"seguridad/login", handlers.Seguridad_login).Methods("POST")
	mux.HandleFunc(prefijo+"seguridad/protegido", middleware.ValidarJWT(handlers.Seguridad_protegido)).Methods("GET")
	handler := cors.AllowAll().Handler(mux)
	// log.Fatal(http.ListenAndServe(":8084", mux))  -> Se usa cuando no ponemos tenemos mux
	log.Fatal(http.ListenAndServe(":8084", handler))
}
