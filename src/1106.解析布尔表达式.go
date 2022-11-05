package main

func parseBoolExpr(expression string) bool {
	stk := []byte{}
	for i := 0; i < len(expression); i++ {
		if expression[i] != ',' && expression[i] != ')' {
			stk = append(stk, expression[i])
		}
		tf := []byte{}
		var expr byte
		var res byte
		if expression[i] == ')' {
			for stk[len(stk)-1] != '(' {
				tf = append(tf, stk[len(stk)-1])
				stk = stk[:len(stk)-1]
			}
			//fmt.Printf("%c\n", tf)
			//fmt.Printf("%c\n", stk)
			stk = stk[:len(stk)-1]
			expr = stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			switch expr {
			case '&':
				res = 't'
				for _, b2 := range tf {
					if b2 == 'f' {
						res = 'f'
					}
				}
			case '|':
				res = 'f'
				for _, b2 := range tf {
					if b2 == 't' {
						res = 't'
					}
				}
			case '!':
				res = 't'
				for _, b2 := range tf {
					if b2 == 't' {
						res = 'f'
					}
				}
			}
			stk = append(stk, res)
		}
	}
	//fmt.Printf("%c\n", stk)
	if stk[0] == 'f' {
		return false
	}
	return true
}

//func main() {
//	fmt.Println(parseBoolExpr("|(&(t,f,t),!(t))"))
//}
