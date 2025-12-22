package main

func main() {

}

const prefixoOlaPortugues = "Olá, "
const prefixoOlaIngles = "Hello, "
const prefixoOlaEspanhol = "Hola, "

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}

	if idioma == "en" {
		return prefixoOlaIngles + nome
	}

	if idioma == "es" {
		return prefixoOlaEspanhol + nome
	}

	return prefixoOlaPortugues + nome
}
