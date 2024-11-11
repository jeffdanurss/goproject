package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "Hola mundo distribuida")
	})
	//Crea el servidor
	fmt.Println("El servidor esta corriendo en el puerto 3000")
	fmt.Println("Run server:http://localhost:3000")
	http.ListenAndServe("localhost:3000", nil)
}
