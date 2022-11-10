// Write a function that reverses a string (the input strings elements are given s array of characters )
// . DO NOT allocate extra space for another array, you must do this by modifying the input array in place with O(1) extra memory
// . All the elements are printable ascii characters (symbols, numbers and alphabets)
package reversestring_test

import (
	"reflect"
	"testing"

	"github.com/abdalla/go_dsa/challenges/reversestring"
)

func TestReverseStringSuccess(t *testing.T) {
	word := []rune{'A', 'b', 'd', 'a', 'l', 'l', 'a'}

	got := reversestring.Reverse(word)
	want := []rune{'a', 'l', 'l', 'a', 'd', 'b', 'A'}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v got: %v", want, got)
	}
}

func TestReverseStringFailure(t *testing.T) {
	word := []rune{'A', 'b', 'd', 'a', 'l', 'l', 'a'}

	got := reversestring.Reverse(word)
	want := []rune{'A', 'b', 'd', 'a', 'l', 'l', 'a'}

	if reflect.DeepEqual(want, got) {
		t.Errorf("want: %v got: %v", want, got)
	}
}
