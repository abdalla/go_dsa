package traversal

import (
	"fmt"

	"github.com/abdalla/go_dsa/tree"
)

func Levelorder(t *tree.Tree) (traversal []int) {
	if t == nil {
		return
	}

	queue := []*tree.Tree{t}

	for len(queue) > 0 {
		current := queue[0]
		traversal = append(traversal, current.Value)
		queue = queue[1:]

		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		if current.Right != nil {
			queue = append(queue, current.Right)
		}

	}

	return
}

func bfs() {
	t := tree.CreateBinaryTree()

	fmt.Println(Levelorder(t))
}
