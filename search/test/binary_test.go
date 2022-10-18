package test

import (
	"testing"

	"github.com/abdalla/go_dsa/search"
)

func TestBinarySearch11(t *testing.T) {

	want, got := true, search.BinarySearch(numbers, 11)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBinarySearch1(t *testing.T) {

	want, got := false, search.BinarySearch(numbers, 1)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBinarySearch7(t *testing.T) {

	want, got := true, search.BinarySearch(numbers, 7)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBinarySearch23(t *testing.T) {

	want, got := true, search.BinarySearch(numbers, 23)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBinarySearch2(t *testing.T) {

	want, got := true, search.BinarySearch(numbers, 2)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBinarySearch30(t *testing.T) {

	want, got := true, search.BinarySearch(numbers, 30)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBinarySearch50(t *testing.T) {

	want, got := false, search.BinarySearch(numbers, 50)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBinarySearch10Negative(t *testing.T) {

	want, got := false, search.BinarySearch(numbers, -10)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBinarySearch100(t *testing.T) {

	want, got := false, search.BinarySearch(numbers, 100)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
