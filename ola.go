package main

import "fmt"

func main() {
	fmt.Println(Ola("Mundo"))
}

const prefixoOlaPortugues = "Olá, "

func Ola(msg string) string {
	if msg == "" {
		msg = "mundo"
	}

	return prefixoOlaPortugues + msg
}
