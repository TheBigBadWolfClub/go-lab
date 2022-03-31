package pkg

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/exp/constraints"
	"reflect"
	"testing"
)

var defaultArray []rune

func NewTree[T constraints.Ordered](values ...T) *BinaryTree[T] {
	b := &BinaryTree[T]{}
	for _, v := range values {
		b.Insert(v)
	}
	return b
}

func Test_Alpha(t *testing.T) {
	tree := BinaryTree[rune]{}
	tree.Insert('F')
	tree.Insert('D')
	tree.Insert('B')
	tree.Insert('C')
	tree.Insert('A')

	s := tree.root.String()
	fmt.Println(s)
}

func TestBinaryTree_Insert(t *testing.T) {

	tests := []struct {
		name   string
		expect *BinaryTree[int]
		values []int
	}{
		{
			name: "treeNodeAdd root",
			expect: &BinaryTree[int]{
				root: &TreeNode[int]{Value: 0},
			},
			values: []int{0},
		}, {
			name: "treeNodeAdd root and left children",
			expect: &BinaryTree[int]{
				root: &TreeNode[int]{
					Left:  &TreeNode[int]{Value: 0},
					Value: 1,
				},
			},
			values: []int{1, 0},
		}, {
			name: "treeNodeAdd root and right children",
			expect: &BinaryTree[int]{
				root: &TreeNode[int]{
					Right: &TreeNode[int]{Value: 2},
					Value: 1,
				},
			},
			values: []int{1, 2},
		}, {
			name: "treeNodeAdd root and full children",
			expect: &BinaryTree[int]{
				root: &TreeNode[int]{
					Right: &TreeNode[int]{Value: 2},
					Left:  &TreeNode[int]{Value: 0},
					Value: 1,
				},
			},
			values: []int{1, 2, 0},
		}, {
			name: "3 levels full",
			expect: &BinaryTree[int]{
				root: &TreeNode[int]{
					Right: &TreeNode[int]{
						Left:  &TreeNode[int]{Value: 6},
						Value: 7,
						Right: &TreeNode[int]{Value: 8},
					},
					Left: &TreeNode[int]{
						Left:  &TreeNode[int]{Value: 0},
						Value: 1,
						Right: &TreeNode[int]{Value: 2},
					},
					Value: 5,
				},
			},
			values: []int{5, 1, 7, 0, 2, 6, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BinaryTree[int]{}
			for _, value := range tt.values {
				tree.Insert(value)
			}
			diff := cmp.Diff(tree, tt.expect, cmp.AllowUnexported(BinaryTree[int]{}), cmp.AllowUnexported(BinaryTree[int]{}))
			if diff != "" {
				t.Errorf("fail to create tree: %s", diff)
			}
		})
	}
}

