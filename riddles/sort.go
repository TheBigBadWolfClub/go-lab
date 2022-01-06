package riddles

//QuickSort
// Space complexity: O(1)
// best case: O(n logn)
// average case: O(n logn)
// worst case: O(n^2)
// memory: O(logn)
// stable: no
// Method: partition
func QuickSort(arr []int, low, high int) []int {

	if low < high {
		var pivot int
		arr, pivot = quickSortPartition(arr, low, high)
		arr = QuickSort(arr, low, pivot-1)
		arr = QuickSort(arr, pivot+1, high)
	}

	return arr
}

func quickSortPartition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	index := low

	for i := low; i < high; i++ {

		if arr[i] < pivot {
			arr[index], arr[i] = arr[i], arr[index]
			index++
		}

	}
	arr[index], arr[high] = arr[high], arr[index]

	return arr, index
}

//BubbleSort
// Space complexity: O(1)
// best case: O(n)
// average case: O(n^2)
// worst case: O(n^2)
// memory: O(1)
// stable: yes
// Method: exchanging
func BubbleSort(arr []int) []int {
	var swapped bool
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
			swapped = true
		}
	}

	if swapped {
		return BubbleSort(arr)
	}

	return arr
}

//SelectionSort
// Space complexity: O(1)
// best case: O(n^2)
// average case: O(n^2)
// worst case: O(n^2)
// memory: O(1)
// stable: no
// Method: selection
func SelectionSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	pos := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[pos] {
			pos = i
		}
	}

	value := arr[pos]
	nextSort := SelectionSort(append(arr[:pos], arr[pos+1:]...))
	return append([]int{value}, nextSort...)
}

//InsertionSort
// Space complexity: O(1)
// memory: O(1)
// best case: O(n)
// average case: O(n^2)
// worst case: O(n^2)
// stable: yes
// Method: insertion
func InsertionSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	for i := len(arr) - 1; i > 0; i-- {
		if arr[i-1] > arr[i] {
			arr[i-1], arr[i] = arr[i], arr[i-1]
		}
	}

	return append(arr[:1], InsertionSort(arr[1:])...)
}

//MergeSort
// Space complexity: O(n)
// memory: O(n)
// best case: O(nlogn)
// average case: O(nlogn)
// worst case: O(nlogn)
// stable: yes
// Method: merging
func MergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := MergeSort(items[:len(items)/2])
	second := MergeSort(items[len(items)/2:])
	return mergeSortMerge(first, second)
}
func mergeSortMerge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
