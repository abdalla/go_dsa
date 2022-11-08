package tree_test

import (
	"testing"

	tr "github.com/abdalla/go_dsa/challenges/tree"
	"github.com/abdalla/go_dsa/tree"
)

func getTree() *tree.Tree {
	node1 := &tree.Tree{Value: 1}
	node2 := &tree.Tree{Value: 2}
	node3 := &tree.Tree{Value: 3}
	node4 := &tree.Tree{Value: 4}
	node5 := &tree.Tree{Value: 5}
	node6 := &tree.Tree{Value: 6}
	node7 := &tree.Tree{Value: 7}

	node2.Left = node4
	node2.Right = node5

	node3.Left = node6
	node3.Right = node7

	node1.Left = node2
	node1.Right = node3

	return node1
}

func getInvertedTree() *tree.Tree {
	node1 := &tree.Tree{Value: 1}
	node2 := &tree.Tree{Value: 2}
	node3 := &tree.Tree{Value: 3}
	node4 := &tree.Tree{Value: 4}
	node5 := &tree.Tree{Value: 5}
	node6 := &tree.Tree{Value: 6}
	node7 := &tree.Tree{Value: 7}

	node2.Left = node5
	node2.Right = node4

	node3.Left = node7
	node3.Right = node6

	node1.Left = node3
	node1.Right = node2

	return node1
}

func TestInvertBT(t *testing.T) {
	tree := getTree()

	got := tr.InvertBT(tree)
	want := getInvertedTree()

	if got.Value != want.Value {
		t.Errorf("got %+v, wanted %+v", got.Value, want.Value)
	}

	if got.Left.Value != want.Left.Value {
		t.Errorf("got %+v, wanted %+v", got.Left.Right.Value, want.Left.Right.Value)
	}
}
