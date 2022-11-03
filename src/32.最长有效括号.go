package main

func longestValidParentheses(s string) int {
	var maxcount int = 0
	stack := []int{}
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				if i-stack[len(stack)-1] > maxcount {
					maxcount = i - stack[len(stack)-1]
				}
			}
		}
	}
	return maxcount
}

func isValid2(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if s[i] == ')' {
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

//func main() {
//	s := "(()"
//	fmt.Println(longestValidParentheses(s))
//	fmt.Println(isValid2(s[1:3]))
//}
