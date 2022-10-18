// O(log n)
package search

func BinarySearch(slice []int, item int) bool {
	if len(slice) <= 0 {
		return false
	}

	mid := len(slice) / 2

	selected := slice[mid]

	if selected == item {
		return true
	}

	if len(slice) < 2 {
		return false
	}

	if item < selected {
		return BinarySearch(slice[:mid], item)
	}

	return BinarySearch(slice[mid:], item)
}
