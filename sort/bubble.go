// Runtime: O(nˆ2)
// Memory: O(n)

package sort

func BubbleSort(slice []int) []int {
	size := len(slice)
	sorted := []int{}
	sorted = append(sorted, slice...) //smalltrick

	for i := 0; i < size-1; i++ {
		for j := 0; j < size-1; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}

	return sorted
}
