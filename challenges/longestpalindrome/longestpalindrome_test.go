package longestpalindrome_test

import (
	"testing"

	"github.com/abdalla/go_dsa/challenges/longestpalindrome"
)

func TestLongestPalindrome1(t *testing.T) {
	word := "xyxraccart"

	want := "raccar"
	got := longestpalindrome.GetLongestPalindrome(word)

	if got != want {
		t.Errorf("got: %v want %v", got, want)
	}

}

func TestLongestPalindrome2(t *testing.T) {
	word := "Saippuakivikauppias"

	want := "aippuakivikauppia"
	got := longestpalindrome.GetLongestPalindrome(word)

	if got != want {
		t.Errorf("got: %v want %v", got, want)
	}

}

func TestLongestPalindrome3(t *testing.T) {
	word := "saippuakivikauppias"

	want := "saippuakivikauppias"
	got := longestpalindrome.GetLongestPalindrome(word)

	if got != want {
		t.Errorf("got: %v want %v", got, want)
	}

}
