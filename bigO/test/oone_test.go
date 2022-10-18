package test

import (
	"testing"

	bigo "github.com/abdalla/go_dsa/bigO"
)

func TestBigO1(t *testing.T) {

	err := bigo.BigO1(studentlist2)

	if err != nil {
		t.Errorf(err.Error())
	}
}
