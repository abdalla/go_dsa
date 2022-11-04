package reverse_test

import (
	"testing"

	"github.com/abdalla/go_dsa/challenges/linkedlist/reverse"
)

func getData() *reverse.Node {
	node1 := &reverse.Node{
		Value: 1,
	}

	node2 := &reverse.Node{
		Value: 2,
	}

	node3 := &reverse.Node{
		Value: 3,
	}

	node4 := &reverse.Node{
		Value: 4,
	}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	return node1
}

func TestNode_Revert(t *testing.T) {
	type fields struct {
		Next  *reverse.Node
		Value int
	}

	tests := []struct {
		got  *reverse.Node
		want *reverse.Node
		name string
	}{
		{
			want: getData(),
			got:  getData(),
			name: "success",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.got = tt.got.Revert()

			if tt.want.Value == tt.got.Value {
				t.Errorf("want: %v got: %v", 4, tt.got.Value)
			}
		})
	}
}
