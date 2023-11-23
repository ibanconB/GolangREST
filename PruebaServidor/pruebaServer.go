package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Handle all HTTP requests here
	fmt.Fprintf(w, "¡Hola, mundo!")
}

func main() {
	http.HandleFunc("/", handler)

	port := 8080

	fmt.Printf("Servidor escuchando en el puerto %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
