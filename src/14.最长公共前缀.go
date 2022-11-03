package main

func longestCommonPrefix(strs []string) string {
	lenstrs := len(strs)
	minlen := 200
	minstr := ""
	for i := 0; i < lenstrs; i++ {
		if len(strs[i]) < minlen {
			minlen = len(strs[i])
			minstr = strs[i]
		}
	}
	sameindex := 0
	for i := 0; i < minlen; i++ {
		for j := 0; j < lenstrs; j++ {
			if strs[j][i] != minstr[i] {
				if sameindex == 0 {
					return ""
				}
				return minstr[:sameindex]
			} else {
				if j == lenstrs-1 {
					sameindex++
				}
			}
		}
	}
	return minstr
}

//func main() {
//	var str = []string{"aaa", "bbb"}
//	fmt.Println(longestCommonPrefix(str))
//}
