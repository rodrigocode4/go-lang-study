package main

import "fmt"

func main() {
	fmt.Println(Ola("Mundo"))
}

func Ola(msg string) string {
	return "Olá " + msg
}
