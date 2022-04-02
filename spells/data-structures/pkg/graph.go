package pkg

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"strings"
)

// Graph
// done:
// -implement graph data structure
// - BFS
// - DFS
// - Dijkstra
// - shortPathUnweighted
// todos:
// - union-find
// - topological sort
// - prims (spanning tree)
// - kruskal (spanning tree)
// - floyd warshall

type Graph[T constraints.Ordered] struct {
	nodes map[T][]*GraphEdge[T]
}

func NewGraph[T constraints.Ordered]() *Graph[T] {
	return &Graph[T]{
		nodes: make(map[T][]*GraphEdge[T]),
	}
}

func (g *Graph[T]) AddNode(node T) {
	g.nodes[node] = []*GraphEdge[T]{}
}

func (g *Graph[T]) AddEdge(nodeA T, nodeB T, cost int) {
	exist := func(edges []*GraphEdge[T], node T) bool {
		for _, e := range edges {
			if e.node == node {
				return true
			}
		}
		return false
	}
	appendEdge := func(edges []*GraphEdge[T], node T, cost int) []*GraphEdge[T] {
		if exist(edges, node) {
			return edges
		}
		return append(edges, &GraphEdge[T]{
			cost: cost,
			node: node,
		})

	}
	g.nodes[nodeA] = appendEdge(g.nodes[nodeA], nodeB, cost)
	g.nodes[nodeB] = appendEdge(g.nodes[nodeB], nodeA, cost)
}

func (g Graph[T]) String() string {
	builder := strings.Builder{}
	for k, v := range g.nodes {
		str := fmt.Sprintf("\nnode:%v => %v", k, v)
		builder.WriteString(str)
	}
	builder.WriteString("\n")
	return builder.String()
}

func (g *Graph[T]) Print() {
	fmt.Println(g)
}

func (g *Graph[T]) BFSLevelOrder(start T) []T {
	return graphBFSLevelOrder(start, g.nodes)
}

func (g *Graph[T]) DFSLevelOrder(start T) []T {
	return graphDFSLevelOrder(start, g.nodes)
}

func (g *Graph[T]) ShortPathUnweighted(start T, end T) []T {
	return shortPathUnweighted(start, end, g.nodes)
}

func graphBFSLevelOrder[T constraints.Ordered](start T, nodes map[T][]*GraphEdge[T]) []T {
	if len(nodes) == 0 {
		return nil
	}
	var result []T
	var queue Queue[*GraphEdge[T]]
	visited := make(map[T]bool)

	// add first elem and add level one to queue
	queue.Enqueue(&GraphEdge[T]{node: start})
	for !queue.IsEmpty() {
		cursor := (*GraphEdge[T])(queue.Dequeue())
		if visited[cursor.node] {
			continue
		}
		result = append(result, cursor.node)
		queue.EnqueueAll(nodes[cursor.node]...)
		visited[cursor.node] = true
	}
	return result
}

func graphDFSLevelOrder[T constraints.Ordered](start T, nodes map[T][]*GraphEdge[T]) []T {
	if len(nodes) == 0 {
		return nil
	}
	var result []T
	var stack Stack[*GraphEdge[T]]
	visited := make(map[T]bool)

	stack.Push(&GraphEdge[T]{node: start})
	for !stack.IsEmpty() {
		cursor := (*GraphEdge[T])(stack.Pop())
		if visited[cursor.node] {
			continue
		}
		result = append(result, cursor.node)
		stack.PushAll(nodes[cursor.node]...)
		visited[cursor.node] = true
	}

	return result
}

func shortPathUnweighted[T constraints.Ordered](start, end T, nodes map[T][]*GraphEdge[T]) []T {
	return nil
}

type GraphEdge[T constraints.Ordered] struct {
	cost int
	node T
}

func (g GraphEdge[T]) String() string {
	return fmt.Sprintf("node:%v c:%d", g.node, g.cost)
}
