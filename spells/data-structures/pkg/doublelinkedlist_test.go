package pkg

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestLinkedList_size(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		list DoubleLinkedList[rune]
		want int
	}{
		{
			name: "empty list",
			list: DoubleLinkedList[rune]{},
			want: 0,
		},
		{
			name: "one  item",
			list: linkedListGenerator(1, 0),
			want: 1,
		}, {
			name: "ten  items",
			list: linkedListGenerator(10, 0),
			want: 10,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.list.Size(); got != tt.want {
				t.Errorf("size() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}

func TestLinkedList_AddHead(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		list     DoubleLinkedList[rune]
		expected DoubleLinkedList[rune]
		value    rune
	}{
		{
			name:     "empty list",
			list:     DoubleLinkedList[rune]{},
			expected: linkedListGenerator(1, 0),
			value:    rune(0),
		}, {
			name:     "not empty list",
			list:     linkedListGenerator(1, 1),
			expected: linkedListGenerator(2, 0),
			value:    rune(0),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.list.AddHead(tt.value)

			diff := cmp.Diff(tt.list, tt.expected, cmp.AllowUnexported(DoubleLinkedList[rune]{}, DoubleLinkedNode[rune]{}))
			if diff != "" {
				t.Errorf("fail to add to list: %s", diff)
			}
		})
	}
}

func TestLinkedList_AddTail(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		list     DoubleLinkedList[rune]
		expected DoubleLinkedList[rune]
		value    rune
	}{
		{
			name:     "empty list",
			list:     DoubleLinkedList[rune]{},
			expected: linkedListGenerator(1, 0),
			value:    rune(0),
		}, {
			name:     "not empty list",
			list:     linkedListGenerator(1, 0),
			expected: linkedListGenerator(2, 0),
			value:    rune(1),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.list.AddTail(tt.value)
			diff := cmp.Diff(tt.list, tt.expected, cmp.AllowUnexported(DoubleLinkedList[rune]{}, DoubleLinkedNode[rune]{}))
			if diff != "" {
				t.Errorf("fail to add to list: %s", diff)
			}
		})
	}
}

func TestLinkedList_Add(t *testing.T) {

	insertIndex := 1
	t.Parallel()
	tests := []struct {
		name     string
		list     DoubleLinkedList[rune]
		expected DoubleLinkedList[rune]
		value    rune
		wantErr  bool
	}{
		{
			name:     "empty list",
			list:     DoubleLinkedList[rune]{},
			expected: DoubleLinkedList[rune]{},
			value:    1,
			wantErr:  true,
		}, {
			name:     "one element",
			list:     linkedListGenerator(1, 0),
			expected: linkedListGenerator(1, 0),
			value:    1,
			wantErr:  true,
		}, {
			name:     "insert middle",
			list:     linkedListGenerator(2, 0, 1),
			expected: linkedListGenerator(3, 0),
			value:    1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.list.Add(tt.value, insertIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}

			diff := cmp.Diff(tt.list, tt.expected, cmp.AllowUnexported(DoubleLinkedList[rune]{}, DoubleLinkedNode[rune]{}))
			if diff != "" && !tt.wantErr {
				t.Errorf("fail to add to list: %s", diff)
			}

		})
	}
}

