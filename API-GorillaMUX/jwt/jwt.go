package jwt

import (
	"API_MUX_GORM/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func GenerarJWT(usuario models.Usuario) (string, error) {
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic("Error al cargar el archivo .env")
	}

	key := []byte(os.Getenv("SECRET_JWT"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"correo":         usuario.Correo,
		"nombre":         usuario.Nombre,
		"generado_desde": "ww.xd.com",
		"id":             usuario.Id,
		"iat":            time.Now().Unix(),
		"exp":            time.Now().Add(time.Hour * 24).Unix(), //Expira en 24 hrs
	})
	tokenString, err := token.SignedString(key)
	return tokenString, err

}
