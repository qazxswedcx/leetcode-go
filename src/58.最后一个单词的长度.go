package main

func lengthOfLastWord(s string) int {
	right := len(s) - 1
	left := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			right = i
			break
		}
	}
	for i := right; i >= 0; i-- {
		if s[i] == ' ' {
			left = i + 1
			break
		}
	}
	return right - left + 1
}

//func main() {
//	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
//}
