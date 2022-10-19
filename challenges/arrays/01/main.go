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

func setDuplicateIndex(a []int, seen []int, index int, c chan bool) {
	// var wg sync.WaitGroup
	size := len(a)

	// exists := make(chan bool)
	exists := false

	if size == 2 {
		if a[0] == a[1] {
			seen[1] = 1
			exists = true
		}

		// return seen, exists
	}

	for i := range a {
		if seen[index+i] == 1 {
			exists = true
			break
			// return seen, exists
		}

		j := index + i + 1
		// wg.Add(1)
		//channel := make(chan bool)

		//go func(item int, cx chan bool) {
		for ; j < size; j++ {
			if seen[j] == 1 {
				exists = true
				break
			}

			if a[i] == a[j] {
				seen[j] = 1
				exists = true
				break
			}
		}

		//cx <- exists
		//}(a[i], channel)

		//exists = <-channel
	}

	// wg.Wait()

	c <- exists

	// return seen, exists
}

//func solution3(a []int) int {
//
//}

func solution2(a []int) int {
	// runtime.Gosched()

	size := len(a)
	mid := size / 2
	seen := make([]int, size, size)
	exists := make(chan bool)
	exists2 := make(chan bool)

	if size == 2 {
		if a[0] == a[1] {
			return a[0]
		}

		return -1
	}

	// var wg sync.WaitGroup

	// wg.Add(2)

	// go func() {
	// defer wg.Done()
	go setDuplicateIndex(a[:mid], seen, 0, exists)
	// }()

	// go func() {
	// defer wg.Done()
	go setDuplicateIndex(a[mid:], seen, mid, exists2)
	// }()

	// wg.Wait()

	has1, has2 := <-exists, <-exists2

	if !has1 {
		x := make(chan bool)
		go setDuplicateIndex(a, seen, 0, x)

		has1 = <-x
	}

	if has1 || has2 {
		for i := range seen {
			if seen[i] == 1 {
				return a[i]
			}
		}
	}

	return -1
}

//func solution(a []int) int {
//
//	// runtime.GOMAXPROCS(runtime.NumCPU())
//	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
//
//	size := len(a)
//	seen := make([]int, size, size)
//
//	if size == 2 {
//		if a[0] == a[1] {
//			return a[0]
//		}
//
//		return -1
//	}
//
//	for i := range a {
//		if seen[i] == 1 {
//			return a[i]
//		}
//
//		j := i + 1
//		go func(item int) {
//			for ; j < size; j++ {
//				if seen[j] == 1 {
//					break
//				}
//
//				if item == a[j] {
//					seen[j] = 1
//					break
//				}
//			}
//		}(a[i])
//
//		if size < 5000 {
//			time.Sleep(500 * time.Nanosecond)
//		}
//	}
//
//	return -1
//}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	file, _ := ioutil.ReadFile("./challenges/arrays/01/data.json")

	data := Data{}

	_ = json.Unmarshal([]byte(file), &data)

	//data = Data{
	//	Items: []int{2, 1, 3, 5, 3, 2},
	//}

	//data = Data{
	//	Items: []int{28, 19, 13, 6, 34, 48, 50, 3, 47, 18, 15, 34, 16, 11, 13, 3, 2, 4, 46, 6, 7, 3, 18, 9, 32, 21, 3, 21, 50, 10, 45, 13, 22, 1, 27, 18, 3, 27, 30, 44, 12, 30, 40, 1, 1, 31, 6, 18, 33, 5},
	//}

	now := time.Now()

	fmt.Println(solution2(data.Items))

	diff := time.Since(now)

	fmt.Println(diff)
}
