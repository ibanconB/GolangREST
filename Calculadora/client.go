package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func realizarOperacion(op string, numero1, numero2 int) {
	url := fmt.Sprintf("http://localhost:8080/%s?numero1=%d&numero2=%d", op, numero1, numero2)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error al realizar la solicitud:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta:", err)
		return
	}

	fmt.Printf("Resultado de %s: %s\n", op, string(body))
}

func main() {
	realizarOperacion("sumar", 10, 3)
	realizarOperacion("restar", 8, 2)
	realizarOperacion("multiplicar", 4, 6)
	realizarOperacion("dividir", 10, 2)
}
