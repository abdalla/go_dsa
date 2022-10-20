// Given a string s consisting of small English letters, find and return the first instance of a non-repeating character in it. If there is no such character, return '_'.

// Example

// For s = "abacabad", the output should be
// solution(s) = 'c'.

// There are 2 non-repeating characters in the string: 'c' and 'd'. Return c since it appears in the string first.

// For s = "abacabaabacaba", the output should be
// solution(s) = '_'.

// There are no characters in this string that do not repeat.

// Input/Output

// [execution time limit] 4 seconds (go)

// [input] string s

// A string that contains only lowercase English letters.

// Guaranteed constraints:
// 1 ≤ s.length ≤ 105.

// [output] char

// The first non-repeating character in s, or '_' if there are no characters that do not repeat.
package main

import (
	"fmt"
	"strings"
	"time"
)

type Data struct {
	Input  string
	Output string
}

func solution(s string) string {
	checked := make(map[rune]bool)
	aux := "_"

	for _, current := range s {
		if _, ok := checked[current]; ok {
			checked[current] = false
		} else {
			checked[current] = true
		}
	}

	lIndex := len(s)
	for key, value := range checked {
		if value {
			index := strings.Index(s, string(key))
			if lIndex > index {
				lIndex = index
				aux = string(key)
			}
		}
	}

	return aux
}

func main() {
	datas := []Data{
		{
			Input:  "abacabad",
			Output: "c",
		},
		{
			Input:  "abacabaabacaba",
			Output: "_",
		},
		{
			Input:  "z",
			Output: "z",
		},
		{
			Input:  "bcb",
			Output: "c",
		},
		{
			Input:  "xdnxxlvupzuwgigeqjggosgljuhliybkjpibyatofcjbfxwtalc",
			Output: "d",
		},
	}

	for _, data := range datas {
		fmt.Println("===============>")
		now := time.Now()
		fmt.Println("input: ", data.Input)
		fmt.Println("output: ", data.Output)
		fmt.Println("solution: ", solution(data.Input))
		diff := time.Since(now)

		fmt.Println("Runtime: ", diff)
		fmt.Println("<===============")
	}

}
