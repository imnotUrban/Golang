// database.go
package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB // Change the type to *gorm.DB

func init() {
	// Database initialization code

	// Load .env file
	errVar := godotenv.Load()
	if errVar != nil {
		panic(errVar)
	}

	// Build DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_SERVER"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	// Open the database connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error de conexión")
		panic(err)
	}

	// Assign the opened database connection to the Database variable
	Database = db
	fmt.Println("Conexión Exitosa")
}
