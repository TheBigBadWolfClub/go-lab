package pkg

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"math"
	"testing"
)

func NewTree(values ...rune) *TreeNode {
	var tree *TreeNode
	for _, v := range values {
		tree = insert(tree, v)
	}
	return tree
}

func Test_Alpha(t *testing.T) {

	var tree *TreeNode
	tree = insert(tree, 'F')
	insert(tree, 'D')
	insert(tree, 'B')
	insert(tree, 'C')
	insert(tree, 'A')

	s := tree.String()
	fmt.Println(s)
}

func Test_insert(t *testing.T) {

	tests := []struct {
		name  string
		tree  *TreeNode
		value int
		want  *TreeNode
	}{
		{
			name:  "first value of tree",
			tree:  nil,
			value: 4,
			want: &TreeNode{
				Left:  nil,
				Value: 4,
				Right: nil,
			},
		}, {
			name: "value already in tree, do not add",
			tree: &TreeNode{
				Left:  nil,
				Value: 4,
				Right: nil,
			},
			value: 4,
			want: &TreeNode{
				Left:  nil,
				Value: 4,
				Right: nil,
			},
		}, {
			name: "insert left of tree",
			tree: &TreeNode{
				Left:  nil,
				Value: 4,
				Right: nil,
			},
			value: 2,
			want: &TreeNode{
				Left: &TreeNode{
					Left:  nil,
					Value: 2,
					Right: nil,
				},
				Value: 4,
				Right: nil,
			},
		}, {
			name: "insert right of tree",
			tree: &TreeNode{
				Left:  nil,
				Value: 4,
				Right: nil,
			},
			value: 6,
			want: &TreeNode{
				Left:  nil,
				Value: 4,
				Right: &TreeNode{
					Left:  nil,
					Value: 6,
					Right: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := insert(tt.tree, rune(tt.value))
			fmt.Println(got)
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxDepth(t *testing.T) {

	tests := []struct {
		name string
		tree *TreeNode
		want int
	}{
		{
			name: "nil tree, maxDepth=0",
			tree: nil,
			want: 0,
		}, {
			name: "only root, maxDepth=1",
			tree: &TreeNode{
				Left:  nil,
				Value: 0,
				Right: nil,
			},
			want: 1,
		}, {
			name: "level 2, maxDepth=2",
			tree: NewTree('2', '1', '3'),
			want: 2,
		}, {
			name: "level 2 only right, maxDepth=2",
			tree: NewTree('2', '3'),
			want: 2,
		}, {
			name: "level 2 only left, maxDepth=2",
			tree: NewTree('2', '1'),
			want: 2,
		}, {
			name: "only right nodes, maxDepth=10",
			tree: NewTree('0', '1', '2', '3', '4', '5', '6', '7', '8', '9'),
			want: 10,
		}, {
			name: "only right nodes, maxDepth=10",
			tree: NewTree('9', '8', '7', '6', '5', '4', '3', '2', '1', '0'),
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.tree)
			got := maxDepth(tt.tree)
			if got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	tests := []struct {
		name string
		tree *TreeNode
		want *TreeNode
	}{
		{
			name: "reverse",
			tree: NewTree('2', '1', '3'),
			want: &TreeNode{
				Left:  &TreeNode{nil, '3', nil},
				Value: '2',
				Right: &TreeNode{nil, '1', nil},
			},
		}, {
			name: "reverse",
			tree: NewTree('3', '2', '1', '4', '5'),
			want: &TreeNode{
				Left:  &TreeNode{&TreeNode{nil, '5', nil}, '4', nil},
				Value: '3',
				Right: &TreeNode{nil, '2', &TreeNode{nil, '1', nil}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverse(tt.tree)
			fmt.Println(got)
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxValue(t *testing.T) {

	tests := []struct {
		name string
		tree *TreeNode
		want float64
	}{
		{
			name: "nil, expect infinity",
			tree: nil,
			want: math.Inf(-1),
		}, {
			name: "only root",
			tree: &TreeNode{nil, '1', nil},
			want: '1',
		}, {
			name: "2 level",
			tree: NewTree('5', '4', '6', '2', '3', '1'),
			want: '6',
		}, {
			name: "some leaf",
			tree: NewTree('5', '4', '6', '2', '3', '1', '8', '7', '9'),
			want: '9',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxValue(tt.tree); got != tt.want {
				t.Errorf("maxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minValue(t *testing.T) {

	tests := []struct {
		name string
		tree *TreeNode
		want float64
	}{{
		name: "nil, expect infinity",
		tree: nil,
		want: math.Inf(1),
	},
		{
			name: "only root",
			tree: &TreeNode{nil, '1', nil},
			want: '1',
		}, {
			name: "2 level",
			tree: NewTree('7', '9', '8', '6', 'A'),
			want: '6',
		}, {
			name: "some leaf",
			tree: NewTree('5', '4', '6', '2', '3', '1'),
			want: '1',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minValue(tt.tree); got != tt.want {
				t.Errorf("minValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_depthFirstSearch(t *testing.T) {

	tests := []struct {
		name string
		tree *TreeNode
		want []rune
	}{
		{
			name: "nil tree",
			tree: nil,
			want: []rune{},
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
			got := dfsPreOrder(tt.tree)
			fmt.Println(string(got))
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf("depthFirstSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dfsInOrder(t *testing.T) {

	tests := []struct {
		name string
		tree *TreeNode
		want []rune
	}{
		{
			name: "nil tree",
			tree: nil,
			want: []rune{},
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
			got := dfsInOrder(tt.tree)
			fmt.Println(string(got))
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf("depthFirstSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dfsPostOrder(t *testing.T) {

	tests := []struct {
		name string
		tree *TreeNode
		want []rune
	}{
		{
			name: "nil tree",
			tree: nil,
			want: []rune{},
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
			got := dfsPostOrder(tt.tree)
			fmt.Println(string(got))
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf("depthFirstSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bfsLevelOrder(t *testing.T) {

	tests := []struct {
		name string
		tree *TreeNode
		want []rune
	}{
		{
			name: "nil tree",
			tree: nil,
			want: []rune{},
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
			got := bfsLevelOrder(tt.tree)
			fmt.Println(string(got))
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf("bfsLevelOrder() = %v, want %v", got, tt.want)
			}

			gotRec := bfsLevelOrderRecursive(tt.tree, []*TreeNode{})
			fmt.Println(string(gotRec))
			diffRec := cmp.Diff(gotRec, tt.want)
			if diffRec != "" {
				t.Errorf("bfsLevelOrder() = %v, want %v", gotRec, tt.want)
			}
		})
	}
}

func Test_find(t *testing.T) {

	tests := []struct {
		name  string
		tree  *TreeNode
		value rune
		want  bool
	}{
		{
			name: "nil tree",
			tree: nil,
			want: false,
		}, {
			name:  "root only, is value",
			tree:  NewTree('4'),
			value: '4',
			want:  true,
		}, {
			name:  "root only, is not value",
			tree:  NewTree('4'),
			value: '1',
			want:  false,
		}, {
			name:  "level 1, is  value at left",
			tree:  NewTree('4', '3', '5'),
			value: '3',
			want:  true,
		}, {
			name:  "level 1, is  value at right",
			tree:  NewTree('4', '3', '5'),
			value: '5',
			want:  true,
		}, {
			name:  "level 1, value not found",
			tree:  NewTree('4', '3', '5'),
			value: '1',
			want:  false,
		}, {
			name:  "level 2, value not found",
			tree:  NewTree('4', '3', '5', '2', '6'),
			value: '1',
			want:  false,
		}, {
			name:  "level 2, value  found at left",
			tree:  NewTree('4', '3', '5', '2', '6'),
			value: '2',
			want:  true,
		}, {
			name:  "level 2, value  found at right",
			tree:  NewTree('4', '3', '5', '2', '6'),
			value: '6',
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := find(tt.tree, tt.value)
			fmt.Println(got)
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf("depthFirstSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_breadthFirstSearch(t *testing.T) {

	tests := []struct {
		name string
		tree *TreeNode
		want []rune
	}{
		{
			name: "nil tree",
			tree: nil,
			want: []rune{},
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
			got := breadthFirstSearch(tt.tree, nil)
			fmt.Println(string(got))
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf("depthFirstSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
