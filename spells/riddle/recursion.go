package riddle

func Factorial(i uint) uint {
	if i <= 1 {
		return 1
	}

	return i * Factorial(i-1)
}

func Sum(n uint) uint {
	if n == 0 {
		return n
	}

	return n + Sum(n-1)
}

func FibonacciNumber(n uint) uint {
	if n < 2 {
		return n
	}

	return FibonacciNumber(n-1) + FibonacciNumber(n-2)
}

func FibonacciSequence(n uint) []uint {
	if n <= 1 {
		return []uint{0, 1}
	}

	return append(FibonacciSequence(n-1), FibonacciNumber(n-2)+FibonacciNumber(n-1))
}
