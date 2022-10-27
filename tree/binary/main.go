package main

import (
	"fmt"
)

type Tree struct {
	Value interface{}
	Left  *Tree
	Right *Tree
}

func main() {
	t := &Tree{
		Value: 1,

		Left: &Tree{
			Value: 2,
			Left:  nil,
			Right: &Tree{
				Value: 3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &Tree{
			Value: 4,
			Left: &Tree{
				Value: 5,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}

	currentNode := t
	nodeBefore := t
	aux := t
	isleft := true

	resolved := []int{t.Value.(int)}

	for i := 1; i <= 5; i++ {
		// if currentNode != nil {
		if currentNode == nil {
			break
		}

		if currentNode.Left != nil {
			resolved = append(resolved, currentNode.Left.Value.(int))
		}

		if currentNode.Right != nil {
			resolved = append(resolved, currentNode.Right.Value.(int))
		}

		if isleft {
			currentNode = aux.Left
			isleft = false
		} else {
			currentNode = aux.Right
			isleft = true
		}

		aux = nodeBefore
		nodeBefore = currentNode
	}
	// }

	fmt.Println(resolved)
}
