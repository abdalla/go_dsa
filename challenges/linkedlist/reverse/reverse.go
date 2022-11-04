package reverse

type Node struct {
	Next  *Node
	Value int
}

func (ll *Node) Revert() *Node {
	var prev *Node
	var next *Node
	cur := ll

	for cur != nil {
		next = cur.Next
		cur.Next = prev

		prev = cur
		cur = next
	}

	ll = prev

	return ll
}
