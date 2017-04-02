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

func GenerateVertexCover(g Graph) []VertexCover {
	vertices := GetVertices(g)
	vc := make([]VertexCover, 0)
	var iter func(int, []string)
	iter = func(i int, curVert []string) {
		if i >= len(vertices) {
			vc = append(vc, VertexCover{curVert[:]})
			return
		}
		iter(i+1, curVert)
		copy := append(curVert[:], vertices[i])
		iter(i+1, copy)
	}
	iter(0, make([]string, 0))
	return vc
}

func MinCover(g Graph, vcs []VertexCover) *VertexCover {
	minSize := 1000000
	var minimumVertexCover *VertexCover
	for i, vc := range vcs {
		if vc.Size() < minSize && vc.Covers(g) {
			minSize = vc.Size()
			minimumVertexCover = &vcs[i]
		}
	}
	return minimumVertexCover
}

type SpanningTree struct {
	Edges []Edge
}

func GenerateSpanningTree(g Graph, startVertex string) []SpanningTree {
	trees := make([]SpanningTree, 0)
	var iter func(string, map[string]struct{}, SpanningTree)
	iter = func(fromVertex string, usedVertices map[string]struct{}, currentTree SpanningTree) {
		edges := g.EdgesFrom(fromVertex)
		done := true
		for _,edge := range edges {
			opposite := edge.GetOpposite(v)
			if _,ok :=
		}
	}
	iter(startVertex, make(map[string]struct{}), SpanningTree{})
	return trees
}

func run(s string) {
	// g := toGraph(s)
	// vc := GenerateVertexCover(g)
}
