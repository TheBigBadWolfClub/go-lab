package pkg

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"strings"
)

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

type GraphEdge[T constraints.Ordered] struct {
	cost int
	node T
}

func (g GraphEdge[T]) String() string {
	return fmt.Sprintf("node:%v c:%d", g.node, g.cost)
}
