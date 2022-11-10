package capitaluse_test

import (
	"testing"

	"github.com/abdalla/go_dsa/challenges/capitaluse"
)

func TestCatialUse1(t *testing.T) {
	word := "Abdalla"
	got := capitaluse.DetectCapitalUse(word)
	want := true

	if got != want {
		t.Errorf("want: %v got: %v", want, got)
	}
}

func TestCatialUse2(t *testing.T) {
	word := "AbdalLa"
	got := capitaluse.DetectCapitalUse(word)
	want := false

	if got != want {
		t.Errorf("want: %v got: %v", want, got)
	}
}

func TestCatialUse3(t *testing.T) {
	word := "abdalLa"
	got := capitaluse.DetectCapitalUse(word)
	want := false

	if got != want {
		t.Errorf("want: %v got: %v", want, got)
	}
}

func TestCatialUse4(t *testing.T) {
	word := "abdalla"
	got := capitaluse.DetectCapitalUse(word)
	want := true

	if got != want {
		t.Errorf("want: %v got: %v", want, got)
	}
}

func TestCatialUse5(t *testing.T) {
	word := "abdalla"
	got := capitaluse.DetectCapitalUse(word)
	want := true

	if got != want {
		t.Errorf("want: %v got: %v", want, got)
	}
}

func TestCatialUse6(t *testing.T) {
	word := "abdALLa"
	got := capitaluse.DetectCapitalUse(word)
	want := false

	if got != want {
		t.Errorf("want: %v got: %v", want, got)
	}
}

func TestCatialUse7(t *testing.T) {
	word := "ABDZLLA"
	got := capitaluse.DetectCapitalUse(word)
	want := true

	if got != want {
		t.Errorf("want: %v got: %v", want, got)
	}
}

func TestCatialUse8(t *testing.T) {
	word := "ABDZlLA"
	got := capitaluse.DetectCapitalUse(word)
	want := false

	if got != want {
		t.Errorf("want: %v got: %v", want, got)
	}
}
