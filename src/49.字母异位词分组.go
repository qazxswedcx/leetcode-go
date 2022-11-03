package main

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	mapStr := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		mapStr[sortedStr] = append(mapStr[sortedStr], str)
	}
	ans := [][]string{}
	for _, strings := range mapStr {
		ans = append(ans, strings)
	}
	return ans
}

//
//func main() {
//	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
//	fmt.Println(groupAnagrams(strs))
//}
