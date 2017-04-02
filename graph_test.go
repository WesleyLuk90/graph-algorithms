package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToGraph(t *testing.T) {
	g := toGraph("a,b;c,d")
	assert.Equal(t, len(g.Edges), 2)
	assert.Equal(t, g.Edges[0].A, "a")
	assert.Equal(t, g.Edges[0].B, "b")
	assert.Equal(t, g.Edges[1].A, "c")
	assert.Equal(t, g.Edges[1].B, "d")
}

func TestVertexCover(t *testing.T) {
	g := toGraph("a,b;b,c;c,d;a,d")
	vcs := GenerateVertexCover(g)
	assert.Equal(t, len(vcs), 16)
}

func TestCovers(t *testing.T) {
	testCases := []struct {
		graph  string
		vc     VertexCover
		passes bool
	}{
		{
			"a,b;b,c;c,d;a,d",
			VertexCover{[]string{"a", "c"}},
			true,
		},
		{
			"a,b;b,c;c,d;a,d",
			VertexCover{[]string{"a", "b"}},
			false,
		},
		{
			"a,b;b,c;c,d;a,d;d,e",
			VertexCover{[]string{"a", "c"}},
			false,
		},
		{
			"a,b;b,c;c,d;a,d;d,e",
			VertexCover{[]string{"b", "d"}},
			true,
		},
	}
	for _, testCase := range testCases {
		graph := toGraph(testCase.graph)
		assert.Equal(t, testCase.passes, testCase.vc.Covers(graph), fmt.Sprintf("Expected %v# %v#\n", testCase.graph, testCase.vc))
	}
}

func TestMinCover(t *testing.T) {
	g := toGraph("a,b;b,c;c,d;a,d")
	min := MinCover(g, GenerateVertexCover(g))
	assert.Equal(t, 2, min.Size())
}

func TestMinCover2(t *testing.T) {
	g := toGraph("a,b;b,c;c,d;d,e;e,f;f,g;g,h;h,i")
	vertexCovers := GenerateVertexCover(g)
	assert.Equal(t, 512, len(vertexCovers))
	min := MinCover(g, vertexCovers)
	assert.Equal(t, 4, min.Size())
}

func TestSimpleGenerateSpanningTree(t *testing.T) {
	g := toGraph("a,b;b,c;c,d")
	spanningTrees := GenerateSpanningTree(g, "a")
	assert.Equal(t, 1, len(spanningTrees))
}

func TestComplexGenerateSpanningTree(t *testing.T) {
	g := toGraph("a,b;b,c;c,d")
	spanningTrees := GenerateSpanningTree(g, "b")
	assert.Equal(t, 2, len(spanningTrees))
}

func TestSizeWithoutLeaves(t *testing.T) {
	tree := SpanningTree{[]Edge{
		Edge{"a", "b"},
		Edge{"a", "c"},
		Edge{"b", "d"},
	}}
	assert.Equal(t, 2, tree.SizeWithoutLeaves("a"))
}

func TestHasSpanningTreeCoverWithCount(t *testing.T) {
	g := toGraph("a,b;b,c;c,d")
	min := MinCover(g, GenerateVertexCover(g))
	tree := GetSpanningTreeCoverWithCount(g, min.Size()*2)
	assert.Nil(t, tree)
}

func TestSquares1(t *testing.T) {
	for tailSize := 1; tailSize <= 5; tailSize += 2 {
		g := toGraph("a,b;b,c;c,d;d,a;a,t1")
		for i := 1; i < tailSize; i++ {
			g.Edges = append(g.Edges, Edge{fmt.Sprintf("t%d", i), fmt.Sprintf("t%d", (i + 1))})
		}
		// fmt.Printf("Graph %v#\n", g)
		min := MinCover(g, GenerateVertexCover(g))
		// fmt.Printf("MVC is %d\n%v#\n", min.Size(), min)
		assert.Equal(t, min.Size(), 2+tailSize/2)
		tree := GetSpanningTreeCoverWithCount(g, min.Size()*2)
		assert.NotNil(t, tree)
	}
}

func TestLine(t *testing.T) {
	for tailSize := 3; tailSize <= 21; tailSize += 2 {
		g := Graph{}
		for i := 1; i < tailSize; i++ {
			g.Edges = append(g.Edges, Edge{fmt.Sprintf("t%d", i), fmt.Sprintf("t%d", (i + 1))})
		}
		fmt.Printf("Graph %v#\n", g)
		min := MinCover(g, GenerateVertexCover(g))
		fmt.Printf("MVC is %d\n%v#\n", min.Size(), min)
		tree := GetSpanningTreeCoverWithCount(g, min.Size()*2)
		assert.NotNil(t, tree, fmt.Sprintf("Expected cover of %d in %v", min.Size()*2, g))
	}
}
