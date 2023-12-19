package routes

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("templates/example/home.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, nil)
	}
}

func AboutUs(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("templates/example/nosotros.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, nil)
	}
}
func Params(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	data := map[string]string{
		"id":    vars["id"],
		"slug":  vars["slug"],
		"Texto": "Este es un texto de ejemplo x",
	}
	template, err := template.ParseFiles("templates/example/parametros.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, data)
	}
}

func ParamsQuery(response http.ResponseWriter, request *http.Request) {

	// fmt.Fprintln(response, request.URL)
	// fmt.Fprintln(response, request.URL.RawQuery)
	// data := request.URL.Query()
	data := map[string]string{
		"id":    request.URL.Query().Get("id"),
		"slug":  request.URL.Query().Get("slug"),
		"texto": "Este es un texto de ejemplo x",
	}
	// fmt.Fprintln(response, request.URL.Query().Get("id"))
	// fmt.Fprintln(response, request.URL.Query().Get("slug"))

	template, err := template.ParseFiles("templates/example/parametrosquery.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, data)
	}
}

type Datos struct {
	Nombre      string
	Edad        int
	Perfil      int
	Habilidades []Habilidad
}

type Habilidad struct {
	Nombre string
}

func Estructuras(response http.ResponseWriter, request *http.Request) {
	habilidad1 := Habilidad{"Inteligencia"}
	habilidad2 := Habilidad{"Inteligencia1"}
	habilidad3 := Habilidad{"Inteligencia2"}
	habilidades := []Habilidad{habilidad1, habilidad2, habilidad3}

	template, err := template.ParseFiles("templates/example/estructuras.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, Datos{"Pepe", 1, 3, habilidades})
	}
}
