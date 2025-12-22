package calculadora

import (
	"testing"
)

func Assert[T Number](t *testing.T, resultado, esperado T) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resutado: '%v', esperado: '%v'", resultado, esperado)
	}

}
