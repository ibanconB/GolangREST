package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://localhost:8080/nombre?nombre=Ivan&edad=21"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error al hacer la solicitud: ", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta:", err)
		return
	}

	fmt.Println("Respuesta del servidor:", string(body))

}
