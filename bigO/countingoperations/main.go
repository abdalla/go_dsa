package countingoperations

import "fmt"

func Counter(list []string) int {
	first := list[0]                     // O(1)
	total := 0                           // O(1)
	newList := make([]string, len(list)) //O(1)

	//O(n)
	for _, current := range list {
		total += 1                         //O(n)
		newList = append(newList, current) //O(n)
	}

	fmt.Println(newList) //O(1)
	fmt.Println(first)   //O(1)

	return total //O(1)
}
