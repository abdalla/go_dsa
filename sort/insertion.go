// Runtime: O(nˆ2)
// Memory: O(n)

package sort

func InsertionSort(slice []int) []int {
	size := len(slice)
	sorted := []int{}
	sorted = append(sorted, slice...) //smalltrick

	for i := 1; i < size; i++ {
		current := sorted[i]

		j := i - 1
		for j >= 0 && sorted[j] > current {
			sorted[j+1] = sorted[j]

			j = j - 1
		}

		sorted[j+1] = current

	}

	return sorted
}
