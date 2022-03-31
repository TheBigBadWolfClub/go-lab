package pkg

import (
	"github.com/google/go-cmp/cmp"
	"golang.org/x/exp/constraints"
	"testing"
)

type edgeCreate[T constraints.Ordered] struct {
	A, B T
	cost int
}

func buildTestGraph[T constraints.Ordered](nodes []T, edges []edgeCreate[T]) *Graph[T] {
	graph := NewGraph[T]()
	for _, n := range nodes {
		graph.AddNode(n)
	}

	for _, e := range edges {
		graph.AddEdge(e.A, e.B, e.cost)
	}
	return graph
}
func TestNewGraph(t *testing.T) {
	g := NewGraph[int]()
	if g.nodes == nil {
		t.Fatalf("new graph")
	}
}

func TestGraph_AddNode(t *testing.T) {

	tests := []struct {
		name      string
		graph     *Graph[int]
		nodeValue int
		expected  *Graph[int]
	}{
		{
			name:      "empty graph",
			graph:     NewGraph[int](),
			nodeValue: 1,
			expected: &Graph[int]{
				nodes: map[int][]*GraphEdge[int]{1: {}},
			},
		}, {
			name: "one elem graph",
			graph: &Graph[int]{
				nodes: map[int][]*GraphEdge[int]{1: {}},
			},
			nodeValue: 2,
			expected: &Graph[int]{
				nodes: map[int][]*GraphEdge[int]{1: {}, 2: {}},
			},
		}, {
			name: "duplicated node",
			graph: &Graph[int]{
				nodes: map[int][]*GraphEdge[int]{1: {}},
			},
			nodeValue: 1,
			expected: &Graph[int]{
				nodes: map[int][]*GraphEdge[int]{1: {}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.graph.AddNode(tt.nodeValue)
			diff := cmp.Diff(tt.graph, tt.expected, cmp.AllowUnexported(Graph[int]{}, Graph[int]{}))

			if diff != "" {
				t.Fatalf("AddNode() : %s", diff)
			}
		})
	}
}

func TestGraph_AddEdge(t *testing.T) {

	tests := []struct {
		name     string
		graph    *Graph[int]
		expected *Graph[int]
		edge     []edgeCreate[int]
	}{
		{
			name:     "empty",
			graph:    buildTestGraph[int]([]int{1, 2}, []edgeCreate[int]{}),
			expected: buildTestGraph[int]([]int{1, 2}, []edgeCreate[int]{{1, 2, 10}}),
			edge:     []edgeCreate[int]{{1, 2, 10}},
		}, {
			name:     "add 2 times same edge",
			graph:    buildTestGraph[int]([]int{1, 2}, []edgeCreate[int]{}),
			expected: buildTestGraph[int]([]int{1, 2}, []edgeCreate[int]{{1, 2, 10}}),
			edge:     []edgeCreate[int]{{1, 2, 10}, {1, 2, 10}},
		}, {
			name:     "add 3",
			graph:    buildTestGraph[int]([]int{1, 2, 3}, []edgeCreate[int]{}),
			expected: buildTestGraph[int]([]int{1, 2, 3}, []edgeCreate[int]{{1, 2, 10}, {1, 3, 30}}),
			edge:     []edgeCreate[int]{{1, 2, 10}, {1, 3, 30}},
		}, {
			name:     "fully circular",
			graph:    buildTestGraph[int]([]int{1, 2, 3}, []edgeCreate[int]{}),
			expected: buildTestGraph[int]([]int{1, 2, 3}, []edgeCreate[int]{{1, 2, 10}, {1, 3, 30}, {2, 3, 23}}),
			edge:     []edgeCreate[int]{{1, 2, 10}, {1, 3, 30}, {2, 3, 23}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, edge := range tt.edge {
				tt.graph.AddEdge(edge.A, edge.B, edge.cost)
			}
			diff := cmp.Diff(tt.graph, tt.expected, cmp.AllowUnexported(Graph[int]{}, GraphEdge[int]{}))
			if diff != "" {
				t.Fatalf("AddEdge() : %s", diff)
			}
		})
	}
}
