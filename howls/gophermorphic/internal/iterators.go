package gophermorphic

import (
	"bufio"
	"iter"
	"os"
	"strings"
)

// Counter returns an iterator that yields the numbers from 1 to n.
func Counter(n int) func(yield func(int) bool) {

	return func(yield func(int) bool) {

		for i := 1; i <= n; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// EvenNumbers returns an iterator that yields the even numbers from 0 to n.
func EvenNumbers(limit int) func(func(int) bool) {
	return func(yield func(int) bool) {
		for i := 0; i <= limit; i++ {
			if i%2 != 0 {
				if !yield(i) {
					return
				}
			}
		}
	}
}

// Fibonacci returns an iterator that yields the Fibonacci sequence up to n terms.
func Fibonacci(terms int) func(func(int) bool) {
	return func(yield func(int) bool) {
		a, b := 0, 1
		for i := 0; i < terms; i++ {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

// Tokenize returns an iterator that yields the words in a string.
func Tokenize(s string) func(func(string) bool) {
	return func(yield func(string) bool) {
		words := strings.Split(s, " ")
		for _, w := range words {
			if !yield(w) {
				return
			}
		}
	}
}

// ReadLines returns an iterator that yields the lines in a file.
func ReadLines(filename string) func(func(string, error) bool) {
	return func(yield func(string, error) bool) {
		file, err := os.Open(filename)
		if err != nil {
			yield("", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if !yield(scanner.Text(), nil) {
				return
			}
		}

		if err := scanner.Err(); err != nil {
			yield("", err)
		}
	}
}

// MapIterator returns an iterator that yields the key-value pairs in a map.
func MapIterator[K comparable, V any](m map[K]V) func(yield func(K, V) bool) {

	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}

// MapKeys returns an iterator that yields the keys in a map.
func MapKeys[K comparable, V any](m map[K]V) func(func(K) bool) {
	return func(yield func(K) bool) {
		for k := range m {
			if !yield(k) {
				return
			}
		}
	}

}

// MapValues returns an iterator that yields the values in a map.
func MapValues[K comparable, V any](m map[K]V) func(func(V) bool) {
	return func(yield func(V) bool) {
		for _, v := range m {
			if !yield(v) {
				return
			}
		}
	}
}

// Filter returns an iterator that yields the values in a slice that satisfy the predicate.
func Filter[T any](slice []T, predicate func(T) bool) func(yield func(T) bool) {
	return func(yield func(T) bool) {
		for _, v := range slice {
			if predicate(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Backwards returns an iterator that yields the values in a slice in reverse order.
func Backwards[T any](slice []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := len(slice) - 1; i >= 0; i-- {
			if !yield(slice[i]) {
				return
			}
		}
	}
}
