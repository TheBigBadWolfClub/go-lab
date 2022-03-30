package pkg

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

var testsExt = []struct {
	name     string
	array    []int
	expected []int
}{
	{
		name:     "sort empty",
		array:    []int{},
		expected: []int{},
	}, {
		name:     "sort one elem",
		array:    []int{0},
		expected: []int{0},
	}, {
		name:     "sort 2 elem",
		array:    []int{1, 0},
		expected: []int{0, 1},
	}, {
		name:     "sort multiple elem",
		array:    []int{1, 0, 5, 6, 3, 9, 7, 8, 2, 4},
		expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
}

func TestSelectionSortIterative(t *testing.T) {
	t.Parallel()
	for _, tt := range testsExt {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			testArr := make([]int, len(tt.array))
			copy(testArr, tt.array)
			SelectionSortIterative(testArr)
			diff := cmp.Diff(testArr, tt.expected)
			if diff != "" {
				t.Fatalf("SelectionSortIterative, fail to sort: %s", diff)
			}
		})
	}
}

func TestSelectionSortRecursive(t *testing.T) {
	t.Parallel()
	for _, tt := range testsExt {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			testArr := make([]int, len(tt.array))
			copy(testArr, tt.array)
			SelectionSortRecursive(testArr, 0, 0)
			diff := cmp.Diff(testArr, tt.expected)
			if diff != "" {
				t.Fatalf("SelectionSortRecursive, fail to sort: %s", diff)
			}
		})
	}
}

func TestBubbleSortIterative(t *testing.T) {
	t.Parallel()
	for _, tt := range testsExt {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			testArr := make([]int, len(tt.array))
			copy(testArr, tt.array)
			BubbleSortIterative(testArr)
			diff := cmp.Diff(testArr, tt.expected)
			if diff != "" {
				t.Fatalf("BubbleSortIterative, fail to sort: %s", diff)
			}
		})
	}
}

func TestBubbleSortRecursive(t *testing.T) {
	t.Parallel()
	for _, tt := range testsExt {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			testArr := make([]int, len(tt.array))
			copy(testArr, tt.array)
			BubbleSortRecursive(testArr)
			diff := cmp.Diff(testArr, tt.expected)
			if diff != "" {
				t.Fatalf("BubbleSortRecursive, fail to sort: %s", diff)
			}
		})
	}
}

func TestInsertionSortIterative(t *testing.T) {
	t.Parallel()
	for _, tt := range testsExt {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			testArr := make([]int, len(tt.array))
			copy(testArr, tt.array)
			InsertionSortIterative(testArr)
			diff := cmp.Diff(testArr, tt.expected)
			if diff != "" {
				t.Fatalf("InsertionSortIterative, fail to sort: %s", diff)
			}
		})
	}
}

func TestInsertionSortRecursive(t *testing.T) {
	t.Parallel()
	for _, tt := range testsExt {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			testArr := make([]int, len(tt.array))
			copy(testArr, tt.array)
			InsertionSortRecursive(testArr, 0)
			diff := cmp.Diff(testArr, tt.expected)
			if diff != "" {
				t.Fatalf("InsertionSortRecursive, fail to sort: %s", diff)
			}
		})
	}
}

func TestMergeSortRecursive(t *testing.T) {
	t.Parallel()
	for _, tt := range testsExt {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			testArr := make([]int, len(tt.array))
			copy(testArr, tt.array)
			got := MergeSortRecursive(testArr)
			diff := cmp.Diff(got, tt.expected)
			if diff != "" {
				t.Fatalf("MergeSortRecursive, fail to sort: %s", diff)
			}
		})
	}
}

func TestQuickSortRecursive(t *testing.T) {
	t.Parallel()
	for _, tt := range testsExt {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			testArr := make([]int, len(tt.array))
			copy(testArr, tt.array)
			got := QuickSortRecursive(testArr)
			diff := cmp.Diff(got, tt.expected)
			if diff != "" {
				t.Fatalf("QuickSortRecursive, fail to sort: %s", diff)
			}
		})
	}
}

func TestHeapSortByTypeMax(t *testing.T) {
	array := []int{1, 0, 5, 6, 3, 9, 7, 8, 2, 4}
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	HeapSortByType(array, MAX)
	diff := cmp.Diff(array, expected)
	if diff != "" {
		t.Fatalf("HeapSortRecursive, fail to sort: %s", diff)
	}
}

func TestHeapSortByTypeMin(t *testing.T) {
	array := []int{1, 0, 5, 6, 3, 9, 7, 8, 2, 4}
	expected := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	HeapSortByType(array, MIN)
	diff := cmp.Diff(array, expected)
	if diff != "" {
		t.Fatalf("HeapSortRecursive, fail to sort: %s", diff)
	}
}
