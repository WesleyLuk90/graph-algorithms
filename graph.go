package main

import "strings"

type Edge struct {
	A string
	B string
}

type Graph struct {
	Edges []Edge
}

func (g *Graph) AddEdge(a, b string) {
	g.Edges = append(g.Edges, Edge{a, b})
}

type VertexCover struct {
	Vertices []string
}

func (v *VertexCover) Covers(g Graph) bool {
	return false
}

func (v *VertexCover) Size() int {
	return 0
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

func run(s string) {
	// g := toGraph(s)
	// vc := GenerateVertexCover(g)
}
