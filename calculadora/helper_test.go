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

func Assert2[T Number](t *testing.T, resultado T, esperado T, dado any) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resutado: '%v', esperado: '%v', dado: %v", resultado, esperado, dado)
	}

}
