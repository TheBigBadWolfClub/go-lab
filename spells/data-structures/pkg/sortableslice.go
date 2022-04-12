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

type SortedSlice[T constraints.Ordered] []T

//SelectionSortIterative
// > Selection sort works by selecting the smallest element from an unsorted list
// > and moving it to the front.
// Selection sort: time complexity of O(n²) and space complexity of O(1)
func (s *SortedSlice[T]) SelectionSortIterative() {
	arr := ([]T)(*s)
	for i := range arr {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}

//BubbleSortIterative
// > repeatedly comparing pairs of adjacent elements and then swapping their positions
// > if they are in the wrong order.
// > repeat until no swap was executed
//Bubble sort: time complexity of O(n²) and space complexity of O(1)
func (s *SortedSlice[T]) BubbleSortIterative() {
	arr := ([]T)(*s)
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

//InsertionSortIterative
// > works by inserting elements from an unsorted list
// > into a sorted subsection of the list,
// > one item at a time.
//Insertion sort: time complexity of O(n²) and space complexity of O(1)
func (s *SortedSlice[T]) InsertionSortIterative() {
	arr := ([]T)(*s)
	for i := 1; i < len(arr); i++ {
		pivot := i
		for pivot > 0 && arr[pivot] < arr[pivot-1] {
			arr[pivot], arr[pivot-1] = arr[pivot-1], arr[pivot]
			pivot--
		}
	}
}

//MergeSortRecursive
// > is a recursive algorithm that works like this:
// > - split the input in half
// > - sort each half by recursively using this same process
// > - merge the sorted halves back together
//Merge sort: time complexity of O(n log n) and space complexity of O(n)
func (s *SortedSlice[T]) MergeSortRecursive() SortedSlice[T] {

	arr := ([]T)(*s)
	if len(arr) <= 1 {
		return arr
	}

	var loArr, hiArr SortedSlice[T]
	loArr = arr[:len(arr)/2]
	lo := loArr.MergeSortRecursive()

	hiArr = arr[len(arr)/2:]
	hi := hiArr.MergeSortRecursive()
	return mergeIt(lo, hi)
}

func mergeIt[T constraints.Ordered](lo, hi SortedSlice[T]) SortedSlice[T] {
	var mergeArr SortedSlice[T]
	for len(([]T)(lo)) > 0 && len(([]T)(hi)) > 0 {
		if lo[0] < hi[0] {
			mergeArr = append(([]T)(mergeArr), lo[0])
			lo = lo[1:]
		} else {
			mergeArr = append(([]T)(mergeArr), hi[0])
			hi = hi[1:]
		}
	}

	mergeArr = append(([]T)(mergeArr), lo...)
	mergeArr = append(([]T)(mergeArr), hi...)
	return mergeArr
}

//QuickSortRecursive
// > works by dividing the input into two smaller lists:
// > - one with small items
// > - the other with large items.
// > Then, it recursively sorts both  lists.
//Quicksort: time complexity of O(n log n) and space complexity of O(log n)
func (s *SortedSlice[T]) QuickSortRecursive() SortedSlice[T] {
	arr := ([]T)(*s)
	if len(arr) <= 1 {
		return arr
	}

	pivot := s.quickSortPartitionR()
	var loArr, HiArr SortedSlice[T]
	loArr = arr[:pivot]
	HiArr = arr[pivot:]
	lo := loArr.QuickSortRecursive()
	hi := HiArr.QuickSortRecursive()

	var res SortedSlice[T]
	res = append(([]T)(res), lo...)
	res = append(([]T)(res), hi...)
	return res

}

func (s *SortedSlice[T]) quickSortPartitionR() int {
	arr := ([]T)(*s)
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
func (s *SortedSlice[T]) HeapSortByType(sortType HeapSortType) {
	arr := ([]T)(*s)
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
