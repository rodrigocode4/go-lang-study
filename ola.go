package main

import "fmt"

func main() {
	fmt.Println(Ola("Mundo"))
}

func Ola(msg string) string {
	if msg == "" {
		msg = "mundo"
	}

	return "Olá, " + msg
}
