package main

import (
	"fmt"
	"math/rand"
	"time"

	bigo "github.com/abdalla/go_dsa/bigO"
	"github.com/abdalla/go_dsa/sort"
)

var (
	studentlist1 = []string{"student1", "student2", "student3"}
	studentlist2 = []string{"student1", "student2", "student3", "student4", "student5"}
)

func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(99999) - rand.Intn(99999)
	}
	return slice
}

func main() {
	// bigo.BigOn(studentlist1) //O(N)
	// bigo.BigO1(studentlist2) //O(2)

	// countingoperations.Counter(studentlist2) //O(6 + 2n) ==> O(n) ==> Rule 3: removed unecessary constant

	// total := bigo.NM(studentlist1, studentlist2) //O(n + m)

	total := bigo.Exponential(studentlist2) //O(n^ 2)

	fmt.Println(total) //O(1)

	// slice := generateSlice(50)
	slice := []int{4, 1, 5, 8, 9, 3, 10, 2, 6, 7, 4, 4, 3, 2, 1, 9, 10}
	fmt.Println("======= MERGE SORT =========")
	fmt.Println("\n ---> unsorted <--- \n\n", slice)
	fmt.Println("\n --->  sorted  <---\n\n", sort.MergeSort(slice))

	fmt.Println("\n\n ======= INSERTION SORT =========")
	fmt.Println("\n ---> unsorted <--- \n\n", slice)
	fmt.Println("\n --->  sorted  <---\n\n", sort.InsertionSort(slice))
}
