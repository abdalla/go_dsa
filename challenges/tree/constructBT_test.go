package tree_test

import (
	"testing"

	tr "github.com/abdalla/go_dsa/challenges/tree"
	"github.com/abdalla/go_dsa/tree"
)

var (
	preorder = []int{1, 2, 4, 8, 9, 5, 10, 11, 3, 6, 7} // root > left > right
	inorder  = []int{8, 4, 9, 2, 10, 5, 11, 1, 6, 3, 7} // left > root > right
	want     = &tree.Tree{}
)

func getWant() *tree.Tree {
	node1 := &tree.Tree{Value: 1}
	node2 := &tree.Tree{Value: 2}
	node3 := &tree.Tree{Value: 3}
	node4 := &tree.Tree{Value: 4}
	node5 := &tree.Tree{Value: 5}
	node6 := &tree.Tree{Value: 6}
	node7 := &tree.Tree{Value: 7}
	node8 := &tree.Tree{Value: 8}
	node9 := &tree.Tree{Value: 9}
	node10 := &tree.Tree{Value: 10}
	node11 := &tree.Tree{Value: 11}

	node4.Left = node8
	node4.Right = node9

	node5.Left = node10
	node5.Right = node11

	node2.Left = node4
	node2.Right = node5

	node3.Left = node6
	node3.Right = node7

	node1.Left = node2
	node1.Right = node3

	return node1
}

func TestBuildTree(t *testing.T) {

	got := tr.BuildTree(preorder, inorder)
	want := getWant()

	if got.Value != want.Value {
		t.Errorf("got %+v, wanted %+v", got.Value, want.Value)
	}

	if got.Left.Right.Value != want.Left.Right.Value {
		t.Errorf("got %+v, wanted %+v", got.Left.Right.Value, want.Left.Right.Value)
	}

	if got.Right.Left.Value != want.Right.Left.Value {
		t.Errorf("got %+v, wanted %+v", got.Right.Left.Value, want.Right.Left.Value)
	}

	if got.Left.Left.Value != want.Left.Left.Value {
		t.Errorf("got %+v, wanted %+v", got.Left.Left.Value, want.Left.Left.Value)
	}
}