func TestLinkedList_Delete(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		list     DoubleLinkedList[rune]
		expected DoubleLinkedList[rune]
		index    int
		wantErr  bool
	}{
		{
			name:     "empty list",
			list:     DoubleLinkedList[rune]{},
			expected: DoubleLinkedList[rune]{},
			index:    0,
			wantErr:  true,
		}, {
			name:     "index out of bounds",
			list:     linkedListGenerator(10, 0),
			expected: linkedListGenerator(10, 0),
			index:    10,
			wantErr:  true,
		}, {
			name:     "one elem list",
			list:     linkedListGenerator(1, 0),
			expected: DoubleLinkedList[rune]{},
			index:    0,
			wantErr:  false,
		}, {
			name:     "remove tail",
			list:     linkedListGenerator(2, 0),
			expected: linkedListGenerator(1, 0),
			index:    1,
			wantErr:  false,
		}, {
			name:     "remove head",
			list:     linkedListGenerator(2, 0),
			expected: linkedListGenerator(1, 1),
			index:    0,
			wantErr:  false,
		}, {
			name:     "remove middle",
			list:     linkedListGenerator(3, 0),
			expected: linkedListGenerator(2, 0, 1),
			index:    1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.list.Delete(tt.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}

			diff := cmp.Diff(tt.list, tt.expected, cmp.AllowUnexported(DoubleLinkedList[rune]{}, DoubleLinkedNode[rune]{}))
			if diff != "" && !tt.wantErr {
				t.Errorf("fail to delete at index=%d list: %s", tt.index, diff)
			}

		})
	}
}

func TestLinkedList_DeleteHead(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		list     DoubleLinkedList[rune]
		expected DoubleLinkedList[rune]
		wantErr  bool
	}{
		{
			name:     "empty list",
			list:     DoubleLinkedList[rune]{},
			expected: DoubleLinkedList[rune]{},
			wantErr:  true,
		}, {
			name:     "one elem list",
			list:     linkedListGenerator(1, 0),
			expected: DoubleLinkedList[rune]{},
			wantErr:  false,
		}, {
			name:     "two elem list",
			list:     linkedListGenerator(2, 0),
			expected: linkedListGenerator(1, 1),
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.list.DeleteHead()
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteHead() error = %v, wantErr %v", err, tt.wantErr)
			}

			diff := cmp.Diff(tt.list, tt.expected, cmp.AllowUnexported(DoubleLinkedList[rune]{}, DoubleLinkedNode[rune]{}))
			if diff != "" && !tt.wantErr {
				t.Errorf("fail to delete tail list: %s", diff)
			}

		})
	}
}

func TestLinkedList_DeleteTail(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		list     DoubleLinkedList[rune]
		expected DoubleLinkedList[rune]
		wantErr  bool
	}{
		{
			name:     "empty list",
			list:     DoubleLinkedList[rune]{},
			expected: DoubleLinkedList[rune]{},
			wantErr:  true,
		}, {
			name:     "one elem list",
			list:     linkedListGenerator(1, 0),
			expected: DoubleLinkedList[rune]{},
			wantErr:  false,
		}, {
			name:     "two elem list",
			list:     linkedListGenerator(2, 0),
			expected: linkedListGenerator(1, 0),
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.list.DeleteTail()
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteTail() error = %v, wantErr %v", err, tt.wantErr)
			}

			diff := cmp.Diff(tt.list, tt.expected, cmp.AllowUnexported(DoubleLinkedList[rune]{}, DoubleLinkedNode[rune]{}))
			if diff != "" && !tt.wantErr {
				t.Errorf("fail to delete tail list: %s", diff)
			}

		})
	}
}

func linkedListGenerator(size, start int, skip ...rune) DoubleLinkedList[rune] {

	skipIt := func(index rune) bool {
		for _, sk := range skip {
			if sk == index {
				return true
			}
		}
		return false
	}

	values := func() []int {
		var vs []int
		nextValue := start
		for i := 0; i < size; i++ {
			if skipIt(rune(i + start)) {
				nextValue++
			}
			vs = append(vs, nextValue)
			nextValue++
		}
		return vs
	}

	nodes := make([]DoubleLinkedNode[rune], size)
	ints := values()
	for ix, vs := range ints {
		nodes[ix].valid = true
		nodes[ix].value = rune(vs)

		if ix > 0 {
			nodes[ix].prev = &nodes[ix-1]
		}
		if ix < size-1 {
			nodes[ix].next = &nodes[ix+1]
		}
	}

	return DoubleLinkedList[rune]{
		head:  &nodes[0],
		tail:  &nodes[size-1],
		count: size,
	}

}
