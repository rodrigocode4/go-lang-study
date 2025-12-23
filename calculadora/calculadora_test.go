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

	t.Run("soma dois números decimais", func(t *testing.T) {
		resultado := Soma(2.0, 2.0)
		esperado := 4.0

		Assert(t, resultado, esperado)
	})

	t.Run("diferença de dois números decimais", func(t *testing.T) {
		resultado := Subtrai(2.0, 2.0)
		esperado := 0.0

		Assert(t, resultado, esperado)
	})

	t.Run("multiplicação de dois números decimais", func(t *testing.T) {
		resultado := Multiplica(2, 1.5)
		esperado := 3.0

		Assert(t, resultado, esperado)
	})

	t.Run("divisão de dois números decimais", func(t *testing.T) {
		resultado := Divide(2, 0.5)
		esperado := 4.0

		Assert(t, resultado, esperado)
	})

	t.Run("soma lista com 5 numeros", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5}

		esperado := SomaLista(numeros)
		resultado := 15

		Assert2(t, resultado, esperado, numeros)

	})

	t.Run("soma lista com qualquer tamanho", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4}

		esperado := SomaLista(numeros)
		resultado := 10

		Assert2(t, resultado, esperado, numeros)

	})

}
