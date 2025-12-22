package main

import "testing"

func TestOla(t *testing.T) {

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resutado := Ola("Rodrigo!")
		esperado := "Olá, Rodrigo!"

		Assert(t, resutado, esperado)
	})

	t.Run("diz 'Olá, mundo' quando string estiver vazia", func(t *testing.T) {
		resutado := Ola("")
		esperado := "Olá, mundo"

		Assert(t, resutado, esperado)
	})

}
