package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type PageVariables struct {
	Resultado string
}

func main() {
	http.HandleFunc("/", HomePage)
	http.ListenAndServe(":8080", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		numeroA, _ := strconv.ParseFloat(r.FormValue("numeroA"), 64)
		numeroB, _ := strconv.ParseFloat(r.FormValue("numeroB"), 64)
		calculo := r.FormValue("calculo")

		res := Calcular(numeroA, numeroB, calculo)

		pageVariables := PageVariables{
			Resultado: fmt.Sprintf("El resultado de %.2f %s %.2f es %.2f", numeroA, calculo, numeroB, res),
		}

		renderizarPagina(w, "index.html", pageVariables)
		return

	}

	renderizarPagina(w, "index.html", PageVariables{})
}

func Calcular(numeroA, numeroB float64, calculo string) float64 {
	switch calculo {
	case "+":
		return numeroA + numeroB
	case "-":
		return numeroA - numeroB
	case "*":
		return numeroA * numeroB
	case "/":
		if numeroB != 0 {
			return numeroA / numeroB
		} else {
			return 0
		}
	default:
		return 0
	}
}

func renderizarPagina(w http.ResponseWriter, tmplFile string, data PageVariables) {
	tmpl, err := template.New("index").ParseFiles(tmplFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
