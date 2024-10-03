package internal

import (
	"cmp"
	"fmt"
	"iter"
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	tests := []struct {
		name string
		end  int
		want []int
	}{
		{
			name: "Counter",
			end:  5,
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Counter",
			end:  10,
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res := []int{}
			for v := range Counter(tt.end) {
				res = append(res, v)
			}

			assert.Equal(t, tt.want, res)
		})
	}
}

func TestEvenNumbers(t *testing.T) {
	tests := []struct {
		name string
		end  int
		want []int
	}{
		{
			name: "EvenNumbers",
			end:  10,
			want: []int{1, 3, 5, 7, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := []int{}
			for v := range EvenNumbers(tt.end) {
				res = append(res, v)
			}
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name  string
		terms int
		want  []int
	}{
		{
			name:  "Fibonacci",
			terms: 10,
			want:  []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := []int{}
			for v := range Fibonacci(tt.terms) {
				res = append(res, v)
			}
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestTokenize(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{
			name: "Tokenize",
			s:    "Hello, world!",
			want: []string{"Hello,", "world!"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := []string{}

			for v := range Tokenize(tt.s) {
				res = append(res, v)
			}

			assert.Equal(t, tt.want, res)
		})
	}
}

func TestReadLines(t *testing.T) {
	tests := []struct {
		name    string
		want    []string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "assets/text.txt",
			want:    []string{"line 1", "line 2", "line 3", "line 4", "line 5", "line 6", "line 7", "line 8", "line 9", "line 10"},
			wantErr: assert.NoError,
		},
		{
			name:    "no file",
			want:    []string{},
			wantErr: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := []string{}

			var err2 error
			for v, err := range ReadLines(tt.name) {
				if err != nil {
					err2 = err
					res = []string{}
					break
				}
				res = append(res, v)
			}

			if !tt.wantErr(t, err2, fmt.Sprintf("error should be nil for %s", tt.name)) {
				return
			}

			assert.Equal(t, tt.want, res)
		})
	}
}

func TestMapIterator(t *testing.T) {
	tests := []struct {
		name     string
		m        map[string]int
		wantKeys []string
		wantVals []int
	}{
		{
			name:     "MapIterator",
			m:        map[string]int{"a": 1, "b": 2, "c": 3},
			wantKeys: []string{"a", "b", "c"},
			wantVals: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vals := []int{}
			keys := []string{}
			for k, v := range MapIterator(tt.m) {
				vals = append(vals, v)
				keys = append(keys, k)
			}
			assert.ElementsMatch(t, tt.wantKeys, keys)
			assert.ElementsMatch(t, tt.wantVals, vals)
		})
	}
}

func TestMapValues(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]int
		want []int
	}{
		{
			name: "MapValues",
			m:    map[string]int{"a": 1, "b": 2, "c": 3},
			want: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := []int{}
			for v := range MapValues(tt.m) {
				res = append(res, v)
			}
			assert.ElementsMatch(t, tt.want, res)
		})
	}
}

func TestMapKeys(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]int
		want []string
	}{
		{
			name: "MapKeys",
			m:    map[string]int{"a": 1, "b": 2, "c": 3},
			want: []string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			res := []string{}
			for v := range MapKeys(tt.m) {
				res = append(res, v)
			}
			assert.ElementsMatch(t, tt.want, res)
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name      string
		predicate func(int) bool
		slice     []int
		want      []int
	}{

		{
			name: "filter pair values",
			predicate: func(v int) bool {
				return v%2 == 0
			},
			slice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:  []int{2, 4, 6, 8, 10},
		},
		{
			name: "filter odd values",
			predicate: func(v int) bool {
				return v%2 != 0
			},
			slice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:  []int{1, 3, 5, 7, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := []int{}

			for v := range Filter(tt.slice, tt.predicate) {
				res = append(res, v)
			}

			assert.Equal(t, tt.want, res)
		})
	}
}

func TestBackwards(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{
			name:  "Backwards",
			slice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:  []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := []int{}
			for v := range Backwards(tt.slice) {
				res = append(res, v)
			}
			assert.Equal(t, tt.want, res)
		})
	}
}

// -----------------------------------
// tests for package slices/iter.go
// -----------------------------------

func TestSlicesAll(t *testing.T) {

	testSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	t.Run("slices.All", func(t *testing.T) {

		var got []int
		for i, v := range slices.All(testSlice) {
			got = append(got, v)
			fmt.Printf("%d:%d ", i, v)
		}

		assert.Equal(t, testSlice, got)
	})
}

func TestSlicesBackward(t *testing.T) {
	testSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	t.Run("slices.Backward", func(t *testing.T) {
		want := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
		var got []int
		for i, v := range slices.Backward(testSlice) {
			got = append(got, v)
			fmt.Printf("%d:%d ", i, v)
		}
		assert.Equal(t, want, got)
	})
}

func TestSlicesValues(t *testing.T) {
	testSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	t.Run("slices.Values", func(t *testing.T) {
		want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		var got []int
		for v := range slices.Values(testSlice) {
			got = append(got, v)
			fmt.Printf("%d ", v)
		}
		assert.Equal(t, want, got)
	})
}

