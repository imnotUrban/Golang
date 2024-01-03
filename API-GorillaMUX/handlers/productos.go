package handlers

import (
	"API_MUX_GORM/database"
	"API_MUX_GORM/dto"
	"API_MUX_GORM/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gosimple/slug"
)

func Producto_post(response http.ResponseWriter, request *http.Request) {
	var producto dto.ProductoDto
	if err := json.NewDecoder(request.Body).Decode(&producto); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Recurso no disponible",
		}
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(respuesta)
	}

	datos := models.Producto{
		Nombre:      producto.Nombre,
		Slug:        slug.Make(producto.Nombre),
		Precio:      producto.Precio,
		Stock:       producto.Stock,
		Descripcion: producto.Descripcion,
		CategoriaId: producto.CategoriaId,
		Fecha:       time.Now(),
	}

	database.Database.Save(&datos)
	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Se creó el registro correctamente",
	}
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(respuesta)
}

func Productos_get(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	datos := models.Productos{}

	// database.Database.Find(&datos)
	database.Database.Order("id desc").Preload("Categoria").Find(&datos)
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(datos)

}
