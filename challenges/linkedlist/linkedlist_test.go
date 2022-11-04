package linkedlist_test

import (
	"testing"

	"github.com/abdalla/go_dsa/challenges/linkedlist"
)

func createNode(value int) *linkedlist.Node {
	return &linkedlist.Node{
		Value: value,
	}
}

func getInitialData() *linkedlist.LinkedList {
	node1 := createNode(1)
	node2 := createNode(2)
	node3 := createNode(3)
	node4 := createNode(4)
	node5 := createNode(5)
	node6 := createNode(6)

	node1.Next = node2

	node2.Previus = node1
	node2.Next = node3

	node3.Previus = node2
	node3.Next = node4

	node4.Previus = node3
	node4.Next = node5

	node5.Previus = node4
	node5.Next = node6

	node6.Previus = node5

	ll := &linkedlist.LinkedList{
		Head: node1,
		Tail: node6,
		Size: 6,
	}

	return ll
}

func getWantedWithNewHead(value int) *linkedlist.LinkedList {
	want := getInitialData()
	newNode := createNode(value)
	newNode.Next = want.Head
	want.Head.Previus = newNode
	want.Head = newNode
	want.Size += 1

	return want
}

func getWantedWithNewTail(value int) *linkedlist.LinkedList {
	want := getInitialData()
	newNode := createNode(value)
	newNode.Previus = want.Tail
	want.Tail.Next = newNode
	want.Tail = newNode
	want.Size += 1

	return want
}

func TestGetSuccess(t *testing.T) {

	ll := getInitialData()

	want := 3
	got := ll.Get(2)

	if want != got {
		t.Errorf("want: %v got: %v", want, got)
	}

}

func TestGetFailure(t *testing.T) {

	ll := getInitialData()

	want := 3
	got := ll.Get(1)

	if want == got {
		t.Errorf("want: %v got: %v", want, got)
	}

}

func TestAddHead(t *testing.T) {
	ll := getInitialData()
	want := getWantedWithNewHead(0)

	ll.AddHead(0)

	if want.Size != ll.Size {
		t.Errorf("want: %v got: %v", want.Size, ll.Size)
	}

	if want.Head.Value != ll.Head.Value {
		t.Errorf("want: %v got: %v", want.Head.Value, ll.Head.Value)
	}

}

func TestAddTail(t *testing.T) {
	ll := getInitialData()
	want := getWantedWithNewTail(7)

	ll.AddTail(7)

	if want.Size != ll.Size {
		t.Errorf("want: %v got: %v", want.Size, ll.Size)
	}

	if want.Tail.Value != ll.Tail.Value {
		t.Errorf("want: %v got: %v", want.Tail.Value, ll.Tail.Value)
	}

}

func TestAddInvalidIndexGreaterThanSize(t *testing.T) {
	ll := getInitialData()
	want := getInitialData()

	ll.Add(7, 7)

	if want.Size != ll.Size {
		t.Errorf("want: %v got: %v", want.Size, ll.Size)
	}

	if want.Head.Value != ll.Head.Value {
		t.Errorf("want: %v got: %v", want.Head.Value, ll.Head.Value)
	}

	if want.Tail.Value != ll.Tail.Value {
		t.Errorf("want: %v got: %v", want.Tail.Value, ll.Tail.Value)
	}

}

func TestAddInvalidIndexLessThanZero(t *testing.T) {
	ll := getInitialData()
	want := getInitialData()

	ll.Add(-1, 7)

	if want.Size != ll.Size {
		t.Errorf("want: %v got: %v", want.Size, ll.Size)
	}

	if want.Head.Value != ll.Head.Value {
		t.Errorf("want: %v got: %v", want.Head.Value, ll.Head.Value)
	}

	if want.Tail.Value != ll.Tail.Value {
		t.Errorf("want: %v got: %v", want.Tail.Value, ll.Tail.Value)
	}

}

