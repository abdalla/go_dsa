// Given an array a that contains only numbers in the range from 1 to a.length, find the first duplicate number for which the second occurrence has the minimal index. In other words, if there are more than 1 duplicated numbers, return the number for which the second occurrence has a smaller index than the second occurrence of the other number does. If there are no such elements, return -1.

// Example

// For a = [2, 1, 3, 5, 3, 2], the output should be solution(a) = 3.

// There are 2 duplicates: numbers 2 and 3. The second occurrence of 3 has a smaller index than the second occurrence of 2 does, so the answer is 3.

// For a = [2, 2], the output should be solution(a) = 2;

// For a = [2, 4, 3, 5, 1], the output should be solution(a) = -1.

// Input/Output

// [execution time limit] 4 seconds (go)

// [input] array.integer a

// Guaranteed constraints:
// 1 ≤ a.length ≤ 105,
// 1 ≤ a[i] ≤ a.length.

// [output] integer

// The element in a that occurs in the array more than once and has the minimal index for its second occurrence. If there are no such elements, return -1.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"runtime"
	"time"
)

type Data struct {
	Items []int `json:"items"`
}

func solution(a []int) int {
	checked := make(map[int]bool)

	for _, n := range a {
		if _, ok := checked[n]; ok {
			return n
		} else {
			checked[n] = true
		}
	}

	return -1
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	file, _ := ioutil.ReadFile("./challenges/arrays/01/data.json")

	data := Data{}

	_ = json.Unmarshal([]byte(file), &data)

	// data = Data{
	// 	Items: []int{2, 1, 3, 5, 3, 2},
	// }

	// data = Data{
	// 	Items: []int{28, 19, 13, 6, 34, 48, 50, 3, 47, 18, 15, 34, 16, 11, 13, 3, 2, 4, 46, 6, 7, 3, 18, 9, 32, 21, 3, 21, 50, 10, 45, 13, 22, 1, 27, 18, 3, 27, 30, 44, 12, 30, 40, 1, 1, 31, 6, 18, 33, 5},
	// }

	now := time.Now()

	fmt.Println(solution(data.Items))

	diff := time.Since(now)

	fmt.Println(diff)
}
