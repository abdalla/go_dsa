package test

import (
	"reflect"
	"testing"

	"github.com/abdalla/go_dsa/tree"
	"github.com/abdalla/go_dsa/tree/traversal"
)

func TestBFS(t *testing.T) {

	bt := tree.CreateBinaryTree()

	got := traversal.Levelorder(bt)
	want := []int{3, 4, 5, 6, 7, 8, 9}

	if len(got) != len(want) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}

}
