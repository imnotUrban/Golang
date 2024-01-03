package handlers

import (
	"API_MUX_GORM/database"
	"API_MUX_GORM/dto"
	"API_MUX_GORM/jwt"
	"API_MUX_GORM/models"
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Seguridad_login(response http.ResponseWriter, request *http.Request) {
	var registro dto.LoginDto
	response.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(request.Body).Decode(&registro); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Error no Esperado",
		}
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(respuesta)
	}
	// Validar correo y password (TODO)
	user := models.Usuario{}
	if database.Database.Where("correo =?", registro.Correo).Limit(1).Find(&user).RowsAffected > 0 {
		passwordBytes := []byte(registro.Password)
		passwordDB := []byte(user.Password)
		err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

		if err != nil {
			respuesta := map[string]string{
				"estado":  "error",
				"mensaje": "Las credenciales son inválidas",
			}

			response.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(response).Encode(respuesta)
		} else {
			jwtKey, errJWT := jwt.GenerarJWT(user)
			if errJWT != nil {
				respuesta := map[string]string{
					"estado":  "error",
					"mensaje": "Error al generar el token",
				}

				response.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(response).Encode(respuesta)
			}

			retorno := dto.LoginRespuestaDto{
				Nombre: user.Nombre,
				Token:  jwtKey,
			}
			response.WriteHeader(http.StatusOK)
			json.NewEncoder(response).Encode(retorno)

		}
	} else {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Las credenciales son inválidas",
		}

		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(respuesta)
	}
}
func Seguridad_registro(response http.ResponseWriter, request *http.Request) {
	var registro dto.UsuarioDto
	response.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(request.Body).Decode(&registro); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Error no Esperado",
		}
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(respuesta)
	}
	if len(registro.Nombre) == 0 {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "El nombre es obligatorio",
		}

		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(respuesta)

	}
	if len(registro.Correo) == 0 {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "El Correo es obligatorio",
		}

		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(respuesta)

	}
	if len(registro.Telefono) == 0 {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "El Telefono es obligatorio",
		}

		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(respuesta)

	}
	if len(registro.Password) == 0 {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "El contraseña es obligatorio",
		}

		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(respuesta)

	}

	//Validar si existe el correo

	usuario := models.Usuario{}
	if database.Database.Where("correo =? ", registro.Correo).Limit(1).Find(&usuario).RowsAffected > 0 {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "El email ya está siendo usado por otro usuario",
		}

		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(respuesta)
	} else {
		// Generar hash
		costo := 8
		bytes, _ := bcrypt.GenerateFromPassword([]byte(registro.Password), costo)
		datos := models.Usuario{Nombre: registro.Nombre, Correo: registro.Correo, Telefono: registro.Telefono, PerfilID: registro.PerfilID, Password: string(bytes), Fecha: time.Now()}
		database.Database.Save(&datos)
		respuesta := map[string]string{
			"estado":  "ok",
			"mensaje": "Se creo el usuario correctamente",
		}

		response.WriteHeader(http.StatusCreated)
		json.NewEncoder(response).Encode(respuesta)

	}

}

func Seguridad_protegido(response http.ResponseWriter, request *http.Request) {
	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Recurso protegido",
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(respuesta)
}
