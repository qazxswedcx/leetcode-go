package main

func findAnagrams(s string, p string) []int {
	pcount := [26]int{}
	for _, i := range p {
		pcount[i-'a']++
	}
	scount := [26]int{}
	res := []int{}
	for i := 0; i < len(s)-len(p)+1; i++ {
		if i == 0 {
			for _, i1 := range s[:len(p)] {
				scount[i1-'a']++
			}
		} else {
			scount[s[i-1]-'a']--
			scount[s[i+len(p)-1]-'a']++
		}
		if pcount == scount {
			res = append(res, i)
		}
	}
	return res
}

//func main() {
//	fmt.Println(findAnagrams("cbaebabacd", "abc"))
//}
