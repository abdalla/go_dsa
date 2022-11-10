package capitaluse

func DetectCapitalUse(word string) bool {
	count := 0
	size := len(word)

	//A = 65  | Z = 91
	for _, r := range word {
		if r >= 65 && r < 91 {
			count++
		}
	}

	if count == size || count == 0 {
		return true
	}

	return count == 1 && word[0] >= 65 && word[0] < 91 //only 1st char (rune) could be uper in this case
}
