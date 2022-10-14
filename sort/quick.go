// Quicksort is an efficient sorting algorithm commonly used in production sorting implementations. Like Merge Sort , Quicksort is a divide-and-conquer algorithm .
// As the name implies, Quicksort is one of the fastest sorting algorithms, but you have to pay attention to detail in your implementation because
// if you’re not careful, your speed can drop quickly.

// Divide
// Select a pivot element that will preferably end up close to the center of the sorted pack
// Move everything onto the “greater than” or “less than” side of the pivot
// The pivot is now in its final position
// Recursively repeat the operation on both sides of the pivot

// Runtime: O(n^2)
// Memory: O(log n)

package sort

func partition(slice []int, low, high int) ([]int, int) {
	pivot := slice[high]
	i := low

	for j := i; j < high; j++ {
		if slice[j] <= pivot {
			slice[i], slice[j] = slice[j], slice[i]
			i++
		}
	}

	slice[high], slice[i] = slice[i], slice[high]

	return slice, i
}

func QuickSort(slice []int, low, high int) []int {
	var sorted []int
	sorted = append(sorted, slice...)

	if low < high {
		var p int
		sorted, p = partition(sorted, low, high)
		sorted = QuickSort(sorted, low, p-1)
		sorted = QuickSort(sorted, p+1, high)
	}

	return sorted
}
