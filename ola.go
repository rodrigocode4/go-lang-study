package main

import "fmt"

func main() {
	fmt.Println(Ola("", ""))
}

const prefixoOlaPortugues = "Olá, "
const prefixoOlaIngles = "Hello, "
const prefixoOlaEspanhol = "Hola, "

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}

	return prefixoSaldacao(idioma) + nome
}

func prefixoSaldacao(idioma string) (prefixo string) {
	switch idioma {
	case "en":
		prefixo = prefixoOlaIngles
	case "es":
		prefixo = prefixoOlaEspanhol
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}
