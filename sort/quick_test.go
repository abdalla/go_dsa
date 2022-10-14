package sort_test

import (
	"reflect"
	"testing"

	"github.com/abdalla/go_dsa/sort"
)

func TestQuickSortAlreadySorted(t *testing.T) {

	got := sort.QuickSort(alreadySorted, 0, len(alreadySorted)-1)
	want := alreadySorted

	if len(got) != len(want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func TestQuickSortInversed(t *testing.T) {

	got := sort.QuickSort(inversed, 0, len(inversed)-1)
	want := inversedSorted

	if len(got) != len(want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestQuickSortRepeated(t *testing.T) {

	got := sort.QuickSort(repeated, 0, len(repeated)-1)
	want := repeatedSorted

	if len(got) != len(want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

}
