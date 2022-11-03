package main

func countStudents(students []int, sandwiches []int) int {
	stu0, san0 := 0, 0
	for i := 0; i < len(students); i++ {
		if students[i] == 0 {
			stu0++
		}
		if sandwiches[i] == 0 {
			san0++
		}
	}
	num := 0
	if san0 == stu0 {
		return 0
	} else if san0 < stu0 {
		for i := 0; i < len(sandwiches); i++ {
			if sandwiches[i] == 1 {
				num++
			}
			if num == len(students)-stu0+1 {
				return len(students) - i
			}
		}
	} else {
		for i := 0; i < len(sandwiches); i++ {
			if sandwiches[i] == 0 {
				num++
			}
			if num == stu0+1 {
				return len(students) - i
			}
		}
	}
	return 0
}

//
//func main() {
//	students := []int{1, 1, 0, 0}
//	sandwiches := []int{0, 1, 0, 1}
//	fmt.Println(countStudents(students, sandwiches))
//}
