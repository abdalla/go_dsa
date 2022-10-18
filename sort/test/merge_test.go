package test

import (
	"reflect"
	"testing"

	"github.com/abdalla/go_dsa/sort"
)

func TestMergeSortAlreadySorted(t *testing.T) {

	got := sort.MergeSort(alreadySorted)
	want := alreadySorted

	if len(got) != len(want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func TestMergeSortInversed(t *testing.T) {

	got := sort.MergeSort(inversed)
	want := inversedSorted

	if len(got) != len(want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestMergeSortRepeated(t *testing.T) {

	got := sort.MergeSort(repeated)
	want := repeatedSorted

	if len(got) != len(want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

}
