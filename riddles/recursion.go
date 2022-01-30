package riddles

import "constraints"

func Factorial[T constraints.Unsigned](i T) T {
	if i <= 1 {
		return 1
	}

	return i * Factorial(i-1)
}

func Sum[T constraints.Unsigned](n T) T {
	if n == 0 {
		return n
	}

	return n + Sum(n-1)
}

func FibonacciNumber[T constraints.Unsigned](n T) T {
	if n < 2 {
		return n
	}

	return FibonacciNumber(n-1) + FibonacciNumber(n-2)
}

func FibonacciSequence[T constraints.Unsigned](n T) []T {
	if n <= 1 {
		return []T{0, 1}
	}

	return append(FibonacciSequence(n-1), FibonacciNumber(n-2)+FibonacciNumber(n-1))
}

func GreatestCommonDivisor[T constraints.Unsigned](x, y T) T {
	if y == 0 {
		return x
	}

	return GreatestCommonDivisor(y, x%y)
}
