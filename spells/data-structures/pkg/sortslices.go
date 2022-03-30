package pkg

// Selection sort: time complexity of O(n²) and space complexity of O(1)
// Insertion sort: time complexity of O(n²) and space complexity of O(1)
// Bubble sort: time complexity of O(n²) and space complexity of O(1)
// Merge sort: time complexity of O(n log n) and space complexity of O(n)
// Quicksort: time complexity of O(n log n) and space complexity of O(log n)
// Heap sort: time complexity of O(n log n) and space complexity of O(1)
//
// Resources
// https://www.interviewcake.com/sorting-algorithm-cheat-sheet
// https://www.fullstack.cafe/blog/sorting-algorithms-interview-questions

import (
	"golang.org/x/exp/constraints"
)

//SelectionSortIterative
// > Selection sort works by selecting the smallest element from an unsorted list
// > and moving it to the front.
// Selection sort: time complexity of O(n²) and space complexity of O(1)
func SelectionSortIterative[T constraints.Ordered](arr []T) {
	for i := range arr {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}

func SelectionSortRecursive[T constraints.Ordered](arr []T, index int, minIdx int) {
	arrLen := len(arr)
	if arrLen == 0 {
		return
	}

	min := selectionSortRecursiveMin(arr, 0, 0)
	arr[0], arr[min] = arr[min], arr[0]

	SelectionSortRecursive(arr[1:], 0, 0)
}

func selectionSortRecursiveMin[T constraints.Ordered](arr []T, cursor, minIdx int) int {
	if cursor >= len(arr) {
		return minIdx
	}

	if arr[cursor] < arr[minIdx] {
		minIdx = cursor
	}

	cursor++
	return selectionSortRecursiveMin(arr, cursor, minIdx)
}

//BubbleSortIterative
// > repeatedly comparing pairs of adjacent elements and then swapping their positions
// > if they are in the wrong order.
// > repeat until no swap was executed
//Bubble sort: time complexity of O(n²) and space complexity of O(1)
func BubbleSortIterative[T constraints.Ordered](arr []T) {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func BubbleSortRecursive[T constraints.Ordered](arr []T) {
	if len(arr) == 0 {
		return
	}
	BubbleSortRecursiveSwap(arr, 0)
	BubbleSortRecursive(arr[:len(arr)-1])
}

func BubbleSortRecursiveSwap[T constraints.Ordered](arr []T, cursor int) {
	if cursor >= len(arr)-1 {
		return
	}

	if arr[cursor+1] < arr[cursor] {
		arr[cursor], arr[cursor+1] = arr[cursor+1], arr[cursor]
	}

	cursor++
	BubbleSortRecursiveSwap(arr, cursor)
}

//InsertionSortIterative
// > works by inserting elements from an unsorted list
// > into a sorted subsection of the list,
// > one item at a time.
//Insertion sort: time complexity of O(n²) and space complexity of O(1)
func InsertionSortIterative[T constraints.Ordered](arr []T) {
	for i := 1; i < len(arr); i++ {
		pivot := i
		for pivot > 0 && arr[pivot] < arr[pivot-1] {
			arr[pivot], arr[pivot-1] = arr[pivot-1], arr[pivot]
			pivot--
		}
	}
}

func InsertionSortRecursive[T constraints.Ordered](arr []T, cursor int) {
	if cursor >= len(arr) {
		return
	}

	insertionSortRecursiveSwap(arr, cursor)
	cursor++
	InsertionSortRecursive(arr, cursor)
}
func insertionSortRecursiveSwap[T constraints.Ordered](arr []T, cursor int) {
	if cursor < 1 {
		return
	}
	if arr[cursor] < arr[cursor-1] {
		arr[cursor], arr[cursor-1] = arr[cursor-1], arr[cursor]
	}
	cursor--
	insertionSortRecursiveSwap(arr, cursor)
}

//MergeSortRecursive
// > is a recursive algorithm that works like this:
// > - split the input in half
// > - sort each half by recursively using this same process
// > - merge the sorted halves back together
//Merge sort: time complexity of O(n log n) and space complexity of O(n)
func MergeSortRecursive[T constraints.Ordered](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	lo := MergeSortRecursive(arr[:len(arr)/2])
	hi := MergeSortRecursive(arr[len(arr)/2:])
	return mergeIt(lo, hi)

}

func mergeIt[T constraints.Ordered](lo, hi []T) []T {
	var mergeArr []T
	for len(lo) > 0 && len(hi) > 0 {
		if lo[0] < hi[0] {
			mergeArr = append(mergeArr, lo[0])
			lo = lo[1:]
		} else {
			mergeArr = append(mergeArr, hi[0])
			hi = hi[1:]
		}
	}

	mergeArr = append(mergeArr, lo...)
	mergeArr = append(mergeArr, hi...)
	return mergeArr
}

//QuickSortRecursive
// > works by dividing the input into two smaller lists:
// > - one with small items
// > - the other with large items.
// > Then, it recursively sorts both  lists.
//Quicksort: time complexity of O(n log n) and space complexity of O(log n)
func QuickSortRecursive[T constraints.Ordered](arr []T) []T {

	if len(arr) <= 1 {
		return arr
	}

	pivot := quickSortPartitionR(arr)
	lo := QuickSortRecursive(arr[:pivot])
	hi := QuickSortRecursive(arr[pivot:])

	var res []T
	res = append(res, lo...)
	res = append(res, hi...)
	return res

}

func quickSortPartitionR[T constraints.Ordered](arr []T) int {
	lastHi := 0
	pivotIdx := len(arr) - 1

	for i := 0; i < pivotIdx; i++ {
		if arr[i] < arr[pivotIdx] {
			arr[i], arr[lastHi] = arr[lastHi], arr[i]
			lastHi++
		}
	}

	arr[lastHi], arr[pivotIdx] = arr[pivotIdx], arr[lastHi]
	return lastHi
}

//HeapSortByType
// > repeatedly choosing the largest item and moving it to the end of our list.
// > The main difference is that instead of scanning through the entire list
// > to find the largest item.
// > convert the list into a max heap to speed things up.
//Heap sort: time complexity of O(n log n) and space complexity of O(1)
func HeapSortByType[T constraints.Ordered](arr []T, sortType HeapSortType) {
	// 1- get max
	// Index of last non-leaf node
	startIdx := (len(arr) / 2) - 1
	for i := startIdx; i >= 0; i-- {
		heapify(arr, i, sortType)
	}

	// 2- order all, take one by one from heap top
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr[:i], 0, sortType)
	}
}

