// Runtime: O(nˆ2)
// Memory: O(n)

package sort

func SelectionSort(slice []int) []int {
	size := len(slice)
	sorted := []int{}
	sorted = append(sorted, slice...)

	for i := 0; i < size; i++ {
		minIndex := i
		for j := i; j < size; j++ {
			if sorted[j] < sorted[minIndex] {
				minIndex = j
			}
		}

		if sorted[i] > sorted[minIndex] {
			sorted[i], sorted[minIndex] = sorted[minIndex], sorted[i]
		}

	}

	return sorted
}
