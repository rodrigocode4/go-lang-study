package calculadora

type Number interface {
	int | float64
}

func Soma[T Number](a, b T) T {
	return a + b
}

func Subtrai[T Number](a, b T) T {
	return a - b
}

func Multiplica[T Number](a, b T) T {
	return a * b
}

func Divide[T Number](a, b T) T {
	return a / b
}

func SomaLista(numeros [5]int) int {
	var soma int
	for _, num := range numeros {
		soma += num
	}

	return soma
}
