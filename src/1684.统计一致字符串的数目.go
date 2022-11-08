package main

func countConsistentStrings(allowed string, words []string) int {
	exist := [26]bool{}
	for _, i := range allowed {
		exist[i-'a'] = true
	}
	judge := func(word string) bool {
		for _, i := range word {
			if exist[i-'a'] == false {
				return false
			}
		}
		return true
	}
	res := 0
	for _, i := range words {
		if judge(i) == true {
			res++
		}
	}
	return res
}
