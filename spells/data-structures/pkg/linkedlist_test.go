package pkg

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestSingleLinkList_Add(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		value    int
		list     LinkedList[int]
		expected LinkedList[int]
	}{
		{
			name:  "empty list",
			value: 0,
			list:  LinkedList[int]{},
			expected: LinkedList[int]{
				Len: 1,
				Head: &linkedListNode[int]{
					next:  nil,
					value: 0,
				},
			},
		}, {
			name:  "not empty list",
			value: 1,
			list: LinkedList[int]{
				Len: 1,
				Head: &linkedListNode[int]{
					next:  nil,
					value: 0,
				},
			},
			expected: LinkedList[int]{
				Len: 2,
				Head: &linkedListNode[int]{
					next: &linkedListNode[int]{
						next:  nil,
						value: 1,
					},
					value: 0,
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.list.Add(tt.value)

			diff := cmp.Diff(tt.list, tt.expected, cmp.AllowUnexported(linkedListNode[int]{}, linkedListNode[int]{}))
			if diff != "" {
				t.Errorf("fail to treeNodeAdd to list: %s", diff)
			}
		})
	}
}

func TestSingleLinkList_Get(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		index int
		list  *LinkedList[int]
		value int
		ok    bool
	}{
		{
			name:  "index 0",
			index: 0,
			list:  buildTestList(0),
			value: 0,
			ok:    true,
		}, {
			name:  "index 2",
			index: 2,
			list:  buildTestList(0, 1, 2),
			value: 2,
			ok:    true,
		}, {
			name:  "index not exist",
			index: 3,
			list:  buildTestList(0, 1, 2),
			value: 0,
			ok:    false,
		}, {
			name:  "index on empty list",
			index: 0,
			list:  &LinkedList[int]{},
			value: 0,
			ok:    false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, ok := tt.list.Get(tt.index)
			if !cmp.Equal(got, tt.value) {
				t.Errorf("Get() value got = %v, want %v", got, tt.value)
			}
			if ok != tt.ok {
				t.Errorf("Get() ok got1 = %v, want %v", ok, tt.ok)
			}
		})
	}
}

func TestSingleLinkList_Delete(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		index    int
		want     error
		list     *LinkedList[int]
		expected *LinkedList[int]
	}{
		{
			name:     "empty list",
			index:    1,
			want:     IndexNotFound,
			list:     &LinkedList[int]{},
			expected: &LinkedList[int]{},
		}, {
			name:     "DeleteIndex  index 0",
			index:    0,
			want:     nil,
			list:     buildTestList(0, 1, 2),
			expected: buildTestList(1, 2),
		}, {
			name:     "DeleteIndex  index 1",
			index:    1,
			want:     nil,
			list:     buildTestList(0, 1, 2),
			expected: buildTestList(0, 2),
		}, {
			name:     "DeleteIndex  last index ",
			index:    2,
			want:     nil,
			list:     buildTestList(0, 1, 2),
			expected: buildTestList(0, 1),
		}, {
			name:     "index bigger than list size",
			index:    10,
			want:     IndexNotFound,
			list:     buildTestList(0, 1, 2),
			expected: buildTestList(0, 1, 2),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.list.Delete(tt.index)
			if err != tt.want {
				t.Errorf("Delete() got = %v, want %v", err, tt.want)

			}

			diff := cmp.Diff(tt.list, tt.expected, cmp.AllowUnexported(linkedListNode[int]{}, linkedListNode[int]{}))
			if diff != "" {
				t.Errorf("unexpect list after DeleteIndex: %s", diff)
			}
		})
	}
}

