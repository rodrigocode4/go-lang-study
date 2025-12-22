package main

import "testing"

func TestOla(t *testing.T) {
	resutado := Ola("Mundo!")
	esperado := "Olá Mundo!"

	Assert(t, resutado, esperado)

}
