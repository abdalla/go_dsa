package test

import (
	"testing"

	bigo "github.com/abdalla/go_dsa/bigO"
)

func TestBigOExponential(t *testing.T) {

	got := bigo.Exponential(studentlist2)
	want := 25 // # iteractions

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
