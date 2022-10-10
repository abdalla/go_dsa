package main

import (
	"fmt"

	bigo "github.com/abdalla/go_dsa/bigO"
)

var (
	studentlist1 = []string{"student1", "student2", "student3"}
	studentlist2 = []string{"student1", "student2", "student3", "student4", "student5"}
)

func main() {
	// bigo.BigOn(studentlist1) //O(N)
	// bigo.BigO1(studentlist2) //O(2)

	// countingoperations.Counter(studentlist2) //O(6 + 2n) ==> O(n) ==> Rule 3: removed unecessary constant

	total := bigo.NM(studentlist1, studentlist2) //O(n + m)

	// total := bigo.Exponential(studentlist2) //O(n^ 2)

	fmt.Println(total) //O(1)
}
