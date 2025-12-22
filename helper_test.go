package main

import "testing"

func Assert(t *testing.T, resultado string, esperado string) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resutado: '%s', esperado: '%s'", resultado, esperado)
	}

}
