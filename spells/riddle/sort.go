package riddles

func QuickSort(arr []int, low, high int) []int {

	if low < high {
		var pivot int
		arr, pivot = partition(arr, low, high)
		arr = QuickSort(arr, low, pivot-1)
		arr = QuickSort(arr, pivot+1, high)
	}

	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
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
