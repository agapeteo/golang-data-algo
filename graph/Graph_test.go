package graph

import (
	"sort"
	"testing"
)

func TestGraph_ElementsDfs(t *testing.T) {
	g := createTestGraph()

	expected := []string{"a", "b", "c", "d", "e"}
	result := g.ElementsDfs()
	actual := make([]string, len(result))
	for i := range result {
		actual[i] = result[i].(string)
	}

	sort.Strings(actual)

	if !equalStringSlices(expected, actual) {
		t.Errorf("actual: %v is different than expected: %v", actual, expected)
	}
}

func TestGraph_ElementsBfs(t *testing.T) {
	g := createTestGraph()

	expected := []string{"a", "b", "c", "d", "e"}
	result := g.ElementsBfs()
	actual := make([]string, len(result))
	for i := range result {
		actual[i] = result[i].(string)
	}

	sort.Strings(actual)

	if !equalStringSlices(expected, actual) {
		t.Errorf("actual: %v is different than expected: %v", actual, expected)
	}
}

func TestGraph_IsConnected(t *testing.T) {
	g := createTestGraph()

	if !g.IsConnected("a", "e") {
		t.Errorf("'a' and 'e' should be connected")
	}

	g.AddBothEdges("y", "z")

	if g.IsConnected("a", "z") {
		t.Errorf("'a' and 'z' should not be connected")
	}
}

func equalStringSlices(a []string, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func createTestGraph() *Graph {
	g := NewGraph()

	g.AddBothEdges("a", "b")
	g.AddBothEdges("b", "c")
	g.AddBothEdges("b", "d")
	g.AddBothEdges("c", "e")
	g.AddBothEdges("d", "e")

	return g
}

