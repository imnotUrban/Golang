package handlers

import (
	"bufio"
	"fmt"
	"go_mysql_driver/connection"
	"go_mysql_driver/models"
	"log"
	"os"
)

/*
* Método listar para mostrar los clientes
 */

func List() {
	connection.OpenDbConnection()

	//Construcción de la consulta
	sql := "SELECT id, nombre, correo, telefono FROM clientes ORDER BY id DESC;"

	//Ejecución de la consulta
	datos, err := connection.Db.Query(sql)

	//Si el error != nil (Hay algún error), lo imprime
	if err != nil {
		fmt.Println(err)
	}

	defer connection.CloseDBConnection() //Se ejecutará al final, por lo que da lo mismo donde ponerlo

	//Inicializa la estructura para almanecar los resultados
	clients := models.Clients{}

	//Recorre los resultados de la query y los almacena en la estructura
	for datos.Next() { //Recorrer los datos
		dato := models.Client{}
		datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)
		clients = append(clients, dato)
	}

}

func ListV2() {
	connection.OpenDbConnection()

	//Construcción de la consulta
	sql := "SELECT id, nombre, correo, telefono FROM clientes ORDER BY id DESC;"

	//Ejecución de la consulta
	datos, err := connection.Db.Query(sql)

	//Si el error != nil (Hay algún error), lo imprime
	if err != nil {
		fmt.Println(err)
	}

	defer connection.CloseDBConnection() //Se ejecutará al final, por lo que da lo mismo donde ponerlo

	for datos.Next() {
		var dato = models.Client{}
		err := datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Id: %v | Nombre: %v | Email: %s | Telefono: %s \n", dato.Id, dato.Nombre, dato.Correo, dato.Telefono)
	}

}

func ListById(id int) {

	connection.OpenDbConnection()

	//Construcción de la consulta
	sql := "SELECT id, nombre, correo, telefono FROM clientes WHERE id =?;"

	//Ejecución de la consulta
	datos, err := connection.Db.Query(sql, id)

	//Si el error != nil (Hay algún error), lo imprime
	if err != nil {
		fmt.Println(err)
	}

	defer connection.CloseDBConnection() //Se ejecutará al final, por lo que da lo mismo donde ponerlo

	for datos.Next() {
		var dato = models.Client{}
		err := datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Id: %v | Nombre: %v | Email: %s | Telefono: %s \n", dato.Id, dato.Nombre, dato.Correo, dato.Telefono)
	}
}

func Insert(cliente models.Client) {

	connection.OpenDbConnection()
	//2023-12-16 12:00:00
	sql := "insert into clientes values(null, ?,?,?,'2023-12-16 12:00:00');"

	result, err := connection.Db.Exec(sql, cliente.Nombre, cliente.Correo, cliente.Telefono)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println("Se creó exitosamente el registro")
	defer connection.CloseDBConnection() //Se ejecutará al final, por lo que da lo mismo donde ponerlo

}

func Edit(client models.Client, id int) {
	connection.OpenDbConnection()

	sql := "update clientes set nombre=?, correo=?, telefono=? where id=?;"
	result, err := connection.Db.Exec(sql, client.Nombre, client.Correo, client.Telefono, id)
	if err != nil {
		panic(err)

	}
	fmt.Println(result)
	fmt.Println("Se editó correctamente el registro con el id: ", id)

}

func Delete(id int) {

	connection.OpenDbConnection()

	sql := "delete from clientes where id = ?;"
	_, err := connection.Db.Exec(sql, id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Se eliminó el registro")

}

func Execute() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Seleccione una opción \n\n")
	fmt.Println("1- Listar clientes\n")
	fmt.Println("2- Listar clientes por ID\n")
	fmt.Println("3- Crear cliente\n")
	fmt.Println("4- Editar cliente\n")
	fmt.Println("5- Eliminar cliente\n")
	if scanner.Scan() {
		for { //For infinito
			if scanner.Text() == "1" {
				ListV2()
				return
			}
			// if scanner.Text() == "2" {
			// 	fmt.Println("Ingrese el ID del cliente: \n")
			// 	if scanner.Scan() {
			// 		ID, _ = strconv.Atoi(scanner.Text())
			// 	}
			// 	ListById(ID)
			// 	return

			// }
		}
	}
}
