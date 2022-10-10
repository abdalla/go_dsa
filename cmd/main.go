package main

import (
	bigo "github.com/abdalla/go_dsa/bigO"
	"github.com/abdalla/go_dsa/bigO/countingoperations"
)

var (
	studentlist1 = []string{"student1", "student2", "student3"}
	studentlist2 = []string{"student1", "student2", "student3", "student4", "student5"}
)

func main() {
	bigo.BigOn(studentlist1) //O(N)
	bigo.BigO1(studentlist2) //O(2)

	countingoperations.Counter(studentlist2) //O(6 + 2n) ==> O(n) ==> Rule 3: removed unecessary constant
}
