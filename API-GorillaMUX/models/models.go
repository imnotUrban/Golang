package models

import "API_MUX_GORM/database"

type Categoria struct {
	Id     uint   `json:"id"`
	Nombre string `gorm:"type:varchar(100)" json:"nombre"`
	Slug   string `gorm:"type:varchar(100)" json:"slug"`
}

type Categorias []Categoria

func Migraciones() {
	database.Database.AutoMigrate(&Categoria{})
}
