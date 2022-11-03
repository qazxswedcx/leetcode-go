package main

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	res := 0
	if ruleKey == "type" {
		for _, item := range items {
			if item[0] == ruleValue {
				res++
			}
		}
	} else if ruleKey == "color" {
		for _, item := range items {
			if item[1] == ruleValue {
				res++
			}
		}
	} else if ruleKey == "name" {
		for _, item := range items {
			if item[2] == ruleValue {
				res++
			}
		}
	}
	return res
}

//func main() {
//	items := [][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}
//	fmt.Println(countMatches(items, "type", "phone"))
//}
