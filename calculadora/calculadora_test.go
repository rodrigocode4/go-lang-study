package calculadora

import "testing"

func TestSoma(t *testing.T) {
	t.Run("soma dois números inteiros", func(t *testing.T) {
		resultado := Soma(2, 2)
		esperado := 4

		Assert(t, resultado, esperado)
	})

	t.Run("diferença de dois números inteiros", func(t *testing.T) {
		resultado := Subtrai(2, 2)
		esperado := 0

		Assert(t, resultado, esperado)
	})

	t.Run("multiplicação de dois números inteiros", func(t *testing.T) {
		resultado := Multiplica(2, 2)
		esperado := 4

		Assert(t, resultado, esperado)
	})

	t.Run("divisão de dois números inteiros", func(t *testing.T) {
		resultado := Divide(2, 2)
		esperado := 1

		Assert(t, resultado, esperado)
	})

}
