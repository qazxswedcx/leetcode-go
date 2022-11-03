package main

func generateParenthesis(n int) []string {
	res := []string{}
	if n <= 0 {
		return res
	}
	getParenthesis(&res, "", n, n)
	return res
}

func getParenthesis(res *[]string, str string, left int, right int) {
	if left == 0 && right == 0 {
		*res = append(*res, str)
		return
	}
	if left == right {
		getParenthesis(res, str+"(", left-1, right)
	} else if left < right {
		if left > 0 {
			getParenthesis(res, str+"(", left-1, right)
		}
		getParenthesis(res, str+")", left, right-1)
	}

}

//
//func main() {
//	fmt.Println(generateParenthesis(8))
//	fmt.Println(len(generateParenthesis(8)))
//}
