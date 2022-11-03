package main

import (
	"strconv"
)

func haveConflict(event1 []string, event2 []string) bool {
	if compareTime(event1[0], event2[0]) && compareTime(event2[0], event1[1]) {
		return true
	}
	if compareTime(event1[0], event2[1]) && compareTime(event2[1], event1[1]) {
		return true
	}
	if compareTime(event2[0], event1[0]) && compareTime(event1[1], event2[1]) {
		return true
	}
	return false
}

func compareTime(a string, b string) bool {
	hour1, _ := strconv.Atoi(a[:2])
	hour2, _ := strconv.Atoi(b[:2])
	min1, _ := strconv.Atoi(a[3:])
	min2, _ := strconv.Atoi(b[3:])
	if hour1 < hour2 {
		return true
	} else if hour1 == hour2 {
		if min1 <= min2 {
			return true
		}
	}
	return false
}

//func main() {
//	event1 := []string{"10:00", "11:00"}
//	event2 := []string{"10:00", "15:00"}
//	fmt.Println(haveConflict(event2, event1))
//}
