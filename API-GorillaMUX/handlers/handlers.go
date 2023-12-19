package handlers

import (
	"fmt"
	"net/http"
)

func Ejemplo_get(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Pepito GET")
}
