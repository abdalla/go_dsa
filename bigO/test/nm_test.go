package test

import (
	"testing"

	bigo "github.com/abdalla/go_dsa/bigO"
)

func TestBigONM(t *testing.T) {

	got := bigo.NM(studentlist1, studentlist2)
	want := 15 // # iteractions

	println(got)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