func TestAdd(t *testing.T) {
	index := 3
	value := 7

	ll := getInitialData()
	want := getInitialData()

	ll.Add(index, value)

	if want.Size == ll.Size {
		t.Errorf("want: %v got: %v", want.Size, ll.Size)
	}

	if want.Head.Value != ll.Head.Value {
		t.Errorf("want: %v got: %v", want.Head.Value, ll.Head.Value)
	}

	if want.Tail.Value != ll.Tail.Value {
		t.Errorf("want: %v got: %v", want.Tail.Value, ll.Tail.Value)
	}

	// currentWanted := want.GetNode(index)
	currentGot := ll.GetNode(index)

	if currentGot.Value != value {
		t.Errorf("want: %v got: %v", value, currentGot.Value)
	}

}

func TestDeleteInvalidIndexGreaterThanSize(t *testing.T) {
	ll := getInitialData()
	want := getInitialData()

	ll.Delete(7)

	if want.Size != ll.Size {
		t.Errorf("want: %v got: %v", want.Size, ll.Size)
	}

	if want.Head.Value != ll.Head.Value {
		t.Errorf("want: %v got: %v", want.Head.Value, ll.Head.Value)
	}

	if want.Tail.Value != ll.Tail.Value {
		t.Errorf("want: %v got: %v", want.Tail.Value, ll.Tail.Value)
	}

}

func TestDeleteInvalidIndexLessThanZero(t *testing.T) {
	ll := getInitialData()
	want := getInitialData()

	ll.Delete(-1)

	if want.Size != ll.Size {
		t.Errorf("want: %v got: %v", want.Size, ll.Size)
	}

	if want.Head.Value != ll.Head.Value {
		t.Errorf("want: %v got: %v", want.Head.Value, ll.Head.Value)
	}

	if want.Tail.Value != ll.Tail.Value {
		t.Errorf("want: %v got: %v", want.Tail.Value, ll.Tail.Value)
	}

}

func TestDeleteHead(t *testing.T) {
	ll := getInitialData()
	want := getInitialData()

	ll.Delete(0)

	if want.Size <= ll.Size {
		t.Errorf("want: %v got: %v", want.Size, ll.Size)
	}

	if want.Head.Value == ll.Head.Value {
		t.Errorf("want: %v got: %v", want.Head.Value, ll.Head.Value)
	}

	if want.Tail.Value != ll.Tail.Value {
		t.Errorf("want: %v got: %v", want.Tail.Value, ll.Tail.Value)
	}

	// currentWanted := want.GetNode(index)
	currentGot := ll.GetNode(0)

	if currentGot.Value != 2 {
		t.Errorf("want: %v got: %v", 2, currentGot.Value)
	}

}

func TestDeleteTail(t *testing.T) {
	ll := getInitialData()
	want := getInitialData()

	ll.Delete(6)

	if want.Size <= ll.Size {
		t.Errorf("want: %v got: %v", want.Size, ll.Size)
	}

	if want.Head.Value != ll.Head.Value {
		t.Errorf("want: %v got: %v", want.Head.Value, ll.Head.Value)
	}

	if want.Tail.Value == ll.Tail.Value {
		t.Errorf("want: %v got: %v", want.Tail.Value, ll.Tail.Value)
	}

	// currentWanted := want.GetNode(index)
	currentGot := ll.GetNode(6)
	if currentGot != nil {
		t.Errorf("want: %v got: %v", nil, currentGot.Value)
	}

	val := ll.Get(4)
	if val != 5 {
		t.Errorf("want: %v got: %v", 5, val)
	}

}

func TestDelete(t *testing.T) {
	index := 3

	ll := getInitialData()
	want := getInitialData()

	ll.Delete(index)

	if want.Size <= ll.Size {
		t.Errorf("want: %v got: %v", want.Size, ll.Size)
	}

	if want.Head.Value != ll.Head.Value {
		t.Errorf("want: %v got: %v", want.Head.Value, ll.Head.Value)
	}

	if want.Tail.Value != ll.Tail.Value {
		t.Errorf("want: %v got: %v", want.Tail.Value, ll.Tail.Value)
	}

	// currentWanted := want.GetNode(index)
	currentGot := ll.GetNode(index)
	if currentGot == nil {
		t.Errorf("want: %v got: %v", 4, currentGot.Value)
	}

	if currentGot.Value != 5 {
		t.Errorf("want: %v got: %v", 4, currentGot.Value)
	}

}
