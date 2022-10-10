package bigo

import "fmt"

func NM(list []string, list2 []string) int {
	total := 0 // O(1)

	//O(n)
	for _, current := range list {
		//O(m)
		for _, nasted := range list2 {
			fmt.Println(current, nasted) //O(n + m)
			total += 1                   //O(n + m)
		}
	}

	return total //O(1)
}
