package main

func minAddToMakeValid(s string) int {
	var count int = 0
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				stack = append(stack, pairs[s[i]])
				count++
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return count + len(stack)
}

//func main() {
//	fmt.Println(minAddToMakeValid("(())))))"))
//}
