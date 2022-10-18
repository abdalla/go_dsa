package test

import (
	"reflect"
	"testing"

	"github.com/abdalla/go_dsa/sort"
)

func TestInsertionSortAlreadySorted(t *testing.T) {

	got := sort.InsertionSort(alreadySorted)
	want := alreadySorted

	if len(got) != len(want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func TestInsertionSortInversed(t *testing.T) {

	got := sort.InsertionSort(inversed)
	want := inversedSorted

	if len(got) != len(want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestInsertionSortRepeated(t *testing.T) {

	got := sort.InsertionSort(repeated)
	want := repeatedSorted

	if len(got) != len(want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

}
