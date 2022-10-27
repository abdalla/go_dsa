package tree

type Tree struct {
	Left  *Tree
	Right *Tree
	Value int
}

func CreateBinaryTree() *Tree {
	t := &Tree{}

	t.Value = 3

	t.Left = &Tree{
		Value: 4,
	}

	t.Left.Left = &Tree{
		Value: 6,
	}

	t.Left.Right = &Tree{
		Value: 7,
	}

	t.Right = &Tree{
		Value: 5,
	}

	t.Right.Left = &Tree{
		Value: 8,
	}

	t.Right.Right = &Tree{
		Value: 9,
	}

	return t
}