type HeapSortType string

const MAX = "max"
const MIN = "min"

func heapify[T constraints.Ordered](arr []T, parentIdx int, sortType HeapSortType) {

	// Parent node is index/2.
	// Left child node is 2*index + 1.
	// Right child node is 2*index + 2.
	l := 2*parentIdx + 1
	r := 2*parentIdx + 2

	largerIdx := parentIdx

	switch sortType {
	case MIN:
		largerIdx = heapMinFn(arr, l, largerIdx)
		largerIdx = heapMinFn(arr, r, largerIdx)
	case MAX:
		largerIdx = heapMaxFn(arr, l, largerIdx)
		largerIdx = heapMaxFn(arr, r, largerIdx)
	}

	if largerIdx != parentIdx {
		// Recursively heapify the affected sub-tree
		arr[parentIdx], arr[largerIdx] = arr[largerIdx], arr[parentIdx]
		heapify(arr, largerIdx, sortType)
	}
}

func heapMinFn[T constraints.Ordered](arr []T, testIdx, bestIdx int) int {
	if testIdx < len(arr) && arr[testIdx] < arr[bestIdx] {
		return testIdx
	}
	return bestIdx
}

func heapMaxFn[T constraints.Ordered](arr []T, testIdx, bestIdx int) int {
	if testIdx < len(arr) && arr[testIdx] > arr[bestIdx] {
		return testIdx
	}
	return bestIdx
}
