// Input  : arr[] = {10, 20, 10, 5, 15}
// Output : prefixSum[] = {10, 30, 40, 45, 60}

//"Qual a soma dos elementos entre arr(1) e arr(3)?"

package realworld_test

import (
	"reflect"
	"testing"

	"github.com/abdalla/go_dsa/challenges/realworld"
)

func TestPrefixSum(t *testing.T) {
	arr := []int{10, 20, 10, 5, 15}

	want := []int{10, 30, 40, 45, 60}
	got := realworld.PrefixSum(arr)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}
}

func TestSum2Elements1and3(t *testing.T) {
	el1 := 1
	el2 := 3

	arr := []int{10, 20, 10, 5, 15}
	ps := realworld.PrefixSum(arr)

	want := arr[el1] + arr[el2]
	got := ps[el2] - ps[el1-1]

	if got != want {
		t.Errorf("got %+v, wanted %+v", got, want)
	}
}
