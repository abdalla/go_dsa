package bigo

import "fmt"

func NM(list []string, list2 []string) int {
	total := 0 // O(1)

	//O(n)
	for _, current := range list {
		for _, nasted := range list2 {
			fmt.Println(current, nasted) //O(nˆ2)
			total += 1                   //O(nˆ2)
		}
	}

	return total //O(1)
}
