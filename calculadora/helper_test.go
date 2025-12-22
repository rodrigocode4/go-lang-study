package calculadora

import "testing"

func Assert(t *testing.T, resultado int, esperado int) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resutado: '%d', esperado: '%d'", resultado, esperado)
	}

}
