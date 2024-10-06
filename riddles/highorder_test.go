package riddles

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestFilter(t *testing.T) {

	type testCase[T comparable] struct {
		name   string
		arr    []T
		filter func(T) bool
		want   []T
	}
	tests := []testCase[int]{
		{
			name: "Filter even numbers",
			arr:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			filter: func(i int) bool {
				return i%2 == 0
			},
			want: []int{2, 4, 6, 8},
		},
		{
			name: "Prime numbers",
			arr:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			filter: func(i int) bool {
				if i == 1 {
					return false
				}
				for j := 2; j < i; j++ {
					if i%j == 0 {
						return false
					}
				}
				return true
			},
			want: []int{2, 3, 5, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.arr, tt.filter)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMap(t *testing.T) {
	type testCase[T comparable, Z comparable] struct {
		name  string
		arr   []T
		mapFn func(T) Z
		want  []Z
	}
	tests := []testCase[string, string]{
		{
			name: "Map char to ASCII value",
			arr:  []string{"1", "2", "3", "4", "5"},
			mapFn: func(i string) string {
				return strconv.Itoa(int(i[0]))
			},
			want: []string{"49", "50", "51", "52", "53"},
		},
		{
			name: "decimal to hex code",
			arr:  []string{"1", "10", "100", "1000", "10000"},
			mapFn: func(i string) string {
				str, _ := strconv.Atoi(i)
				return strconv.FormatInt(int64(str), 16)
			},
			want: []string{"1", "a", "64", "3e8", "2710"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Map(tt.arr, tt.mapFn)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestReduce(t *testing.T) {
	type args[T comparable] struct {
		arr []T
		f   func(T, T) T
	}
	type testCase[T comparable] struct {
		name   string
		arr    []T
		reduce func(T, T) T
		want   T
	}
	tests := []testCase[int]{
		{
			name: "Sum of all elements",
			arr:  []int{1, 2, 3, 4, 5},
			reduce: func(a, b int) int {
				return a + b
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(tt.arr, tt.reduce)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAny(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		arr  []T
		any  func(T) bool
		want bool
	}
	tests := []testCase[int]{
		{
			name: "Any even number",
			arr:  []int{1, 2, 3, 4, 5},
			any: func(i int) bool {
				return i%2 == 0
			},
			want: true,
		},
		{
			name: "No prime number",
			arr:  []int{1, 4, 6, 8, 9},
			any: func(i int) bool {
				if i == 1 {
					return false
				}
				for j := 2; j < i; j++ {
					if i%j == 0 {
						return false
					}
				}
				return true
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Any(tt.arr, tt.any)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAll(t *testing.T) {
	type args[T comparable] struct {
	}
	type testCase[T comparable] struct {
		name string
		arr  []T
		all  func(T) bool
		want bool
	}
	tests := []testCase[int]{
		{
			name: "All even number",
			arr:  []int{2, 4, 6, 8, 10},
			all: func(i int) bool {
				return i%2 == 0
			},
			want: true,
		},
		{
			name: "Not all are odd number",
			arr:  []int{1, 2, 3, 4, 5},
			all: func(i int) bool {
				return i%2 != 0
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := All(tt.arr, tt.all)
			assert.Equal(t, tt.want, got)
		})
	}
}
