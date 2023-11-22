package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type joke struct {
	Tipo      string `json:"type"`
	Broma     string `json:"joke"`
	Categoria string `json:"category"`
}

func main() {
	resp, err := http.Get("https://v2.jokeapi.dev/joke/Programming?lang=es&type=single")

	if err != nil {
		fmt.Println("Error en solicitud", err)
		return
	}
	defer resp.Body.Close()

	var respJoke joke

	if err := json.NewDecoder(resp.Body).Decode(&respJoke); err != nil {
		fmt.Println("Error al decodificar la respuesta JSON:", err)
		return
	}

	fmt.Println("Tipo:", respJoke.Tipo)
	fmt.Println("Broma:", respJoke.Broma)
	fmt.Println("Categoria:", respJoke.Categoria)

}
