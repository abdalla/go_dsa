package bigo

func BigOn(studentList []string) error {
	// LINEAR
	// O(n) == worst case - 5 times
	for _, student := range studentList {
		if student == "student3" {
			println("ok")
			break
		}
	}

	return nil
}
