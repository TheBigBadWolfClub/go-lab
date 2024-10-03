package internal

import (
	"maps"
	"slices"
	"time"
	"unique"
)

// Deduplicate returns a slice of unique handles for the given slice of strings.
func Deduplicate(strs []string) []unique.Handle[string] {
	mapping := make(map[unique.Handle[string]]struct{})
	for s := range slices.Values(strs) {
		h := unique.Make(s)
		mapping[h] = struct{}{}
	}

	seqKeys := maps.Keys(mapping)
	return slices.Collect(seqKeys)
}

// CountUnique returns the number of unique elements in the given slice of integers.
func CountUnique(numbers []int) int {
	mapping := make(map[unique.Handle[int]]struct{})
	for n := range slices.Values(numbers) {
		h := unique.Make(n)
		mapping[h] = struct{}{}
	}

	return len(mapping)
}

// DeduplicateStruct returns a slice of unique handles for the given slice of structs.
func DeduplicateStruct[T comparable](numbers []T) []unique.Handle[T] {
	mapping := make(map[unique.Handle[T]]struct{})
	for n := range slices.Values(numbers) {
		h := unique.Make(n)
		mapping[h] = struct{}{}
	}

	seqKeys := maps.Keys(mapping)
	return slices.Collect(seqKeys)
}

/**
* Implementation of a generic set data structure using the unique package.
* The set should support add, remove, and contains operations.
 */

// Set is a generic set data structure.
type Set[T comparable] struct {
	data map[unique.Handle[T]]struct{}
}

// NewSet creates a new set.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		data: make(map[unique.Handle[T]]struct{}),
	}
}

// Add adds a value to the set.
func (s *Set[T]) Add(value T) {
	h := unique.Make(value)
	if _, ok := s.data[h]; !ok {
		s.data[h] = struct{}{}
	}
}

// Remove removes a value from the set.
func (s *Set[T]) Remove(value T) {
	h := unique.Make(value)
	delete(s.data, h)
}

// Contains checks if a value is in the set.
func (s *Set[T]) Contains(value T) bool {
	h := unique.Make(value)
	_, ok := s.data[h]
	return ok
}

// Size returns the number of elements in the set.
func (s *Set[T]) Size() int {
	return len(s.data)
}

/*
* Implement a cache with expiration using the unique package.
* Each cache entry should have a value and an expiration time.
* Test the cache with different types and expiration times
 */

// CacheEntry represents a cache entry.
type CacheEntry[T comparable] struct {
	Value      T
	Expiration time.Time
}

// Cache is a generic cache data structure.
type Cache[T comparable] struct {
	data map[unique.Handle[string]]CacheEntry[T]
}

// NewCache creates a new cache.
func NewCache[T comparable]() *Cache[T] {
	return &Cache[T]{
		data: make(map[unique.Handle[string]]CacheEntry[T]),
	}
}

// Set adds a value to the cache with an expiration time.
func (c *Cache[T]) Set(key string, value T, expiration time.Duration) {
	c.data[unique.Make(key)] = CacheEntry[T]{
		Value:      value,
		Expiration: time.Now().Add(expiration),
	}
}

// Get returns a value from the cache if it exists and has not expired.
func (c *Cache[T]) Get(key string) (T, bool) {
	value, ok := c.data[unique.Make(key)]
	if !ok {
		var zero T
		return zero, false
	}

	if time.Now().After(value.Expiration) {
		delete(c.data, unique.Make(key))
		var zero T
		return zero, false
	}

	return value.Value, true
}

// Clean removes expired entries from the cache.
func (c *Cache[T]) Clean() {
	for key, value := range c.data {
		if time.Now().After(value.Expiration) {
			delete(c.data, key)
		}
	}
}

/*
* Implement a simple undirected graph using the unique package
* for efficient node storage and edge checking.
 */

// Node represents a graph node.
type Node[T comparable] struct {
	Value T
}

// Graph is a generic undirected graph data structure.
type Graph[T comparable] struct {
	nodes map[unique.Handle[T]]struct{}
	edges map[unique.Handle[T]]map[unique.Handle[T]]struct{}
}

// NewGraph creates a new graph.
func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{
		nodes: make(map[unique.Handle[T]]struct{}),
		edges: make(map[unique.Handle[T]]map[unique.Handle[T]]struct{}),
	}
}

// AddNode adds a node to the graph.
func (g *Graph[T]) AddNode(value T) {
	g.nodes[unique.Make(value)] = struct{}{}
}

// AddEdge adds an edge between two nodes in the graph.
func (g *Graph[T]) AddEdge(value1, value2 T) {
	n1 := unique.Make(value1)
	n2 := unique.Make(value2)

	g.nodes[n1] = struct{}{}
	g.nodes[n2] = struct{}{}

	if _, ok := g.edges[n1]; !ok {
		g.edges[n1] = make(map[unique.Handle[T]]struct{})
	}
	g.edges[n1][n2] = struct{}{}

	if _, ok := g.edges[n2]; !ok {
		g.edges[n2] = make(map[unique.Handle[T]]struct{})
	}
	g.edges[n2][n1] = struct{}{}
}

// HasEdge checks if an edge exists between two nodes in the graph.
func (g *Graph[T]) HasEdge(value1, value2 T) bool {
	n1 := unique.Make(value1)
	n2 := unique.Make(value2)

	if _, ok := g.edges[n1]; !ok {
		return false
	}
	if _, ok := g.edges[n1][n2]; !ok {
		return false
	}
	return true

}

// GetNeighbors returns a list of neighbors for a given node in the graph.
func (g *Graph[T]) GetNeighbors(value T) []T {
	n := unique.Make(value)

	if _, ok := g.edges[n]; !ok {
		return nil
	}

	neighbors := make([]T, 0)
	for neighbor := range g.edges[n] {
		neighbors = append(neighbors, neighbor.Value())
	}

	return neighbors
}

/*
* Implement a generic memoization function that uses the unique
* package to efficiently cache function results for complex input types.
*
* Test with:
* fibonacci := Memoize(func(n int) int {
*     if n <= 1 {
*         return n
*     }
*     return fibonacci(n-1) + fibonacci(n-2)
* })
*
* fmt.Println(fibonacci(40))
* fmt.Println(fibonacci(40))  // Should be much faster
 */

func Memoize[T comparable, R any](f func(T) R) func(T) R {

	cache := make(map[unique.Handle[T]]R)

	return func(input T) R {
		hand := unique.Make(input)
		if val, ok := cache[hand]; ok {
			return val
		}

		result := f(input)
		cache[hand] = result
		return result
	}
}
