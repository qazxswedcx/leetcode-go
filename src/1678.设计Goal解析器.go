package main

import "fmt"

func interpret(command string) string {
	res := ""
	for i := 0; i < len(command); i++ {
		switch command[i] {
		case 'G':
			res += "G"
		case '(':
			if command[i+1] == ')' {
				res += "o"
				i++
			} else if command[i+1] == 'a' {
				res += "al"
				i += 3
			}
		}
	}
	return res
}

func main() {
	fmt.Println(interpret("(al)G(al)()()G"))
}
