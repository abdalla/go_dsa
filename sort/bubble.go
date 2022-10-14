// Runtime: O(nˆ2)
// Memory: O(n)

package sort

func BubbleSort(slice []int) []int {
	size := len(slice)

	for i := 0; i < size-1; i++ {
		for j := 0; j < size-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	return slice
}
