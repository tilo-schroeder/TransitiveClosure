//Further ideas: Topological Sort implementieren

package main

import (
	"fmt"
	"sync"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s", name, elapsed)
}

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it already exists in the graph", k)
		fmt.Println(err.Error())
	}
	g.vertices = append(g.vertices, &Vertex{key: k})
}

func (g *Graph) AddEdge(u int, v int) {
	fromVertex := g.getVertex(u)
	toVertex := g.getVertex(v)

	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v --> %v)", u, v)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, v) {
		err := fmt.Errorf("Edge (%v --> %v) not added because it already exists in the graph", u, v)
		fmt.Println(err.Error())
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf("%v ", v.key)
		}
	}
}

func (g *Graph) DFS(start int) []*Vertex {
	var v *Vertex
	v = g.getVertex(start)

	var discovered []*Vertex

	var stack []*Vertex
	stack = append(stack, v)

	for len(stack) > 0 {
		v = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !contains(discovered, v.key) {
			discovered = append(discovered, v)
			g.AddEdge(start, v.key) //Maybe create a new graph instead of trying to add the Edges (?)

			for _, w := range v.adjacent {
				if !contains(discovered, w.key) {
					stack = append(stack, w)
				}
			}
		}
	}
	return discovered
}

//Sequential execution of the Depth-First-Search (~1.5ms for a graph of 5)
func (g *Graph) sq_dfs_trc() {
	defer timeTrack(time.Now(), "sq_dfs_trc")
	for v := range g.vertices {
		g.DFS(v)
	}
}

//Execution using go-routines (~550 microseconds for a graph of 5)
func (g *Graph) dfs_trc() {
	defer timeTrack(time.Now(), "dfs_trc")

	var wg sync.WaitGroup
	wg.Add(len(g.vertices))

	for v := range g.vertices {
		go func(v int) {
			defer wg.Done()
			g.DFS(v)
		}(v)
	}
	wg.Wait()
}

func main() {
	test := &Graph{}
	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}
	test.AddEdge(1, 0)
	test.AddEdge(2, 1)
	test.AddEdge(0, 2)
	test.AddEdge(0, 3)
	test.AddEdge(3, 4)

	test.Print()
	fmt.Println()
	test.sq_dfs_trc()
	test.dfs_trc()
	test.Print()
}
