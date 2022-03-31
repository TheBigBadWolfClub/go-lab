package pkg

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math"
)

const strNilNode = "_"

type BinaryTree[T constraints.Ordered] struct {
	root *TreeNode[T]
}

func (b *BinaryTree[T]) Insert(value T) {
	if b.root == nil {
		b.root = &TreeNode[T]{Value: value}
		return
	}

	treeNodeAdd(b.root, value)
}

func (b *BinaryTree[T]) Exists(value T) bool {
	if b.root == nil {
		return false
	}

	return treeNodeExists(b.root, value)
}

func (b *BinaryTree[T]) Delete(value T) bool {
	if b.root == nil {
		return false
	}

	var ok bool
	b.root, ok = delete(b.root, value)
	return ok
}

func (b *BinaryTree[T]) MaxDepth() int {
	if b.root == nil {
		return 0
	}
	return maxDepth(b.root)
}

func (b *BinaryTree[T]) Reverse() {
	if b.root == nil {
		return
	}
	reverse(b.root)
}

func (b *BinaryTree[T]) MaxValue() (T, bool) {
	if b.root == nil {
		var zero T
		return zero, false
	}
	return maxValue(b.root), true
}

func (b *BinaryTree[T]) MinValue() (T, bool) {
	if b.root == nil {
		var zero T
		return zero, false
	}
	return minValue(b.root), true
}

func (b *BinaryTree[T]) DFSPreOrder() []T {
	var result []T
	if b.root == nil {
		return result
	}

	return dfsPreOrder(b.root)
}

func (b *BinaryTree[T]) DFSInOrder() []T {
	var result []T
	if b.root == nil {
		return result
	}

	return dfsInOrder(b.root)
}

func (b *BinaryTree[T]) DFSPostOrder() []T {
	var result []T
	if b.root == nil {
		return result
	}

	return dfsPostOrder(b.root)
}

func (b *BinaryTree[T]) BFSLevelOrder() []T {
	var result []T
	if b.root == nil {
		return result
	}

	nodes := []TreeNode[T]{*b.root}
	return bfsLevelOrder(nodes)
}

func (b *BinaryTree[T]) BFSLevelOrderRecursive() []T {
	var result []T
	if b.root == nil {
		return result
	}

	nodes := []TreeNode[T]{*b.root}
	return bfsLevelOrderRecursive(nodes)
}

type TreeNode[T constraints.Ordered] struct {
	Left  *TreeNode[T]
	Value T
	Right *TreeNode[T]
}

func treeNodeAdd[T constraints.Ordered](node *TreeNode[T], value T) {
	if node.Left == nil && value < node.Value {
		node.Left = &TreeNode[T]{Value: value}

	}

	if node.Right == nil && value > node.Value {
		node.Right = &TreeNode[T]{Value: value}
	}

	if value < node.Value {
		treeNodeAdd(node.Left, value)
	}

	if value > node.Value {
		treeNodeAdd(node.Right, value)
	}
}

func treeNodeExists[T constraints.Ordered](node *TreeNode[T], value T) bool {
	if node == nil {
		return false
	}
	if value == node.Value {
		return true
	}
	if value < node.Value {
		return treeNodeExists(node.Left, value)
	}
	if value > node.Value {
		return treeNodeExists(node.Right, value)
	}

	return false
}

func maxDepth[T constraints.Ordered](node *TreeNode[T]) int {
	if node == nil {
		return 0
	}

	r := maxDepth(node.Right) + 1
	l := maxDepth(node.Left) + 1
	max := math.Max(float64(r), float64(l))
	return int(max)
}

func reverse[T constraints.Ordered](node *TreeNode[T]) {
	if node == nil {
		return
	}

	node.Right, node.Left = node.Left, node.Right
	reverse(node.Left)
	reverse(node.Right)
}

func maxValue[T constraints.Ordered](node *TreeNode[T]) T {

	max := node.Value

	if node.Right != nil {
		if r := maxValue(node.Right); r > max {
			max = r
		}
	}

	if node.Left != nil {
		if l := maxValue(node.Left); l > max {
			max = l
		}
	}

	return max

}

func minValue[T constraints.Ordered](node *TreeNode[T]) T {
	max := node.Value

	if node.Right != nil {
		if r := minValue(node.Right); r < max {
			max = r
		}
	}

	if node.Left != nil {
		if l := minValue(node.Left); l < max {
			max = l
		}
	}

	return max
}

func dfsPreOrder[T constraints.Ordered](node *TreeNode[T]) []T {
	var result []T
	if node == nil {
		return result
	}

	result = append(result, node.Value)
	result = append(result, dfsPreOrder(node.Left)...)
	result = append(result, dfsPreOrder(node.Right)...)
	return result
}

func dfsInOrder[T constraints.Ordered](node *TreeNode[T]) []T {
	var result []T
	if node == nil {
		return result
	}

	result = append(result, dfsInOrder(node.Left)...)
	result = append(result, node.Value)
	result = append(result, dfsInOrder(node.Right)...)
	return result
}

func dfsPostOrder[T constraints.Ordered](node *TreeNode[T]) []T {
	var result []T
	if node == nil {
		return result
	}

	result = append(result, dfsPostOrder(node.Left)...)
	result = append(result, dfsPostOrder(node.Right)...)
	result = append(result, node.Value)
	return result
}

func bfsLevelOrder[T constraints.Ordered](queue []TreeNode[T]) []T {
	var result []T
	for len(queue) > 0 {
		result = append(result, queue[0].Value)
		if queue[0].Left != nil {
			queue = append(queue, *queue[0].Left)
		}

		if queue[0].Right != nil {
			queue = append(queue, *queue[0].Right)
		}

		queue = queue[1:]
	}
	return result
}

func bfsLevelOrderRecursive[T constraints.Ordered](queue []TreeNode[T]) []T {
	var result []T
	if len(queue) == 0 {
		return result
	}

	cur := queue[0]
	result = append(result, cur.Value)
	if cur.Left != nil {
		queue = append(queue, *cur.Left)
	}
	if cur.Left != nil {
		queue = append(queue, *cur.Right)
	}

	if len(queue) > 0 {
		return append(result, bfsLevelOrderRecursive(queue[1:])...)
	}

	return result
}

func delete[T constraints.Ordered](node *TreeNode[T], value T) (*TreeNode[T], bool) {

	if node == nil {
		return node, false
	}

	var ok bool
	if value < node.Value {
		node.Left, ok = delete(node.Left, value)
		return node, ok
	}

	if value > node.Value {
		node.Right, ok = delete(node.Right, value)
		return node, ok
	}

	if node.Left == nil && node.Right == nil {
		return nil, true
	}

	if node.Left == nil {
		return node.Right, true
	}
	if node.Right == nil {
		return node.Left, true
	}

	node.Value = minValue(node.Right)
	node.Right, _ = delete(node.Right, node.Value)
	return node, true
}

func (t *TreeNode[T]) String() string {
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

	return fmt.Sprintf("(%s %v %s)", left, t.Value, right)
}
