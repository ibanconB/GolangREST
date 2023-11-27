package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func sumar(w http.ResponseWriter, r *http.Request) {
	valores, ok := obtenerValores(r)
	if !ok {
		http.Error(w, "Parámetros inválidos", http.StatusBadRequest)
		return
	}

	resultado := valores.numero1 + valores.numero2
	fmt.Fprintf(w, "Suma: %d", resultado)
}

func restar(w http.ResponseWriter, r *http.Request) {
	valores, ok := obtenerValores(r)
	if !ok {
		http.Error(w, "Parámetros inválidos", http.StatusBadRequest)
		return
	}

	resultado := valores.numero1 - valores.numero2
	fmt.Fprintf(w, "Resta: %d", resultado)
}

func multiplicar(w http.ResponseWriter, r *http.Request) {
	valores, ok := obtenerValores(r)
	if !ok {
		http.Error(w, "Parámetros inválidos", http.StatusBadRequest)
		return
	}

	resultado := valores.numero1 * valores.numero2
	fmt.Fprintf(w, "Multiplicación: %d", resultado)
}

func dividir(w http.ResponseWriter, r *http.Request) {
	valores, ok := obtenerValores(r)
	if !ok || valores.numero2 == 0 {
		http.Error(w, "Parámetros inválidos o división por cero", http.StatusBadRequest)
		return
	}

	resultado := valores.numero1 / valores.numero2
	fmt.Fprintf(w, "División: %d", resultado)
}

type valoresOperacion struct {
	numero1 int
	numero2 int
}

func obtenerValores(r *http.Request) (valoresOperacion, bool) {
	numero1Str := r.FormValue("numero1")
	numero2Str := r.FormValue("numero2")

	numero1, err1 := strconv.Atoi(numero1Str)
	numero2, err2 := strconv.Atoi(numero2Str)

	if err1 != nil || err2 != nil {
		return valoresOperacion{}, false
	}

	return valoresOperacion{numero1, numero2}, true
}

func main() {
	http.HandleFunc("/sumar", sumar)
	http.HandleFunc("/restar", restar)
	http.HandleFunc("/multiplicar", multiplicar)
	http.HandleFunc("/dividir", dividir)

	fmt.Println("Servidor en ejecución en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
