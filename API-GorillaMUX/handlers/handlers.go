package handlers

import (
	"API_MUX_GORM/dto"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ResponseGenerico struct {
	Estado  string
	Mensaje string
}

func Ejemplo_get(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json") // Se usa para que la api devuelva un json y no un texto plano
	response.Header().Add("Nuevo-Header", "Header bkn")       // Se usa para agregar nuevos campos al header
	output, _ := json.Marshal(ResponseGenerico{"ok", "Método get"})
	fmt.Fprintln(response, string(output))
}

func Ejemplo_get_con_parametros(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprintln(response, "Método get con parámetros | id = ", vars["id"])
}
func Ejemplo_get_query_string(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Método get con query | id = ", request.URL.Query().Get("id"))
}

func Ejemplo_post(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json") // Se usa para que la api devuelva un json y no un texto plano
	response.Header().Add("Nuevo-Header", "Header bkn")       // Se usa para agregar nuevos campos al header

	var categoria dto.CategoriaDto

	err := json.NewDecoder(request.Body).Decode(&categoria)

	if err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode((respuesta))
		return
	}

	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": " Método POST",
		"nombre":  categoria.Nombre,
	}
	// response.WriteHeader(201)                // Se usa para retornar el estado
	response.WriteHeader(http.StatusCreated) // Esta es la forma recomendada, hay muuchos
	json.NewEncoder(response).Encode(respuesta)
}
func Ejemplo_put(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprintln(response, "Método put, id: "+vars["id"])
}
func Ejemplo_delete(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprintln(response, "Método delete, id: "+vars["id"])
}
