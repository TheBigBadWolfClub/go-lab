package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"unique"
)

func TestDeduplicate(t *testing.T) {
	t.Run("Deduplicate", func(t *testing.T) {
		strs := []string{"a", "b", "c", "a", "b", "c"}
		hdls := Deduplicate(strs)
		assert.Equal(t, 3, len(hdls))

		var got []string
		for _, h := range hdls {
			got = append(got, h.Value())
		}
		assert.Equal(t, []string{"a", "b", "c"}, got)
	})
}

func TestCountUnique(t *testing.T) {
	t.Run("CountUnique", func(t *testing.T) {
		numbers := []int{1, 2, 3, 2, 1, 4, 5, 4}
		assert.Equal(t, 5, CountUnique(numbers))
	})
}

func TestDeduplicateStruct(t *testing.T) {
	t.Run("DeduplicateStruct", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{
			{"Alice", 30},
			{"Bob", 25},
			{"Alice", 30},
			{"Charlie", 35},
			{"Bob", 25},
		}

		hdls := DeduplicateStruct(people)
		assert.Equal(t, 3, len(hdls))

		var got []Person
		for _, h := range hdls {
			got = append(got, h.Value())
		}
		assert.ElementsMatch(t, []Person{
			{"Alice", 30},
			{"Bob", 25},
			{"Charlie", 35},
		}, got)
	})
}

func TestNewSet(t *testing.T) {
	t.Run("NewSet", func(t *testing.T) {
		s := NewSet[string]()
		assert.NotNil(t, s)
		assert.NotNil(t, s.data)
	})
}

func TestSet_Add(t *testing.T) {
	type person struct {
		name string
		age  int
	}

	type testCase[T comparable] struct {
		name      string
		set       *Set[T]
		args      []T
		finalSize int
	}
	tests := []testCase[person]{
		{
			name:      "should add a person to empty set",
			set:       NewSet[person](),
			args:      []person{{"Alice", 30}},
			finalSize: 1,
		},
		{
			name: "should add to not empty set",
			set: &Set[person]{data: map[unique.Handle[person]]struct{}{
				unique.Make(person{"Alice", 30}): {},
			}},
			args:      []person{{"Bob", 25}},
			finalSize: 2,
		},
		{
			name: "should not add duplicate person",
			set: &Set[person]{data: map[unique.Handle[person]]struct{}{
				unique.Make(person{"Alice", 30}): {},
			}},
			args:      []person{{"Alice", 30}},
			finalSize: 1,
		},
		{
			name:      "should add multiple persons",
			set:       NewSet[person](),
			args:      []person{{"Alice", 30}, {"Bob", 25}, {"Bob", 25}},
			finalSize: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			for _, arg := range tt.args {
				tt.set.Add(arg)
			}

			assert.Equal(t, tt.finalSize, len(tt.set.data))
		})
	}
}

func TestSet_Contains(t *testing.T) {
	t.Run("Contains", func(t *testing.T) {
		newSet := NewSet[string]()
		newSet.Add("Alice")
		newSet.Add("Bob")

		assert.True(t, newSet.Contains("Alice"))
		assert.False(t, newSet.Contains("Charlie"))
	})
}

func TestSet_Remove(t *testing.T) {
	t.Run("Remove", func(t *testing.T) {
		newSet := NewSet[string]()
		newSet.Add("Alice")
		newSet.Add("Bob")

		newSet.Remove("Alice")
		assert.False(t, newSet.Contains("Alice"))
		assert.True(t, newSet.Contains("Bob"))
	})
}

func TestSet_String(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		newSet := NewSet[string]()
		newSet.Add("Alice")
		newSet.Add("Bob")

		assert.Equal(t, 2, newSet.Size())
	})
}

func TestNewCache(t *testing.T) {
	t.Run("NewCache", func(t *testing.T) {
		c := NewCache[string]()
		assert.NotNil(t, c)
		assert.NotNil(t, c.data)
	})
}

func TestCache_Set(t *testing.T) {
	type testCase[T comparable] struct {
		name       string
		cache      *Cache[T]
		key        string
		value      T
		expiration time.Duration
		finalSize  int
	}
	tests := []testCase[int]{
		{
			name:       "should add a value to empty cache",
			cache:      NewCache[int](),
			key:        "Alice",
			value:      30,
			expiration: 1,
			finalSize:  1,
		},
		{
			name: "should update the value in the cache",
			cache: &Cache[int]{data: map[unique.Handle[string]]CacheEntry[int]{
				unique.Make("Alice"): {Value: 30, Expiration: time.Now().Add(1 * time.Second)},
			}},
			key:        "Alice",
			value:      25,
			expiration: 1,
			finalSize:  1,
		},
		{
			name: "should add multiple values to the cache",
			cache: &Cache[int]{data: map[unique.Handle[string]]CacheEntry[int]{
				unique.Make("Alice"): {Value: 30, Expiration: time.Now().Add(1 * time.Second)},
			}},
			key:        "Mike",
			value:      30,
			expiration: 1,
			finalSize:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cache.Set(tt.key, tt.value, tt.expiration)
			assert.Equal(t, tt.finalSize, len(tt.cache.data))
		})
	}
}

