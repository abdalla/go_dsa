package maxstack_test

import (
	"reflect"
	"testing"

	"github.com/abdalla/go_dsa/challenges/maxstack"
)

func TestTop(t *testing.T) {
	ms := maxstack.MaxStack{
		Max:   []int{1, 3, 5},
		Stack: []int{5, 7, 8, 9, 3, 4, 1},
	}

	got := ms.Top()
	want := 1

	if got != want {
		t.Errorf("got %+v, wanted %+v", got, want)
	}
}

func TestGetMax(t *testing.T) {
	ms := maxstack.MaxStack{
		Max:   []int{5, 7, 8, 9},
		Stack: []int{5, 7, 8, 9, 3, 4},
	}

	got := ms.GetMax()
	want := 9

	if got != want {
		t.Errorf("got %+v, wanted %+v", got, want)
	}
}

func TestPopStackAndMax(t *testing.T) {
	ms := maxstack.MaxStack{
		Max:   []int{5, 7, 8, 9},
		Stack: []int{5, 7, 8, 1, 3, 4, 9},
	}

	ms.Pop()
	wantStack := []int{5, 7, 8, 1, 3, 4}
	wantMax := []int{5, 7, 8}

	if len(ms.Stack) != len(wantStack) {
		t.Errorf("got %+v, wanted %+v", ms.Stack, wantStack)
	}

	if !reflect.DeepEqual(ms.Stack, wantStack) {
		t.Errorf("got %+v, wanted %+v", ms.Stack, wantStack)
	}

	if len(ms.Max) != len(wantMax) {
		t.Errorf("got %+v, wanted %+v", ms.Max, wantMax)
	}

	if !reflect.DeepEqual(ms.Max, wantMax) {
		t.Errorf("got %+v, wanted %+v", ms.Max, wantMax)
	}
}

func TestPopStackOnly(t *testing.T) {
	ms := maxstack.MaxStack{
		Max:   []int{5, 7, 8},
		Stack: []int{5, 7, 8, 1, 3, 4},
	}

	ms.Pop()
	wantStack := []int{5, 7, 8, 1, 3}
	wantMax := []int{5, 7, 8}

	if len(ms.Stack) != len(wantStack) {
		t.Errorf("got %+v, wanted %+v", ms.Stack, wantStack)
	}

	if !reflect.DeepEqual(ms.Stack, wantStack) {
		t.Errorf("got %+v, wanted %+v", ms.Stack, wantStack)
	}

	if len(ms.Max) != len(wantMax) {
		t.Errorf("got %+v, wanted %+v", ms.Max, wantMax)
	}

	if !reflect.DeepEqual(ms.Max, wantMax) {
		t.Errorf("got %+v, wanted %+v", ms.Max, wantMax)
	}
}

func TestPushStackAndMax(t *testing.T) {
	ms := maxstack.MaxStack{
		Max:   []int{5, 7, 8},
		Stack: []int{5, 7, 8, 1, 3, 4},
	}

	ms.Push(9)
	wantStack := []int{5, 7, 8, 1, 3, 4, 9}
	wantMax := []int{5, 7, 8, 9}

	if len(ms.Stack) != len(wantStack) {
		t.Errorf("got %+v, wanted %+v", ms.Stack, wantStack)
	}

	if !reflect.DeepEqual(ms.Stack, wantStack) {
		t.Errorf("got %+v, wanted %+v", ms.Stack, wantStack)
	}

	if len(ms.Max) != len(wantMax) {
		t.Errorf("got %+v, wanted %+v", ms.Max, wantMax)
	}

	if !reflect.DeepEqual(ms.Max, wantMax) {
		t.Errorf("got %+v, wanted %+v", ms.Max, wantMax)
	}
}

func TestPushStackOnly(t *testing.T) {
	ms := maxstack.MaxStack{
		Max:   []int{5, 7, 8, 9},
		Stack: []int{5, 7, 8, 9, 3, 4},
	}

	ms.Push(1)
	wantStack := []int{5, 7, 8, 9, 3, 4, 1}
	wantMax := []int{5, 7, 8, 9}

	if len(ms.Stack) != len(wantStack) {
		t.Errorf("got %+v, wanted %+v", ms.Stack, wantStack)
	}

	if !reflect.DeepEqual(ms.Stack, wantStack) {
		t.Errorf("got %+v, wanted %+v", ms.Stack, wantStack)
	}

	if len(ms.Max) != len(wantMax) {
		t.Errorf("got %+v, wanted %+v", ms.Max, wantMax)
	}

	if !reflect.DeepEqual(ms.Max, wantMax) {
		t.Errorf("got %+v, wanted %+v", ms.Max, wantMax)
	}
}
