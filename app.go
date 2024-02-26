package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Skills string
	Age    int
}

func Index(rw http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/index.html")

	user := User{"Virginia Castellano", "Desarrolladora", 21}

	if err != nil {
		panic(err)
	} else {
		template.Execute(rw, user)
	}
}

func main() {
	// Manejar archivos estáticos (imágenes)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", Index)

	fmt.Println("El puerto está corriendo en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000")
	http.ListenAndServe("localhost:3000", nil)
}