func TestSlicesAppendSeq(t *testing.T) {
	testSlice := []int{0, 1, 2, 3, 4, 5}
	appendSeq := slices.Values([]int{6, 7, 8, 9, 10})

	t.Run("slices.AppendSeq", func(t *testing.T) {
		want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got := slices.AppendSeq(testSlice, appendSeq)
		assert.Equal(t, want, got)
	})
}

func TestSlicesCollect(t *testing.T) {
	testSlice := slices.Values([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	t.Run("slices.Collect", func(t *testing.T) {
		want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got := slices.Collect(testSlice)
		assert.Equal(t, want, got)
	})
}

func TestSlicesSorted(t *testing.T) {
	testSlice := slices.Values([]int{0, 1, 4, 3, 2, 5, 7, 6, 8, 9, 10})
	t.Run("slices.Sorted", func(t *testing.T) {
		want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got := slices.Sorted(testSlice)
		assert.Equal(t, want, got)
	})
}

func TestSlicesSortedFunc(t *testing.T) {
	testSlice := slices.Values([]int{0, 1, 4, 3, 2, 5, 7, 6, 8, 9, 10})
	t.Run("slices.SortedFunc", func(t *testing.T) {

		compare := func(a, b int) int {
			return cmp.Compare(a, b) * -1
		}

		want := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
		got := slices.SortedFunc(testSlice, compare)
		assert.Equal(t, want, got)
	})
}

func TestSlicesSortedStableFunc(t *testing.T) {
	testSlice := slices.Values([]int{0, 1, 4, 3, 2, 5, 7, 6, 8, 9, 10})
	t.Run("slices.SortedStableFunc", func(t *testing.T) {
		compare := func(a, b int) int {
			return cmp.Compare(a, b) * -1
		}
		want := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
		got := slices.SortedStableFunc(testSlice, compare)
		assert.Equal(t, want, got)
	})
}

func TestSlicesChunk(t *testing.T) {
	testSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t.Run("slices.Chunk", func(t *testing.T) {
		got := [][]int{}
		for sl := range slices.Chunk(testSlice, 2) {
			got = append(got, sl)
		}

		assert.Equal(t, [][]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {8, 9}, {10}}, got)
	})

}

// -----------------------------------
// test for package maps/iter.go
// -----------------------------------

func TestMapsAll(t *testing.T) {
	t.Run("map.All", func(t *testing.T) {
		testMap := map[string]int{"a": 1, "b": 2, "c": 3}

		var gotKeys []string
		var gotValues []int
		for k, v := range maps.All(testMap) {
			gotKeys = append(gotKeys, k)
			gotValues = append(gotValues, v)
		}
		assert.ElementsMatch(t, []string{"a", "b", "c"}, gotKeys)
		assert.ElementsMatch(t, []int{1, 2, 3}, gotValues)
	})
}

func TestMapsKeys(t *testing.T) {
	t.Run("map.Keys", func(t *testing.T) {
		testMap := map[string]int{"a": 1, "b": 2, "c": 3}
		var gotKeys []string
		for key := range maps.Keys(testMap) {
			gotKeys = append(gotKeys, key)
		}
		assert.ElementsMatch(t, []string{"a", "b", "c"}, gotKeys)
	})
}

func TestMapsValues(t *testing.T) {
	t.Run("map.Values", func(t *testing.T) {
		testMap := map[string]int{"a": 1, "b": 2, "c": 3}
		var gotValues []int
		for v := range maps.Values(testMap) {
			gotValues = append(gotValues, v)
		}
		assert.Equal(t, []int{1, 2, 3}, gotValues)
	})
}

func TestMapsInsert(t *testing.T) {
	t.Run("map.Insert", func(t *testing.T) {
		testMap := map[string]int{"a": 1, "b": 2, "c": 3}
		testMapSeq := maps.All(testMap)

		gotMap := map[string]int{}
		maps.Insert(gotMap, testMapSeq)
		assert.Equal(t, testMap, gotMap)
	})
}

func TestMapsCollect(t *testing.T) {
	t.Run("map.Collect", func(t *testing.T) {
		testMap := map[string]int{"a": 1, "b": 2, "c": 3}
		testMapSeq := maps.All(testMap)
		got := maps.Collect(testMapSeq)
		assert.Equal(t, testMap, got)
	})

}

// -----------------------------------
// test pull iterators
// -----------------------------------

func TestIteratorPull(t *testing.T) {

	t.Run("iter.Pull", func(t *testing.T) {

		pushIter := slices.Values([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
		next, stop := iter.Pull(pushIter)
		defer stop()

		var got []int
		n := 5
		for {
			if n == 0 {
				break
			}
			n--

			if v, ok := next(); ok {
				got = append(got, v)
			}
		}
		assert.Equal(t, []int{0, 1, 2, 3, 4}, got)
	})

}

func TestIteratorPull2(t *testing.T) {
	t.Run("iter.Pull2", func(t *testing.T) {
		pushIter := maps.All(map[string]int{"a": 1, "b": 2, "c": 3})
		next, stop := iter.Pull2(pushIter)
		defer stop()

		var gotK []string
		var gotV []int
		for {
			k, v, ok := next()
			if !ok {
				break
			}

			gotK = append(gotK, k)
			gotV = append(gotV, v)
		}
		assert.ElementsMatch(t, []string{"a", "b", "c"}, gotK)
		assert.ElementsMatch(t, []int{1, 2, 3}, gotV)

	})

}
