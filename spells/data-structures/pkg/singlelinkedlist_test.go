package pkg

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestSingleLinkList(t *testing.T) {
	l := SingleLinkList[int]{}

	if l.next != nil || l.valid {
		t.Errorf("expected empty list")
	}

	l.Add(1)
	if !l.valid {
		t.Errorf("expected first elem to be valid")
	}
}

func TestSingleLinkList_Add(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		value    int
		list     SingleLinkList[int]
		expected SingleLinkList[int]
	}{
		{
			name:  "empty list",
			value: 0,
			list:  SingleLinkList[int]{},
			expected: SingleLinkList[int]{
				next:  nil,
				valid: true,
				value: 0,
			},
		}, {
			name:  "not empty list",
			value: 1,
			list: SingleLinkList[int]{
				next:  nil,
				valid: true,
				value: 0,
			},
			expected: SingleLinkList[int]{
				next: &SingleLinkList[int]{
					next:  nil,
					valid: true,
					value: 1,
				},
				valid: true,
				value: 0,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.list.Add(tt.value)

			diff := cmp.Diff(tt.list, tt.expected, cmp.AllowUnexported(SingleLinkList[int]{}, SingleLinkList[int]{}))
			if diff != "" {
				t.Errorf("fail to add to list: %s", diff)
			}
		})
	}
}

func TestSingleLinkList_Get(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		index int
		list  SingleLinkList[int]
		value int
		ok    bool
	}{
		{
			name:  "index 0",
			index: 0,
			list:  testList(),
			value: 0,
			ok:    true,
		}, {
			name:  "index 2",
			index: 2,
			list:  testList(),
			value: 2,
			ok:    true,
		}, {
			name:  "index not exist",
			index: 3,
			list:  testList(),
			value: 0,
			ok:    false,
		}, {
			name:  "index on empty list",
			index: 0,
			list:  SingleLinkList[int]{},
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
		list     SingleLinkList[int]
		expected SingleLinkList[int]
	}{
		{
			name:     "empty list",
			index:    1,
			want:     IndexNotFound,
			list:     SingleLinkList[int]{},
			expected: SingleLinkList[int]{},
		}, {
			name:  "delete  index 0",
			index: 0,
			want:  nil,
			list:  testList(),
			expected: SingleLinkList[int]{
				next: &SingleLinkList[int]{
					next:  nil,
					valid: true,
					value: 2,
				},
				valid: true,
				value: 1,
			},
		}, {
			name:  "delete  index 1",
			index: 1,
			want:  nil,
			list:  testList(),
			expected: SingleLinkList[int]{
				next: &SingleLinkList[int]{
					next:  nil,
					valid: true,
					value: 2,
				},
				valid: true,
				value: 0,
			},
		}, {
			name:  "delete  last index ",
			index: 2,
			want:  nil,
			list:  testList(),
			expected: SingleLinkList[int]{
				next: &SingleLinkList[int]{
					next:  nil,
					valid: true,
					value: 1,
				},
				valid: true,
				value: 0,
			},
		}, {
			name:     "index bigger than list size",
			index:    10,
			want:     IndexNotFound,
			list:     testList(),
			expected: testList(),
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

			diff := cmp.Diff(tt.list, tt.expected, cmp.AllowUnexported(SingleLinkList[int]{}, SingleLinkList[int]{}))
			if diff != "" {
				t.Errorf("unexpect list after delete: %s", diff)
			}
		})
	}
}

func TestSingleLinkList_IsEmpty(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		list SingleLinkList[string]
		want bool
	}{
		{
			name: "is empty",
			list: SingleLinkList[string]{
				next:  nil,
				valid: false,
				value: "nil",
			},
			want: true,
		}, {
			name: "not empty",
			list: SingleLinkList[string]{
				next:  nil,
				valid: true,
				value: "a string",
			},
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
		list SingleLinkList[int]
		want int
	}{
		{
			name: "empty list",
			list: SingleLinkList[int]{},
			want: 0,
		}, {
			name: "size one",
			list: SingleLinkList[int]{
				next:  nil,
				valid: true,
				value: 0,
			},
			want: 1,
		}, {
			name: "size two",
			list: SingleLinkList[int]{
				next: &SingleLinkList[int]{
					next:  nil,
					valid: true,
					value: 1,
				},
				valid: true,
				value: 0,
			},
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
	tail := &SingleLinkList[int]{
		next:  nil,
		valid: true,
		value: 1000,
	}
	tests := []struct {
		name string
		list SingleLinkList[int]
		want *SingleLinkList[int]
	}{
		{
			name: "empty list",
			list: SingleLinkList[int]{},
			want: nil,
		}, {
			name: "one item",
			list: SingleLinkList[int]{
				next:  tail.next,
				valid: tail.valid,
				value: tail.value,
			},
			want: tail,
		}, {
			name: "two item",
			list: SingleLinkList[int]{
				next:  tail,
				valid: true,
				value: 0,
			},
			want: tail,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tail := tt.list.tail()
			diff := cmp.Diff(tail, tt.want, cmp.AllowUnexported(SingleLinkList[int]{}))
			if diff != "" {
				t.Errorf("unexpect list after delete: %s", diff)
			}
		})
	}
}

func TestSingleLinkList_deleteHead(t *testing.T) {

	t.Parallel()

	tests := []struct {
		name     string
		list     SingleLinkList[int]
		expected SingleLinkList[int]
		index    int
		wantOk   bool
	}{
		{
			name:     "empty list",
			list:     SingleLinkList[int]{},
			expected: SingleLinkList[int]{},
			index:    0,
			wantOk:   true,
		}, {
			name: "head is last node",
			list: SingleLinkList[int]{
				next:  nil,
				valid: true,
				value: 1,
			},
			expected: SingleLinkList[int]{
				next:  nil,
				valid: false,
				value: 0,
			},
			index:  0,
			wantOk: true,
		}, {
			name: "link head to next node",
			list: testList(),
			expected: SingleLinkList[int]{
				next:  &SingleLinkList[int]{},
				valid: false,
				value: 1,
			},
			index:  0,
			wantOk: true,
		}, {
			name:     "not head index",
			list:     SingleLinkList[int]{},
			expected: SingleLinkList[int]{},
			index:    1,
			wantOk:   false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if gotOk := tt.list.deleteHead(tt.index); gotOk != tt.wantOk {
				t.Errorf("deleteHead() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func testList() SingleLinkList[int] {
	return SingleLinkList[int]{
		next: &SingleLinkList[int]{
			next: &SingleLinkList[int]{
				next:  nil,
				valid: true,
				value: 2,
			},
			valid: true,
			value: 1,
		},
		valid: true,
		value: 0,
	}
}
