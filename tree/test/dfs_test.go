package test

import (
	"reflect"
	"testing"

	"github.com/abdalla/go_dsa/tree"
	"github.com/abdalla/go_dsa/tree/traversal"
)

func TestPreOrder(t *testing.T) {

	bt := tree.CreateBinaryTree()

	got := traversal.Preorder(bt)
	want := []int{3, 4, 6, 7, 5, 8, 9}

	if len(got) != len(want) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}

}

func TestInOrder(t *testing.T) {

	bt := tree.CreateBinaryTree()

	got := traversal.Inorder(bt)
	want := []int{6, 4, 7, 3, 8, 5, 9}

	if len(got) != len(want) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}

}

func TestPostOrder(t *testing.T) {

	bt := tree.CreateBinaryTree()

	got := traversal.Postorder(bt)
	want := []int{6, 7, 4, 8, 9, 5, 3}

	if len(got) != len(want) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}

}
