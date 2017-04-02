package main

import "strings"

type Edge struct {
	A string
	B string
}

func (e *Edge) GetOpposite(v string) string {
	if e.A == v {
		return e.B
	}
	if e.B == v {
		return e.A
	}
	panic("Invalid vertex")
}

type Graph struct {
	Edges []Edge
}

func (g *Graph) AddEdge(a, b string) {
	g.Edges = append(g.Edges, Edge{a, b})
}

func (g *Graph) EdgesFrom(a string) []Edge {
	edges := make([]Edge, 0)
	for _, e := range g.Edges {
		if e.A == a || e.B == a {
			edges = append(edges, e)
		}
	}
	return edges
}

func (g *Graph) Vertices() []string {
	vertices := make(VertexSet)
	for _, e := range g.Edges {
		vertices.Add(e.A)
		vertices.Add(e.B)
	}
	return vertices.Vertices()
}

type VertexCover struct {
	Vertices []string
}

func (vc *VertexCover) HasVertex(v string) bool {
	for _, vertex := range vc.Vertices {
		if vertex == v {
			return true
		}
	}
	return false
}

func (v *VertexCover) Covers(g Graph) bool {
	for _, edge := range g.Edges {
		if !v.HasVertex(edge.A) && !v.HasVertex(edge.B) {
			return false
		}
	}
	return true
}

func (v *VertexCover) Size() int {
	return len(v.Vertices)
}

func toGraph(s string) Graph {
	g := Graph{}
	edgeData := strings.Split(s, ";")
	for _, e := range edgeData {
		edgePair := strings.Split(e, ",")
		g.AddEdge(edgePair[0], edgePair[1])
	}
	return g
}

func GetVertices(g Graph) []string {
	vertices := make(map[string]struct{})
	for _, edge := range g.Edges {
		vertices[edge.A] = struct{}{}
		vertices[edge.B] = struct{}{}
	}
	verticesArray := make([]string, 0)
	for edge := range vertices {
		verticesArray = append(verticesArray, edge)
	}
	return verticesArray
}

func CopyVertices(vertices []string) []string {
	c := make([]string, len(vertices))
	copy(c, vertices)
	return c
}

func GenerateVertexCover(g Graph) []VertexCover {
	vertices := GetVertices(g)
	vc := make([]VertexCover, 0)
	var iter func(int, []string)
	iter = func(i int, curVert []string) {
		if i >= len(vertices) {
			vc = append(vc, VertexCover{CopyVertices(curVert)})
			return
		}
		iter(i+1, CopyVertices(curVert))
		copy := append(CopyVertices(curVert), vertices[i])
		iter(i+1, copy)
	}
	iter(0, make([]string, 0))
	return vc
}

func MinCover(g Graph, vcs []VertexCover) *VertexCover {
	minSize := 1000000
	var minimumVertexCover VertexCover
	for _, vc := range vcs {
		if vc.Size() < minSize {
			if vc.Covers(g) {
				minSize = vc.Size()
				minimumVertexCover = vc
			}
		}
	}
	return &minimumVertexCover
}

type SpanningTree struct {
	Edges []Edge
}

func (st *SpanningTree) SizeWithoutLeaves(rootVertex string) int {
	vertexCount := make(map[string]int)
	for _, edge := range st.Edges {
		vertexCount[edge.A] += 1
		vertexCount[edge.B] += 1
	}
	total := 0
	for vertex, count := range vertexCount {
		if vertex == rootVertex || count > 1 {
			total += 1
		}
	}
	return total
}

type VertexSet map[string]struct{}

func (v VertexSet) Has(vertex string) bool {
	_, ok := v[vertex]
	return ok
}

func (v VertexSet) Add(vertex string) {
	v[vertex] = struct{}{}
}
func (v VertexSet) Remove(vertex string) {
	delete(v, vertex)
}

