package connection

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

/*
* Función para conectarse a la base de datos
 */

var Db *sql.DB

func OpenDbConnection() {

	errorVar := godotenv.Load()
	if errorVar != nil {
		panic(errorVar)
	}

	db, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_SERVER")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))

	if err != nil {
		panic(err) //Si hay algún error, causa panic
	}
	Db = db //Asigna la conexión a la variable global DB

}

func CloseDBConnection() {
	Db.Close()
}
