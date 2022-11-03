package main

func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	var index []int
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			index = append(index, i)
		}
	}
	if len(index) != 2 {
		return false
	} else {
		if s1[index[0]] == s2[index[1]] && s1[index[1]] == s2[index[0]] {
			return true
		}
	}
	return false
}

//func main() {
//	fmt.Println(areAlmostEqual("bank", "kanb"))
//}
