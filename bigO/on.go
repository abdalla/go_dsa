package main

var (
	studentlist1 = []string{"student1", "student2", "student3"}
	studentlist2 = []string{"student1", "student2", "student3", "student4", "student5"}
)

func checkStudent(studentList []string) {

	// LINEAR
	// O(n) == worst case - 5 times
	for _, student := range studentList {
		if student == "student3" {
			println("ok")
			break
		}
	}
}

func main() {
	checkStudent(studentlist1)
}