func TestCache_Get(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		c := NewCache[int]()
		c.Set("Alice", 30, 1*time.Second)
		c.Set("Bob", 25, 1*time.Second)

		got, ok := c.Get("Alice")
		assert.True(t, ok)
		assert.Equal(t, 30, got)

		got, ok = c.Get("Mike")
		assert.False(t, ok)
		assert.Zero(t, got)
	})
}

func TestCache_Delete(t *testing.T) {
	t.Run("Delete", func(t *testing.T) {
		c := NewCache[int]()
		c.Set("Alice", 30, 1*time.Nanosecond)
		c.Set("Bob", 25, 1*time.Second)

		time.Sleep(1 * time.Millisecond)
		c.Clean()
		assert.Equal(t, 1, len(c.data))
	})
}

func TestNewGraph(t *testing.T) {
	t.Run("NewGraph", func(t *testing.T) {
		g := NewGraph[string]()
		assert.NotNil(t, g)
		assert.NotNil(t, g.nodes)
		assert.NotNil(t, g.edges)
	})
}

func TestGraph_AddNode(t *testing.T) {

	type testCase[T comparable] struct {
		name         string
		graph        *Graph[T]
		value        T
		expectedSize int
	}
	tests := []testCase[string]{
		{
			name:         "should add a node to empty graph",
			graph:        NewGraph[string](),
			value:        "Alice",
			expectedSize: 1,
		},
		{
			name: "should add a node to not empty graph",
			graph: &Graph[string]{
				nodes: map[unique.Handle[string]]struct{}{
					unique.Make("Alice"): {},
				},
			},
			value:        "Bob",
			expectedSize: 2,
		},
		{
			name: "should not add duplicate node",
			graph: &Graph[string]{
				nodes: map[unique.Handle[string]]struct{}{
					unique.Make("Alice"): {},
				},
			},
			value:        "Alice",
			expectedSize: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.graph.AddNode(tt.value)
			assert.Equal(t, tt.expectedSize, len(tt.graph.nodes))
		})
	}
}

func TestGraph_AddEdge(t *testing.T) {

	type testCase[T comparable] struct {
		name      string
		g         Graph[T]
		edge1     T
		edge2     T
		sizeNodes int
		sizeEdges int
	}
	tests := []testCase[string]{
		{
			name:      "should add an edge to empty graph",
			g:         *NewGraph[string](),
			edge1:     "Alice",
			edge2:     "Bob",
			sizeNodes: 2,
			sizeEdges: 2,
		},
		{
			name: "should add an edge to not empty graph",
			g: Graph[string]{
				nodes: map[unique.Handle[string]]struct{}{
					unique.Make("Alice"): {},
				},
				edges: map[unique.Handle[string]]map[unique.Handle[string]]struct{}{},
			},
			edge1:     "Alice",
			edge2:     "Bob",
			sizeNodes: 2,
			sizeEdges: 2,
		},
		{
			name: "should not add duplicate edge",
			g: Graph[string]{
				nodes: map[unique.Handle[string]]struct{}{
					unique.Make("Alice"): {},
					unique.Make("Bob"):   {},
				},
				edges: map[unique.Handle[string]]map[unique.Handle[string]]struct{}{
					unique.Make("Alice"): {unique.Make("Bob"): {}},
					unique.Make("Bob"):   {unique.Make("Alice"): {}},
				},
			},
			edge1:     "Alice",
			edge2:     "Bob",
			sizeNodes: 2,
			sizeEdges: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.AddEdge(tt.edge1, tt.edge2)
			assert.Equal(t, tt.sizeNodes, len(tt.g.nodes))
			assert.Equal(t, tt.sizeEdges, len(tt.g.edges))
		})
	}
}

func TestGraph_HasEdge(t *testing.T) {
	t.Run("HasEdge", func(t *testing.T) {
		g := NewGraph[string]()
		g.AddEdge("Alice", "Bob")
		g.AddEdge("Alice", "Charlie")

		assert.True(t, g.HasEdge("Alice", "Bob"))
		assert.False(t, g.HasEdge("Alice", "Mike"))
	})
}

func TestGraph_GetNeighbors(t *testing.T) {
	t.Run("GetNeighbors", func(t *testing.T) {
		g := NewGraph[string]()
		g.AddEdge("Alice", "Bob")
		g.AddEdge("Alice", "Charlie")

		neighbors := g.GetNeighbors("Alice")
		assert.ElementsMatch(t, []string{"Bob", "Charlie"}, neighbors)
	})
}

func TestMemoize(t *testing.T) {
	t.Run("Memoize", func(t *testing.T) {
		f := func(n int) int {
			return n * n
		}

		memoized := Memoize(f)
		assert.Equal(t, 25, memoized(5))
	})
}
