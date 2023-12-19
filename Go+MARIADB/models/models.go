package models

type Client struct {
	Id       int
	Nombre   string
	Correo   string
	Telefono string
}

type Clients []Client
