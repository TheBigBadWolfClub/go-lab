package internal

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math"
)

type PriorityNode[R constraints.Ordered] struct {
	cost int
	from R
}

func (p PriorityNode[R]) String() string {
	return fmt.Sprintf("[%v => %d]", p.from, p.cost)
}

func NewPriorityMap[R constraints.Ordered]() PriorityMap[R] {
	return PriorityMap[R]{
		m:       make(map[R]*PriorityNode[R]),
		visited: make(map[R]bool),
	}
}

type PriorityMap[R constraints.Ordered] struct {
	m       map[R]*PriorityNode[R]
	visited map[R]bool
}

func (p *PriorityMap[R]) Len() int {
	return len((map[R]*PriorityNode[R])(p.m))
}

func (p *PriorityMap[R]) MarkVisited(r R) {
	p.visited[r] = true
}

func (p *PriorityMap[R]) HaveVisitAll() bool {
	return p.Len() == len((map[R]bool)(p.visited))
}
func (p *PriorityMap[R]) Add(defaultCost int, vertices ...R) {
	for _, v := range vertices {
		_, okVisited := p.visited[v]
		_, okAdd := p.m[v]
		if !okAdd && !okVisited {
			p.m[v] = &PriorityNode[R]{cost: defaultCost}
		}
	}
}

func (p *PriorityMap[R]) MinCostVert() R {
	minCost := math.MaxInt
	var minKey R
	for k, v := range p.m {
		if _, ok := p.visited[k]; !ok && v.cost <= minCost {
			minCost = v.cost
			minKey = k
		}
	}
	return minKey
}

func (p *PriorityMap[R]) UpdateCost(vertice R, from R, cost int) {
	if _, ok := p.visited[vertice]; !ok {
		p.m[vertice] = &PriorityNode[R]{cost: cost, from: from}
	}
}

func (p *PriorityMap[R]) Path(start, end R) ([]R, int) {

	path := []R{end}
	curEdge := p.m[end]
	cost := 0
	for i := 1; i < p.Len(); i++ {
		path = append([]R{curEdge.from}, path...)
		cost += curEdge.cost
		if curEdge.from == start {
			break
		}
		curEdge = p.m[curEdge.from]
	}
	return path, cost
}
