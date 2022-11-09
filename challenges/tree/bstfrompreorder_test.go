// Preorder: [5, 3, 1, 4, 8, 6, 9]
// Result:
//						5
//					/   \
//				 3     8
//				/ \   / \
//	 		 1   4 6 . 9

package tree_test

import (
	"fmt"
	"reflect"
	"testing"

	tr "github.com/abdalla/go_dsa/challenges/tree"
	"github.com/abdalla/go_dsa/tree"
)

func getWanted() *tree.Tree {
	root := &tree.Tree{Value: 5}
	node3 := &tree.Tree{Value: 3}
	node1 := &tree.Tree{Value: 1}
	node4 := &tree.Tree{Value: 4}
	node8 := &tree.Tree{Value: 8}
	node6 := &tree.Tree{Value: 6}
	node9 := &tree.Tree{Value: 9}

	node3.Left = node1
	node3.Right = node4

	node8.Left = node6
	node8.Right = node9

	root.Left = node3
	root.Right = node8

	return root
}

func TestBSTFromPreorder(t *testing.T) {
	porder := []int{5, 3, 1, 4, 8, 6, 9}

	got := tr.BstFromPreorder(porder)
	want := getWanted()

	fmt.Printf("got: %+v want: %+v \n", got, want)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}
}
