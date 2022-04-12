package pkg

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

var testsExt = []struct {
	name     string
	array    SortedSlice[int]
	expected SortedSlice[int]
}{
	{
		name:     "sort empty",
		array:    SortedSlice[int]{},
		expected: SortedSlice[int]{},
	}, {
		name:     "sort one elem",
		array:    SortedSlice[int]{0},
		expected: SortedSlice[int]{0},
	}, {
		name:     "sort 2 elem",
		array:    SortedSlice[int]{1, 0},
		expected: SortedSlice[int]{0, 1},
	}, {
		name:     "sort multiple elem",
		array:    SortedSlice[int]{1, 0, 5, 6, 3, 9, 7, 8, 2, 4},
		expected: SortedSlice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
}

func TestSelectionSortIterative(t *testing.T) {
	t.Parallel()
	for _, tt := range testsExt {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			testArr := make(SortedSlice[int], len(([]int)(tt.array)))
			copy(testArr, ([]int)(tt.array))
			testArr.SelectionSortIterative()
			diff := cmp.Diff(testArr, tt.expected)
			if diff != "" {
				t.Fatalf("SelectionSortIterative, fail to sort: %s", diff)
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
			testArr := make(SortedSlice[int], len(([]int)(tt.array)))
			copy(testArr, ([]int)(tt.array))
			testArr.BubbleSortIterative()
			diff := cmp.Diff(testArr, tt.expected)
			if diff != "" {
				t.Fatalf("BubbleSortIterative, fail to sort: %s", diff)
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
			testArr := make(SortedSlice[int], len(([]int)(tt.array)))
			copy(testArr, ([]int)(tt.array))
			testArr.InsertionSortIterative()
			diff := cmp.Diff(testArr, tt.expected)
			if diff != "" {
				t.Fatalf("InsertionSortIterative, fail to sort: %s", diff)
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
			testArr := make(SortedSlice[int], len(([]int)(tt.array)))
			copy(testArr, ([]int)(tt.array))
			final := testArr.MergeSortRecursive()
			diff := cmp.Diff(final, tt.expected)
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
			testArr := make(SortedSlice[int], len(([]int)(tt.array)))
			copy(testArr, ([]int)(tt.array))
			final := testArr.QuickSortRecursive()
			diff := cmp.Diff(final, tt.expected)
			if diff != "" {
				t.Fatalf("QuickSortRecursive, fail to sort: %s", diff)
			}
		})
	}
}

func TestHeapSortByTypeMax(t *testing.T) {
	array := SortedSlice[int]{1, 0, 5, 6, 3, 9, 7, 8, 2, 4}
	expected := SortedSlice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	array.HeapSortByType(MAX)
	diff := cmp.Diff(array, expected)
	if diff != "" {
		t.Fatalf("HeapSortRecursive, fail to sort: %s", diff)
	}
}

func TestHeapSortByTypeMin(t *testing.T) {
	array := SortedSlice[int]{1, 0, 5, 6, 3, 9, 7, 8, 2, 4}
	expected := SortedSlice[int]{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	array.HeapSortByType(MIN)
	diff := cmp.Diff(array, expected)
	if diff != "" {
		t.Fatalf("HeapSortRecursive, fail to sort: %s", diff)
	}
}
