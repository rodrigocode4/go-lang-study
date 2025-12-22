package main

import "testing"

func TestOla(t *testing.T) {
	resutado := Ola("Mundo!")
	esperado := "Olá Mundo!"

	if resutado != esperado {
		t.Errorf("\nresutado: '%s'\nesperado: '%s'", resutado, esperado)
	}

}
