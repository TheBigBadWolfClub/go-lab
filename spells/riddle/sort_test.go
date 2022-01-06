package riddles

import (
	"reflect"
	"testing"
)

type tableTest struct {
	name       string
	initialArr []int
	want       []int
}

func testTableSetup() []tableTest {
	return []tableTest{
		{
			name:       "first out of order",
			initialArr: []int{5, 1, 2, 3, 4},
			want:       []int{1, 2, 3, 4, 5},
		}, {
			name:       "last out of order",
			initialArr: []int{2, 3, 4, 5, 1},
			want:       []int{1, 2, 3, 4, 5},
		}, {
			name:       "sort full inverted",
			initialArr: []int{4, 3, 2, 1},
			want:       []int{1, 2, 3, 4},
		}, {
			name:       "sort half-end inverted",
			initialArr: []int{4, 2, 3, 1},
			want:       []int{1, 2, 3, 4},
		}, {
			name:       "sort half-middle inverted",
			initialArr: []int{1, 3, 2, 4},
			want:       []int{1, 2, 3, 4},
		},
	}
}

func TestBubbleSort(t *testing.T) {

	for _, tt := range testTableSetup() {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.initialArr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {

	for _, tt := range testTableSetup() {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectionSort(tt.initialArr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {

	for _, tt := range testTableSetup() {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertionSort(tt.initialArr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {

	for _, tt := range testTableSetup() {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.initialArr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
