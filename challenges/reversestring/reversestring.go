package reversestring

func Reverse(word []rune) []rune {
	pointerI := 0
	pointerJ := len(word) - 1

	for pointerI < pointerJ {
		word[pointerI], word[pointerJ] = word[pointerJ], word[pointerI]

		pointerI++
		pointerJ--
	}

	return word
}
