package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	// URL del servidor
	serverURL := "http://localhost:8080"

	// Leer el contenido del archivo index.html
	htmlContent, err := ioutil.ReadFile("./index.html")
	if err != nil {
		fmt.Println("Error al leer el archivo index.html:", err)
		return
	}

	// Datos para la solicitud POST
	numeroA := 10.5
	numeroB := 5.2
	calculo := "+"

	// Crear los datos del formulario
	formData := url.Values{
		"numeroA": {strconv.FormatFloat(numeroA, 'f', -1, 64)},
		"numeroB": {strconv.FormatFloat(numeroB, 'f', -1, 64)},
		"calculo": {calculo},
	}

	// Convertir el contenido HTML en un cuerpo de solicitud
	formData.Set("valor1", string(htmlContent))

	// Realizar la solicitud POST al servidor con el contenido HTML como datos
	resp, err := http.PostForm(serverURL, formData)
	if err != nil {
		fmt.Println("Error en la solicitud POST:", err)
		return
	}
	defer resp.Body.Close()

	// Leer la respuesta del servidor
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	responseBody := buf.String()

	// Imprimir la respuesta del servidor
	fmt.Println("Respuesta del servidor:")
	fmt.Println(responseBody)
}
