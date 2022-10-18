package test

import (
	"testing"

	bigo "github.com/abdalla/go_dsa/bigO"
)

func TestBigON(t *testing.T) {

	err := bigo.BigOn(studentlist1)

	if err != nil {
		t.Errorf(err.Error())
	}
}
