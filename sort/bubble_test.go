package sort_test

import (
	"reflect"
	"testing"

	"github.com/abdalla/go_dsa/sort"
)

func TestBubbleSortAlreadySorted(t *testing.T) {

	got := sort.BubbleSort(alreadySorted)
	want := alreadySorted

	if len(got) != len(want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func TestBubbleSortInversed(t *testing.T) {

	got := sort.BubbleSort(inversed)
	want := inversedSorted

	if len(got) != len(want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestBubbleSortRepeated(t *testing.T) {

	got := sort.BubbleSort(repeated)
	want := repeatedSorted

	if len(got) != len(want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

}
