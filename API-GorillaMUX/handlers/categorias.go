package handlers

import (
	"API_MUX_GORM/database"
	"API_MUX_GORM/dto"
	"API_MUX_GORM/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gosimple/slug"
)

func Categoria_get(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	datos := models.Categorias{}

	// database.Database.Find(&datos)
	database.Database.Order("id desc").Find(&datos)
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(datos)

}

func Categoria_get_params(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	var id, _ = strconv.Atoi(vars["id"])
	response.Header().Set("Content-Type", "application/json")
	datos := models.Categorias{}

	if err := database.Database.First(&datos, id); err.Error != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Recurso no disponible",
		}
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(respuesta)
	} else {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(datos)
	}

}

func Categoria_post(response http.ResponseWriter, request *http.Request) {
	var categoria dto.CategoriaDto
	if err := json.NewDecoder(request.Body).Decode(&categoria); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Recurso no disponible",
		}
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(respuesta)
	}

	datos := models.Categoria{
		Nombre: categoria.Nombre,
		Slug:   slug.Make(categoria.Nombre),
	}

	database.Database.Save(&datos)
	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Se creó el registro correctamente",
	}
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(respuesta)
}
