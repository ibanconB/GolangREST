package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Handle all HTTP requests here
	fmt.Fprintf(w, "¡Hola, mundo!")
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	// Handle all HTTP requests here
	fmt.Fprintf(w, "¡He añadido un nuevo Handler!")
}

func pedirNombreEdad(w http.ResponseWriter, r *http.Request) {

	nombre := r.URL.Query().Get("nombre")
	edad := r.URL.Query().Get("edad")

	message := fmt.Sprintf("Me llamo %s y tengo %s anios", nombre, edad)

	fmt.Fprintf(w, message)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/nombre", pedirNombreEdad)

	port := 8080

	fmt.Printf("Servidor escuchando en el puerto %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