func (v VertexSet) Vertices() []string {
	vertices := make([]string, 0)
	for vertex := range v {
		vertices = append(vertices, vertex)
	}
	return vertices
}

func (v VertexSet) Count() int {
	return len(v.Vertices())
}

type EdgeSet map[Edge]struct{}

func (es EdgeSet) Has(edge Edge) bool {
	_, ok := es[edge]
	return ok
}

func (es EdgeSet) Add(edge Edge) {
	es[edge] = struct{}{}
}
func (es EdgeSet) Remove(edge Edge) {
	delete(es, edge)
}

type VertexStack []string

func (vs *VertexStack) Push(vertex string) {
	*vs = append(*vs, vertex)
}
func (vs *VertexStack) Pop() string {
	index := len(*vs) - 1
	last := (*vs)[index]
	*vs = (*vs)[:index]
	return last
}
func (vs *VertexStack) Peak() string {
	index := len(*vs) - 1
	last := (*vs)[index]
	return last
}

func NewSpanningTreeGenerator(g Graph) SpanningTreeGenerator {
	return SpanningTreeGenerator{
		Graph:         g,
		UsedVertices:  make(VertexSet),
		UsedEdges:     make(EdgeSet),
		Trees:         make([]SpanningTree, 0),
		TotalVertices: len(g.Vertices()),
		Stack:         make(VertexStack, 0),
	}
}

type SpanningTreeGenerator struct {
	Graph         Graph
	UsedVertices  VertexSet
	UsedEdges     EdgeSet
	Trees         []SpanningTree
	TotalVertices int
	Stack         VertexStack
	Count         int
}

func CopyAppendEdges(edges []Edge, edge Edge) []Edge {
	c := make([]Edge, len(edges)+1)
	copy(c, edges)
	c[len(edges)] = edge
	return c
}

func (g *SpanningTreeGenerator) IterDown(fromVertex string, currentTree SpanningTree) {
	g.UsedVertices.Add(fromVertex)
	g.Stack.Push(fromVertex)

	if !g.CheckDone(currentTree) {
		edges := g.Graph.EdgesFrom(fromVertex)
		noMore := true
		for _, edge := range edges {
			opposite := edge.GetOpposite(fromVertex)
			if !g.UsedVertices.Has(opposite) && !g.UsedEdges.Has(edge) {
				g.UsedEdges.Add(edge)
				g.IterDown(opposite, SpanningTree{CopyAppendEdges(currentTree.Edges, edge)})
				g.UsedEdges.Remove(edge)
				noMore = false
			}
		}
		if noMore {
			g.IterUp(currentTree)
		}
	}
	g.Stack.Pop()
	g.UsedVertices.Remove(fromVertex)
}

func (g *SpanningTreeGenerator) IterUp(currentTree SpanningTree) {
	if len(g.Stack) < 2 {
		return
	}
	last2 := g.Stack.Pop()
	last1 := g.Stack.Pop()
	g.IterDown(last1, currentTree)
	g.Stack.Push(last1)
	g.Stack.Push(last2)
}

func (g *SpanningTreeGenerator) CheckDone(currentTree SpanningTree) bool {
	if g.UsedVertices.Count() == g.TotalVertices {
		g.Trees = append(g.Trees, currentTree)
		return true
	}
	return false
}

func GenerateSpanningTree(g Graph, startVertex string) []SpanningTree {
	generator := NewSpanningTreeGenerator(g)
	generator.IterDown(startVertex, SpanningTree{})
	return generator.Trees
}

func GetSpanningTreeCoverWithCount(g Graph, count int) *SpanningTree {
	vertices := g.Vertices()
	for _, vertex := range vertices {
		trees := GenerateSpanningTree(g, vertex)
		for _, tree := range trees {
			if tree.SizeWithoutLeaves(vertex) == count {
				return &tree
			}
		}
	}
	return nil
}

func run(s string) {
	// g := toGraph(s)
	// vc := GenerateVertexCover(g)
}