func TestBinaryTree_Exists(t *testing.T) {
	tests := []struct {
		name  string
		tree  *BinaryTree[int]
		value int
		want  bool
	}{
		{
			name: "tree root is nil",
			tree: &BinaryTree[int]{},
			want: false,
		}, {
			name:  "only root, value not match",
			tree:  NewTree(0),
			value: 100,
			want:  false,
		}, {
			name:  "only root, value  match",
			tree:  NewTree(0),
			value: 0,
			want:  true,
		}, {
			name:  "found on left",
			tree:  NewTree(1, 0, 2),
			value: 0,
			want:  true,
		}, {
			name:  "found on right",
			tree:  NewTree(1, 0, 2),
			value: 2,
			want:  true,
		}, {
			name:  "not found",
			tree:  NewTree(1, 0, 2),
			value: 3,
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tree.Exists(tt.value)
			if tt.want != got {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_Delete(t *testing.T) {
	tests := []struct {
		name     string
		tree     *BinaryTree[int]
		expected *BinaryTree[int]
		value    int
		want     bool
	}{
		{
			name:     "tree root is nil",
			tree:     &BinaryTree[int]{},
			expected: &BinaryTree[int]{},
			want:     false,
		}, {
			name:     "only root, value not match",
			tree:     NewTree(0),
			expected: NewTree(0),
			value:    100,
			want:     false,
		}, {
			name:     "only root, value  match",
			tree:     NewTree(0),
			expected: &BinaryTree[int]{},
			value:    0,
			want:     true,
		}, {
			name:     "deleteNode on left leaf",
			tree:     NewTree(1, 0, 2),
			expected: NewTree(1, 2),
			value:    0,
			want:     true,
		}, {
			name:     "found on right leaf",
			tree:     NewTree(1, 0, 2),
			expected: NewTree(1, 0),
			value:    2,
			want:     true,
		}, {
			name:     "deleteNode root  with 2 leafs",
			tree:     NewTree(1, 0, 2),
			expected: NewTree(2, 0),
			value:    1,
			want:     true,
		}, {
			name:     "deleteNode root  3 levels",
			tree:     NewTree(5, 3, 10, 4, 2, 6, 8, 7, 9),
			expected: NewTree(5, 3, 6, 4, 2, 8, 7, 9),
			value:    10,
			want:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tree.Delete(tt.value)
			if tt.want != got {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}

			diff := cmp.Diff(tt.tree, tt.expected, cmp.AllowUnexported(BinaryTree[int]{}, BinaryTree[int]{}))
			if diff != "" {
				t.Errorf("Delete() =  %s", diff)
			}
		})
	}
}

func TestBinaryTree_MaxDepth(t *testing.T) {

	tests := []struct {
		name string
		tree *BinaryTree[int]
		want int
	}{
		{
			name: "root is nil",
			tree: &BinaryTree[int]{},
			want: 0,
		}, {
			name: "only root",
			tree: NewTree(0),
			want: 1,
		}, {
			name: "left 2 levels",
			tree: NewTree(1, 0),
			want: 2,
		}, {
			name: "left 3 levels",
			tree: NewTree(2, 1, 0),
			want: 3,
		}, {
			name: "right 2 levels",
			tree: NewTree(1, 2),
			want: 2,
		}, {
			name: "right 3 levels",
			tree: NewTree(1, 2, 3),
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tree.MaxDepth()
			if got != tt.want {
				t.Errorf("MaxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_Reverse(t *testing.T) {

	tests := []struct {
		name     string
		tree     *BinaryTree[int]
		expected *BinaryTree[int]
	}{
		{
			name:     "is empty",
			tree:     &BinaryTree[int]{},
			expected: &BinaryTree[int]{},
		}, {
			name:     "only root",
			tree:     NewTree(1),
			expected: NewTree(1),
		}, {
			name: "2 level",

			tree: NewTree(1, 0, 2),
			expected: &BinaryTree[int]{
				root: &TreeNode[int]{
					Left:  &TreeNode[int]{Value: 2},
					Value: 1,
					Right: &TreeNode[int]{Value: 0},
				},
			},
		}, {
			name: "3 levels",
			tree: NewTree(5, 3, 7, 2, 4, 6, 8),
			expected: &BinaryTree[int]{
				root: &TreeNode[int]{
					Left: &TreeNode[int]{
						Left:  &TreeNode[int]{Value: 8},
						Value: 7,
						Right: &TreeNode[int]{Value: 6},
					},
					Value: 5,
					Right: &TreeNode[int]{
						Left:  &TreeNode[int]{Value: 4},
						Value: 3,
						Right: &TreeNode[int]{Value: 2},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tree.Reverse()
			diff := cmp.Diff(tt.tree, tt.expected, cmp.AllowUnexported(BinaryTree[int]{}), cmp.AllowUnexported(BinaryTree[int]{}))
			if diff != "" {
				t.Errorf("fail to reverse tree: %s", diff)
			}
		})
	}
}

func TestBinaryTree_MaxValue(t *testing.T) {

	tests := []struct {
		name string
		tree *BinaryTree[int]
		want int
		ok   bool
	}{
		{
			name: "nil root",
			tree: &BinaryTree[int]{},
			want: 0,
			ok:   false,
		}, {
			name: "root is max",
			tree: NewTree(1),
			want: 1,
			ok:   true,
		}, {
			name: "root with children, max is root",
			tree: NewTree(5, 3, 4),
			want: 5,
			ok:   true,
		}, {
			name: "root with children, children is max",
			tree: NewTree(5, 7, 9),
			want: 9,
			ok:   true,
		}, {
			name: "several levels, max is in the mid levels",
			tree: NewTree(5, 7, 10, 8, 11, 9),
			want: 11,
			ok:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, ok := tt.tree.MaxValue()
			if ok != tt.ok {
				t.Fatalf("MaxValue() ok = %v, want %v", ok, tt.ok)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxValue() got = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestBinaryTree_MinValue(t *testing.T) {

	tests := []struct {
		name string
		tree *BinaryTree[int]
		want int
		ok   bool
	}{
		{
			name: "nil root",
			tree: &BinaryTree[int]{},
			want: 0,
			ok:   false,
		}, {
			name: "root is max",
			tree: NewTree(1),
			want: 1,
			ok:   true,
		}, {
			name: "root with children, max is root",
			tree: NewTree(2, 3, 4),
			want: 2,
			ok:   true,
		}, {
			name: "root with children, children is max",
			tree: NewTree(10, 7, 9),
			want: 7,
			ok:   true,
		}, {
			name: "several levels, max is in the mid levels",
			tree: NewTree(5, 3, 2, 0, 1),
			want: 0,
			ok:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, ok := tt.tree.MinValue()
			if ok != tt.ok {
				t.Fatalf("MinValue() ok = %v, want %v", ok, tt.ok)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinValue() got = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestBinaryTree_DepthFirstSearch(t *testing.T) {

	tests := []struct {
		name string
		tree *BinaryTree[rune]
		want []rune
	}{
		{
			name: "nil root",
			tree: &BinaryTree[rune]{},
			want: defaultArray,
		}, {
			name: "root only",
			tree: NewTree('4'),
			want: []rune{'4'},
		}, {
			name: "level one full",
			tree: NewTree('4', '3', '5'),
			want: []rune{'4', '3', '5'},
		}, {
			name: "level two full",
			tree: NewTree('5', '3', '4', '2', '7', '6', '8'),
			want: []rune{'5', '3', '2', '4', '7', '6', '8'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tree.DFSPreOrder()
			fmt.Println(got)
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf("DFSPreOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_DFSInOrder(t *testing.T) {

	tests := []struct {
		name string
		tree *BinaryTree[rune]
		want []rune
	}{
		{
			name: "nil tree",
			tree: &BinaryTree[rune]{},
			want: defaultArray,
		}, {
			name: "root only",
			tree: NewTree('4'),
			want: []rune{'4'},
		}, {
			name: "level one full",
			tree: NewTree('4', '3', '5'),
			want: []rune{'3', '4', '5'},
		}, {
			name: "level two full",
			tree: NewTree('5', '3', '4', '2', '7', '6', '8'),
			want: []rune{'2', '3', '4', '5', '6', '7', '8'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tree.DFSInOrder()
			fmt.Println(got)
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf("DFSInOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_DFSPostOrder(t *testing.T) {

	tests := []struct {
		name string
		tree *BinaryTree[rune]
		want []rune
	}{
		{
			name: "nil tree",
			tree: &BinaryTree[rune]{},
			want: defaultArray,
		}, {
			name: "root only",
			tree: NewTree('4'),
			want: []rune{'4'},
		}, {
			name: "level one full",
			tree: NewTree('4', '3', '5'),
			want: []rune{'3', '5', '4'},
		}, {
			name: "level two full",
			tree: NewTree('5', '3', '4', '2', '7', '6', '8'),
			want: []rune{'2', '4', '3', '6', '8', '7', '5'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tree.DFSPostOrder()
			fmt.Println(got)
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf("DFSPostOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_BFSLevelOrder(t *testing.T) {

	tests := []struct {
		name string
		tree *BinaryTree[rune]
		want []rune
	}{
		{
			name: "nil tree",
			tree: &BinaryTree[rune]{},
			want: defaultArray,
		}, {
			name: "root only",
			tree: NewTree('4'),
			want: []rune{'4'},
		}, {
			name: "level one full",
			tree: NewTree('4', '3', '5'),
			want: []rune{'4', '3', '5'},
		}, {
			name: "level two full",
			tree: NewTree('5', '3', '4', '2', '7', '6', '8'),
			want: []rune{'5', '3', '7', '2', '4', '6', '8'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tree.BFSLevelOrder()
			fmt.Println(got)
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf("BFSLevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_BFSLevelOrderRecursive(t *testing.T) {

	tests := []struct {
		name string
		tree *BinaryTree[rune]
		want []rune
	}{
		{
			name: "nil tree",
			tree: &BinaryTree[rune]{},
			want: defaultArray,
		}, {
			name: "root only",
			tree: NewTree('4'),
			want: []rune{'4'},
		}, {
			name: "level one full",
			tree: NewTree('4', '3', '5'),
			want: []rune{'4', '3', '5'},
		}, {
			name: "level two full",
			tree: NewTree('5', '3', '4', '2', '7', '6', '8'),
			want: []rune{'5', '3', '7', '2', '4', '6', '8'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tree.BFSLevelOrderRecursive()
			fmt.Println(got)
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf("BFSLevelOrderRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
