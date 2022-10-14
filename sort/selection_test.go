package sort_test

import (
	"reflect"
	"testing"

	"github.com/abdalla/go_dsa/sort"
)

func TestSelectionSortAlreadySorted(t *testing.T) {

	got := sort.SelectionSort(alreadySorted)
	want := alreadySorted

	if len(got) != len(want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func TestSelectionSortInversed(t *testing.T) {

	got := sort.SelectionSort(inversed)
	want := inversedSorted

	if len(got) != len(want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestSelectionSortRepeated(t *testing.T) {

	got := sort.SelectionSort(repeated)
	want := repeatedSorted

	if len(got) != len(want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

}
