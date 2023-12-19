package main

import (
	"fmt"
	"go_web/routes"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	mux := mux.NewRouter()

	mux.HandleFunc("/", routes.Home)
	mux.HandleFunc("/AboutUs", routes.AboutUs)
	mux.HandleFunc("/params/{id:.*}/{slug:.*}", routes.Params)
	mux.HandleFunc("/params-query", routes.ParamsQuery)
	mux.HandleFunc("/estructuras", routes.Estructuras)

	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
	}

	/*
		* Carga archivos estáticos
		 -> Así reconoce los archivos .js, .css ,etc
	*/

	s := http.StripPrefix("/public/", http.FileServer(http.Dir("./public/")))

	mux.PathPrefix("/public/").Handler(s)

	/*
	* Ejecución del servidor
	 */
	server := &http.Server{
		Addr:         os.Getenv("SERVER") + ":" + os.Getenv("PORT"),
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Server running from " + os.Getenv("Server") + ":" + os.Getenv("PORT"))
	log.Fatal(server.ListenAndServe())
}
