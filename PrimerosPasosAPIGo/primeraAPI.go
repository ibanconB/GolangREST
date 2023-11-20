package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Respuesta representa la estructura de la respuesta de la API
type Respuesta struct {
	Respuesta string `json:"answer"`
	ImagenURL string `json:"image"`
}

func main() {
	// Realiza una solicitud GET a la API
	resp, err := http.Get("https://yesno.wtf/api")
	if err != nil {
		fmt.Println("Error al realizar la solicitud:", err)
		return
	}
	defer resp.Body.Close()

	// Decodifica la respuesta JSON
	var respuesta Respuesta
	if err := json.NewDecoder(resp.Body).Decode(&respuesta); err != nil {
		fmt.Println("Error al decodificar la respuesta JSON:", err)
		return
	}

	// Imprime la respuesta
	fmt.Println("Respuesta:", respuesta.Respuesta)
	fmt.Println("URL de la imagen:", respuesta.ImagenURL)
}
