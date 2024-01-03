package models

import (
	"API_MUX_GORM/database"
	"time"
)

type Categoria struct {
	Id     uint   `json:"id"`
	Nombre string `gorm:"type:varchar(100)" json:"nombre"`
	Slug   string `gorm:"type:varchar(100)" json:"slug"`
}

type Categorias []Categoria

type Producto struct {
	Id          uint      `json:"id"`
	Nombre      string    `gorm:"type:varchar(100)" json:"nombre"`
	Slug        string    `gorm:"type:varchar(100)" json:"slug"`
	Precio      int       `json:"precio"`
	Stock       int       `json:"stock"`
	Descripcion string    `json:"descripcion"`
	Fecha       time.Time `json:"fecha"`
	CategoriaId uint      `json:"categoria_id"`
	Categoria   Categoria `json:"categoria"`
}
type Productos []Producto

func Migraciones() {
	database.Database.AutoMigrate(&Usuario{}, &Perfil{})
}

type Perfil struct {
	Id     uint   `json:"id"`
	Nombre string `gorm:"type:varchar(100)" json:"nombre"`
}

type Perfiles []Perfil
type Usuario struct {
	Id       uint      `json:"id"`
	PerfilID uint      ` json:"perfil_id"`
	Perfil   Perfil    ` json:"perfil"`
	Nombre   string    `gorm:"type:varchar(100)" json:"nombre"`
	Correo   string    `gorm:"type:varchar(100)" json:"correo"`
	Telefono string    `gorm:"type:varchar(100)" json:"telefono"`
	Password string    `gorm:"type:varchar(100)" json:"password"`
	Fecha    time.Time `json:"fecha"`
}

type Usuarios []Usuario
