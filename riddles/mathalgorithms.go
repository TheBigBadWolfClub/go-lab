package riddles

import (
	"golang.org/x/exp/constraints"
)

func Factorial[T constraints.Unsigned](i T) T {
	if i == 0 {
		return 1
	}

	return i * Factorial(i-1)
}

func Sum[T constraints.Unsigned](n T) T {
	if n == 0 {
		return 0
	}
	return n + Sum(n-1)
}

func GreatestCommonDivisor[T constraints.Unsigned](x, y T) T {
	if y == 0 {
		return x
	}

	return GreatestCommonDivisor(y, x%y)
}

func FibonacciNumber[T constraints.Unsigned](n T) T {
	if n <= 1 {
		return n
	}
	return FibonacciNumber(n-1) + FibonacciNumber(n-2)
}

func FibonacciSequence[T constraints.Unsigned](n T) []T {
	if n == 0 {
		return []T{n}
	}

	sequence := FibonacciSequence(n - 1)
	return append(([]T)(sequence), FibonacciNumber(n))
}

func FibonacciNumberMemoization[T constraints.Unsigned](n T) T {
	cache := make(map[T]T)

	memoiz := func(m T) T {
		if _, ok := cache[m]; !ok {
			cache[m] = FibonacciNumber(m)
		}
		return cache[m]
	}

	fibo := func(m T) T {
		if m <= 1 {
			return m
		}
		return memoiz(m-1) + memoiz(m-2)
	}

	return fibo(n)
}

func ZeroOnKnapsack[T constraints.Unsigned](n T) T {
	panic("todo")

}

func UnboundedKnapsack[T constraints.Unsigned](n T) T {
	panic("todo")
}
