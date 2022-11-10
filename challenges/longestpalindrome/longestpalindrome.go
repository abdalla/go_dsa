package longestpalindrome

func checkPalindrome(word string, left, right int) string {
	for left >= 0 && right < len(word) && word[left] == word[right] {
		left--
		right++
	}

	return word[left+1 : right]
}

func GetLongestPalindrome(word string) string {
	result := ""

	for i := range word {
		word1 := checkPalindrome(word, i, i)
		word2 := checkPalindrome(word, i, i+1)

		longest := ""
		if len(word1) >= len(word2) {
			longest = word1
		} else {
			longest = word2
		}

		if len(longest) >= len(result) {
			result = longest
		}
	}

	return result
}
