package handlers

import (
	"API_MUX_GORM/dto"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

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
		"estado":        "ok",
		"mensaje":       " Método POST",
		"nombre":        categoria.Nombre,
		"Authorization": request.Header.Get("Authorization"),
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

// func Upload_file(response http.ResponseWriter, request *http.Request) {
// 	file, handler, err := request.FormFile("foto")
// 	var extension = strings.Split(handler.Filename, ".")[1]
// 	time := strings.Split(time.Now().String(), "")
// 	foto := string(time[4][6:14]) + "." + extension
// 	var archivo string = "public/uploads/fotos/" + foto
// 	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0777)
// 	if err != nil {
// 		http.Error(response, "Error al subir la imagen"+err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	_, err = io.Copy(f, file)
// 	if err != nil {
// 		http.Error(response, "Error al copiar la imagen "+err.Error(), http.StatusBadGateway)
// 		return
// 	}

//		respuesta := map[string]string{
//			"estado":  "ok",
//			"mensaje": " Imagen subida correctamente ",
//		}
//		response.WriteHeader(http.StatusCreated)
//		json.NewEncoder(response).Encode(respuesta)
//	}
func Upload_file(response http.ResponseWriter, request *http.Request) {
	// Parsea la solicitud para obtener el archivo
	file, handler, err := request.FormFile("foto")
	if err != nil {
		http.Error(response, "Error al obtener el archivo: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Obtiene la extensión del archivo
	extension := strings.Split(handler.Filename, ".")[1]

	// Obtener la fecha actual con un formato específico
	now := time.Now()
	formattedDate := now.Format("20060102150405") // AñoMesDiaHoraMinutoSegundo

	// Generar el nombre del archivo con la fecha y la extensión
	foto := formattedDate + "." + extension

	// Abre el archivo para escritura
	archivo, err := os.OpenFile("public/uploads/fotos/"+foto, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		http.Error(response, "Error al abrir el archivo: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer archivo.Close()

	// Copia el contenido del archivo recibido al archivo destino
	_, err = io.Copy(archivo, file)
	if err != nil {
		http.Error(response, "Error al copiar el archivo: "+err.Error(), http.StatusBadGateway)
		return
	}

	// Envía una respuesta JSON indicando que la imagen se subió correctamente
	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Imagen subida correctamente",
	}
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(respuesta)
}

func View_file(response http.ResponseWriter, request *http.Request) {

	file := request.URL.Query().Get("file")
	if len(file) < 1 {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(respuesta)
		return
	}
	OpenFile, err := os.Open("public/uploads/" + request.URL.Query().Get("folder") + "/" + file)
	if err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(respuesta)
		return

	}
	_, err = io.Copy(response, OpenFile)
	if err != nil {
		http.Error(response, "Error al copiar el archivo", http.StatusBadRequest)
	}
}
