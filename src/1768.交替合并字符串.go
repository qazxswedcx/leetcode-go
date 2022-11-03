package main

func mergeAlternately(word1 string, word2 string) string {
	res := ""
	for {
		if len(word1) == 0 {
			res = res + word2
			break
		}
		if len(word2) == 0 {
			res = res + word1
			break
		}

		res = res + word1[:1] + word2[:1]
		word1 = word1[1:]
		word2 = word2[1:]

	}
	return res
}

//func main() {
//	fmt.Println(mergeAlternately("abcd", "pq"))
//}
