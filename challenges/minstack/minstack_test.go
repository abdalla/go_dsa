package minstack_test

import (
	"reflect"
	"testing"

	"github.com/abdalla/go_dsa/challenges/minstack"
)

func TestTop(t *testing.T) {
	ms := minstack.MinStack{
		Min:   []int{5, 3, 1},
		Stack: []int{5, 7, 8, 9, 3, 4, 1},
	}

	got := ms.Top()
	want := 1

	if got != want {
		t.Errorf("got %+v, wanted %+v", got, want)
	}
}

func TestGetMin(t *testing.T) {
	ms := minstack.MinStack{
		Min:   []int{5, 3},
		Stack: []int{5, 7, 8, 9, 3, 4},
	}

	got := ms.GetMin()
	want := 3

	if got != want {
		t.Errorf("got %+v, wanted %+v", got, want)
	}
}

func TestPopStackAndMin(t *testing.T) {
	ms := minstack.MinStack{
		Min:   []int{5, 3, 1},
		Stack: []int{5, 7, 8, 9, 3, 4, 1},
	}

	ms.Pop()
	wantStack := []int{5, 7, 8, 9, 3, 4}
	wantMin := []int{5, 3}

	if len(ms.Stack) != len(wantStack) {
		t.Errorf("got %+v, wanted %+v", ms.Stack, wantStack)
	}

	if !reflect.DeepEqual(ms.Stack, wantStack) {
		t.Errorf("got %+v, wanted %+v", ms.Stack, wantStack)
	}

	if len(ms.Min) != len(wantMin) {
		t.Errorf("got %+v, wanted %+v", ms.Min, wantMin)
	}

	if !reflect.DeepEqual(ms.Min, wantMin) {
		t.Errorf("got %+v, wanted %+v", ms.Min, wantMin)
	}
}

func TestPopStackOnly(t *testing.T) {
	ms := minstack.MinStack{
		Min:   []int{5, 3},
		Stack: []int{5, 7, 8, 9, 3, 4},
	}

	ms.Pop()
	wantStack := []int{5, 7, 8, 9, 3}
	wantMin := []int{5, 3}

	if len(ms.Stack) != len(wantStack) {
		t.Errorf("got %+v, wanted %+v", ms.Stack, wantStack)
	}

	if !reflect.DeepEqual(ms.Stack, wantStack) {
		t.Errorf("got %+v, wanted %+v", ms.Stack, wantStack)
	}

	if len(ms.Min) != len(wantMin) {
		t.Errorf("got %+v, wanted %+v", ms.Min, wantMin)
	}

	if !reflect.DeepEqual(ms.Min, wantMin) {
		t.Errorf("got %+v, wanted %+v", ms.Min, wantMin)
	}
}

func TestPushStackAndMin(t *testing.T) {
	ms := minstack.MinStack{
		Min:   []int{5, 3},
		Stack: []int{5, 7, 8, 9, 3, 4},
	}

	ms.Push(1)
	wantStack := []int{5, 7, 8, 9, 3, 4, 1}
	wantMin := []int{5, 3, 1}

	if len(ms.Stack) != len(wantStack) {
		t.Errorf("got %+v, wanted %+v", ms.Stack, wantStack)
	}

	if !reflect.DeepEqual(ms.Stack, wantStack) {
		t.Errorf("got %+v, wanted %+v", ms.Stack, wantStack)
	}

	if len(ms.Min) != len(wantMin) {
		t.Errorf("got %+v, wanted %+v", ms.Min, wantMin)
	}

	if !reflect.DeepEqual(ms.Min, wantMin) {
		t.Errorf("got %+v, wanted %+v", ms.Min, wantMin)
	}
}

func TestPushStackOnly(t *testing.T) {
	ms := minstack.MinStack{
		Min:   []int{5, 3},
		Stack: []int{5, 7, 8, 9, 3, 4},
	}

	ms.Push(10)
	wantStack := []int{5, 7, 8, 9, 3, 4, 10}
	wantMin := []int{5, 3}

	if len(ms.Stack) != len(wantStack) {
		t.Errorf("got %+v, wanted %+v", ms.Stack, wantStack)
	}

	if !reflect.DeepEqual(ms.Stack, wantStack) {
		t.Errorf("got %+v, wanted %+v", ms.Stack, wantStack)
	}

	if len(ms.Min) != len(wantMin) {
		t.Errorf("got %+v, wanted %+v", ms.Min, wantMin)
	}

	if !reflect.DeepEqual(ms.Min, wantMin) {
		t.Errorf("got %+v, wanted %+v", ms.Min, wantMin)
	}
}
