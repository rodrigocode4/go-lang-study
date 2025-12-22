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

	prefixo := prefixoOlaPortugues

	switch idioma {
	case "en":
		prefixo = prefixoOlaIngles
	case "es":
		prefixo = prefixoOlaEspanhol
	default:
		prefixo = prefixoOlaPortugues
	}

	return prefixo + nome
}
