package graph_test

import (
	"reflect"
	"testing"

	"github.com/abdalla/go_dsa/graph"
)

// DON'T LOOK AT IT AS TREE... THIS IS A GRAPH DESIGNED THIS WAY JUST FOR VISUALIZATION
//
//	     A
//	   / |  \
//	  B  C   D
//	 / \    /
//	E   F  G
func createInitialGraph() *graph.Graph {
	node1 := &graph.Graph{
		Value: "A",
	}

	node2 := &graph.Graph{
		Value: "B",
	}

	node3 := &graph.Graph{
		Value: "C",
	}

	node4 := &graph.Graph{
		Value: "D",
	}

	node5 := &graph.Graph{
		Value: "E",
	}

	node6 := &graph.Graph{
		Value: "F",
	}

	node7 := &graph.Graph{
		Value: "G",
	}

	node1.AdjacentList = append(node1.AdjacentList, node2, node3, node4)
	node2.AdjacentList = append(node2.AdjacentList, node5, node6)
	node3.AdjacentList = append(node4.AdjacentList, node1)
	node4.AdjacentList = append(node4.AdjacentList, node7)

	return node1
}

func TestBFSGraph(t *testing.T) {
	g := createInitialGraph()

	got := graph.BFSRetrieveTraversal(g)
	want := []string{"A", "B", "C", "D", "E", "F", "G"}

	if len(got) != len(want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func TestDFSGraph(t *testing.T) {
	g := createInitialGraph()

	got := graph.DFSRetrieveTraversal(g)
	want := []string{"A", "B", "E", "F", "C", "D", "G"}

	if len(got) != len(want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

}
