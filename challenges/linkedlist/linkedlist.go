package linkedlist

type Node struct {
	Previus *Node
	Next    *Node
	Value   int
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

// [5, 6, 7, 8]
//GET INDEX
//ADD HEAD
//ADD TAIL
//ADD INDEX
//DELETE INDEX

func (ll *LinkedList) GetNode(index int) *Node {
	if index < 0 || index > ll.Size {
		return nil
	}

	current := ll.Head

	for index != 0 {
		current = current.Next
		index = index - 1
	}

	return current
}

func (ll *LinkedList) Get(index int) int {
	current := ll.GetNode(index)

	if current == nil {
		return -1
	}

	return current.Value
}

func (ll *LinkedList) AddHead(value int) {
	newNode := &Node{
		Value: value,
		Next:  ll.Head,
	}

	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
		ll.Size += 1

		return
	}

	// newNode.Next = ll.Head
	ll.Head.Previus = newNode
	ll.Head = newNode
	ll.Size += 1
}

func (ll *LinkedList) AddTail(value int) {
	newNode := &Node{
		Value:   value,
		Previus: ll.Tail,
	}

	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
		ll.Size += 1

		return
	}

	// newNode.Previus = ll.Tail
	ll.Tail.Next = newNode
	ll.Tail = newNode
	ll.Size += 1
}

func (ll *LinkedList) Add(index, value int) {
	if index > ll.Size || index < 0 {
		return
	}

	if index == 0 {
		ll.AddHead(value)
		return
	}

	if index == ll.Size {
		ll.AddTail(value)
		return
	}

	current := ll.GetNode(index)
	if current == nil {
		return
	}

	newNode := &Node{
		Value:   value,
		Next:    current,
		Previus: current.Previus,
	}

	current.Previus.Next = newNode
	current.Previus = newNode
	// current.Next = newNode
	// newNode.Previus = current

	ll.Size++
}
