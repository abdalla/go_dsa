// Merge Sort was invented by John von Neumann in 1945, and is an algorithm belonging to the “divide and conquer” class of algorithms.
// Meaning, the algorithm splits an input into various pieces, sorts them and then merges them back together.

// Runtime: O(nlogn)
// Memory: O(n)

package sort

// func merge(left, right []int) []int {
// 	size, topLeft, topRight, current := len(left)+len(right), 0, 0, 0
// 	slice := make([]int, size, size)

// 	for topLeft < len(left) && topRight < len(right) {
// 		if left[topLeft] < right[topRight] {
// 			slice[current] = left[topLeft]
// 			topLeft++
// 			current++
// 		} else {
// 			slice[current] = right[topRight]
// 			topRight++
// 			current++
// 		}
// 	}

// 	for ; topLeft < len(left); topLeft++ {
// 		slice[current] = left[topLeft]
// 		current++
// 	}

// 	for ; topRight < len(right); topRight++ {
// 		slice[current] = right[topRight]
// 		current++
// 	}

// 	return slice
// }

func merge(left, right []int) []int {
	size, topLeft, topRight, current := len(left)+len(right), 0, 0, 0
	slice := make([]int, size, size)

	for ; current < size; current++ {
		if topLeft >= len(left) {
			slice[current] = right[topRight]
			topRight++
		} else if topRight >= len(right) {
			slice[current] = left[topLeft]
			topLeft++
		} else if left[topLeft] < right[topRight] {
			slice[current] = left[topLeft]
			topLeft++
		} else {
			slice[current] = right[topRight]
			topRight++
		}
	}

	return slice
}

func MergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	mid := len(slice) / 2

	return merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}