func TestSingleLinkList_IsEmpty(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		list *LinkedList[int]
		want bool
	}{
		{
			name: "is empty",
			list: &LinkedList[int]{},
			want: true,
		}, {
			name: "not empty",
			list: buildTestList(0, 1),
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.list.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingleLinkList_Size(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		list *LinkedList[int]
		want int
	}{
		{
			name: "empty list",
			list: &LinkedList[int]{},
			want: 0,
		}, {
			name: "size one",
			list: buildTestList(0),
			want: 1,
		}, {
			name: "size two",
			list: buildTestList(0, 1),
			want: 2,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.list.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingleLinkList_tail(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		list *LinkedList[int]
		want *linkedListNode[int]
	}{
		{
			name: "empty list",
			list: &LinkedList[int]{},
			want: nil,
		}, {
			name: "one item",
			list: buildTestList(0),
			want: &linkedListNode[int]{
				next:  nil,
				value: 0,
			},
		}, {
			name: "three item",
			list: buildTestList(0, 1, 2),
			want: &linkedListNode[int]{
				next:  nil,
				value: 2,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tail := tt.list.Tail()
			diff := cmp.Diff(tail, tt.want, cmp.AllowUnexported(linkedListNode[int]{}))
			if diff != "" {
				t.Errorf("unexpect list after DeleteIndex: %s", diff)
			}
		})
	}
}

func TestSingleLinkList_deleteHead(t *testing.T) {

	t.Parallel()

	tests := []struct {
		name     string
		list     *LinkedList[int]
		expected *LinkedList[int]
		wantOk   bool
	}{
		{
			name:     "empty list",
			list:     &LinkedList[int]{},
			expected: &LinkedList[int]{},
			wantOk:   true,
		}, {
			name:     "head is last node",
			list:     buildTestList(0),
			expected: &LinkedList[int]{},
			wantOk:   true,
		}, {
			name:     "link head to next node",
			list:     buildTestList(0, 1, 2),
			expected: buildTestList(1, 2),
			wantOk:   true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if gotOk := tt.list.DeleteHead(); gotOk != tt.wantOk {
				t.Errorf("DeleteHead() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestSingleLinkList_Reverse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		list     *LinkedList[int]
		expected *LinkedList[int]
	}{
		{
			name:     "reverse empty",
			list:     &LinkedList[int]{},
			expected: &LinkedList[int]{},
		}, {
			name:     "one elem list",
			list:     buildTestList(1),
			expected: buildTestList(1),
		}, {
			name:     "two elem list",
			list:     buildTestList(1, 2),
			expected: buildTestList(2, 1),
		}, {
			name:     "3 elem list",
			list:     buildTestList(0, 1, 2),
			expected: buildTestList(2, 1, 0),
		}, {
			name:     "10 elem list",
			list:     buildTestList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
			expected: buildTestList(9, 8, 7, 6, 5, 4, 3, 2, 1, 0),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.list.Reverse()
			fmt.Println(tt.list)
			diff := cmp.Diff(tt.list, tt.expected, cmp.AllowUnexported(linkedListNode[int]{}, linkedListNode[int]{}))
			if diff != "" {
				t.Errorf("fail to reverse list: %s", diff)
			}
		})
	}
}

func TestSingleLinkList_ReverseRecursive(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		list     *LinkedList[int]
		expected *LinkedList[int]
	}{
		{
			name:     "reverse empty",
			list:     &LinkedList[int]{},
			expected: &LinkedList[int]{},
		}, {
			name:     "one elem list",
			list:     buildTestList(1),
			expected: buildTestList(1),
		}, {
			name:     "two elem list",
			list:     buildTestList(1, 2),
			expected: buildTestList(2, 1),
		}, {
			name:     "3 elem list",
			list:     buildTestList(0, 1, 2),
			expected: buildTestList(2, 1, 0),
		}, {
			name:     "10 elem list",
			list:     buildTestList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
			expected: buildTestList(9, 8, 7, 6, 5, 4, 3, 2, 1, 0),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.list.ReverseRecursive()
			fmt.Println(tt.list)
			diff := cmp.Diff(tt.list, tt.expected, cmp.AllowUnexported(linkedListNode[int]{}, linkedListNode[int]{}))
			if diff != "" {
				t.Errorf("fail to reverse list: %s", diff)
			}
		})
	}
}

func TestSingleLinkList_Traverse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		list     *LinkedList[int]
		expected []int
	}{
		{
			name:     "reverse empty",
			list:     &LinkedList[int]{},
			expected: nil,
		}, {
			name:     "one elem list",
			list:     buildTestList(1),
			expected: []int{1},
		}, {
			name:     "two elem list",
			list:     buildTestList(1, 2),
			expected: []int{1, 2},
		}, {
			name:     "3 elem list",
			list:     buildTestList(0, 1, 2),
			expected: []int{0, 1, 2},
		}, {
			name:     "10 elem list",
			list:     buildTestList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			traverse := tt.list.Traverse()
			diff := cmp.Diff(traverse, tt.expected)
			if diff != "" {
				t.Errorf("Traverse: %s", diff)
			}

			recursive := tt.list.TraverseRecursive()
			diff = cmp.Diff(recursive, tt.expected)
			if diff != "" {
				t.Errorf("TraverseRecursive: %s", diff)
			}
		})
	}
}

func buildTestList(values ...int) *LinkedList[int] {
	var list LinkedList[int]
	for _, v := range values {
		list.Add(v)
	}

	return &list
}
