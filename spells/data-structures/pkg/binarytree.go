package pkg

import (
	"fmt"
	"math"
)

const strNilNode = "_"

type TreeNode struct {
	Left  *TreeNode
	Value rune
	Right *TreeNode
}

func insert(t *TreeNode, v rune) *TreeNode {
	if t == nil {
		return &TreeNode{Left: nil, Value: v, Right: nil}
	}

	if t.Value > v {
		t.Left = insert(t.Left, v)
	} else if t.Value < v {
		t.Right = insert(t.Right, v)
	}
	return t
}

func find(node *TreeNode, value rune) bool {
	if node == nil {
		return false
	}

	if node.Value > value {
		return find(node.Left, value)
	}

	if node.Value < value {
		return find(node.Right, value)
	}
	return node.Value == value
}

func maxDepth(t *TreeNode) int {
	if t == nil {
		return 0
	}
	left := float64(maxDepth(t.Left))
	right := float64(maxDepth(t.Right))
	return 1 + int(math.Max(left, right))
}

func reverse(t *TreeNode) *TreeNode {
	if t == nil {
		return nil
	}

	t.Left, t.Right = reverse(t.Right), reverse(t.Left)
	return t
}

func maxValue(t *TreeNode) float64 {
	if t == nil {
		return math.Inf(-1)
	}

	l := maxValue(t.Left)
	r := maxValue(t.Right)
	return math.Max(float64(t.Value), math.Max(l, r))
}

func minValue(t *TreeNode) float64 {
	if t == nil {
		return math.Inf(1)
	}

	l := minValue(t.Left)
	r := minValue(t.Right)
	mx := math.Min(l, r)
	return math.Min(mx, float64(t.Value))
}

func depthFirstSearch(node *TreeNode) []rune {

	if node == nil {
		return []rune{}
	}

	res := append([]rune{}, node.Value)
	res = append(res, depthFirstSearch(node.Left)...)
	res = append(res, depthFirstSearch(node.Right)...)
	return res
}

func dfsPreOrder(node *TreeNode) []rune {
	return depthFirstSearch(node)
}

func dfsInOrder(node *TreeNode) []rune {
	if node == nil {
		return []rune{}
	}

	res := append([]rune{}, node.Value)
	l := append(dfsInOrder(node.Left), res...)
	r := dfsInOrder(node.Right)
	return append(l, r...)
}

func dfsPostOrder(node *TreeNode) []rune {
	if node == nil {
		return []rune{}
	}

	res := append([]rune{}, node.Value)
	l := append(dfsPostOrder(node.Left), dfsPostOrder(node.Right)...)
	return append(l, res...)
}

func breadthFirstSearch(node *TreeNode, queue []*TreeNode) []rune {

	if node == nil {
		return []rune{}
	}

	if node.Left != nil {
		queue = append(queue, node.Left)
	}

	if node.Right != nil {
		queue = append(queue, node.Right)
	}

	res := append([]rune{}, node.Value)
	if len(queue) == 0 {
		return res
	}

	next, reminder := queue[:1], queue[1:]
	res = append(res, breadthFirstSearch(next[0], reminder)...)
	return res
}

func delete(node *TreeNode, value rune) *TreeNode {
	if node == nil {
		return nil
	}

	if node.Value == value {
		if node.Left == nil && node.Right == nil {
			return nil
		}
		if node.Left == nil && node.Right != nil {
			node = node.Right
			return node
		}

		if node.Left != nil && node.Right == nil {
			node = node.Left
			return node
		}
	}

	return nil
}

func (t *TreeNode) String() string {
	if t == nil {
		return "nil"
	}

	left := strNilNode
	if t.Left != nil {
		left = t.Left.String()
	}

	right := strNilNode
	if t.Right != nil {
		right = t.Right.String()
	}

	return fmt.Sprintf("(%s %s %s)", left, string(t.Value), right)
}
