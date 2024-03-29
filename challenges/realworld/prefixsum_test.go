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

// func TestSum2Elements1and3(t *testing.T) {
// 	el1 := 1
// 	el2 := 3

// 	idx1 := el1 - 1
// 	idx2 := el2 - 1

// 	arr := []int{10, 20, 10, 5, 15}
// 	ps := realworld.PrefixSum(arr) //=> {10, 30, 40, 45, 60}

// 	want := arr[idx1] + arr[idx2] //==> 10 + 10 = 20
// 	got := ps[idx2] - ps[idx2-1] //=> 40 - 20 = 20

// 	if got != want {
// 		t.Errorf("got %+v, wanted %+v", got, want)
// 	}
// }

// func TestSum2Elements1and2(t *testing.T) {
// 	el1 := 1
// 	el2 := 2

// 	idx1 := el1 - 1
// 	idx2 := el2 - 1

// 	arr := []int{10, 20, 10, 5, 15}
// 	ps := realworld.PrefixSum(arr)

// 	want := arr[idx1] + arr[idx2]
// 	got := ps[idx2] - ps[el1-1]

// 	if got != want {
// 		t.Errorf("got %+v, wanted %+v", got, want)
// 	}
// }
