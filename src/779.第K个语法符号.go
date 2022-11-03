package main

func kthGrammar(n int, k int) int {
	if k == 1 {
		return 0
	}
	if k > 1<<(n-2) {
		return 1 ^ kthGrammar(n-1, k-1<<(n-2))
	}
	return kthGrammar(n-1, k)
}

//func main() {
//	fmt.Println(kthGrammar(3, 1))
//}
