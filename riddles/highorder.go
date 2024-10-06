package riddles

// Filter filters the array arr based on the function f
func Filter[T comparable](arr []T, f func(T) bool) []T {
	var res []T
	for _, v := range arr {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

// Map maps the array arr based on the function f
func Map[T comparable, Z comparable](arr []T, f func(T) Z) []Z {
	var res []Z
	for _, v := range arr {
		res = append(res, f(v))
	}
	return res
}

// Reduce reduces the array arr based on the function f
func Reduce[T comparable](arr []T, f func(T, T) T) T {
	var res T
	if len(arr) == 0 {
		return res
	}

	res = arr[0]
	for i := 1; i < len(arr); i++ {
		res = f(res, arr[i])
	}
	return res
}

// Any returns true if any element in the array arr satisfies the function f
func Any[T comparable](arr []T, f func(T) bool) bool {
	for _, v := range arr {
		if f(v) {
			return true
		}
	}
	return false
}

// All returns true if all elements in the array arr satisfy the function f
func All[T comparable](arr []T, f func(T) bool) bool {
	for _, v := range arr {
		if !f(v) {
			return false
		}
	}
	return true
}
