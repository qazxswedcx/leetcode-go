package main

func ambiguousCoordinates(s string) []string {
	res := []string{}
	for i := 2; i < len(s)-1; i++ {
		lt := getPos(s[1:i])
		if len(lt) == 0 {
			continue
		}
		rt := getPos(s[i : len(s)-1])
		if len(rt) == 0 {
			continue
		}
		for _, i := range lt {
			for _, j := range rt {
				res = append(res, "("+i+", "+j+")")
			}
		}
	}
	return res
}

func getPos(s string) (pos []string) {
	if s[0] != '0' || s == "0" {
		pos = append(pos, s)
	}
	for p := 1; p < len(s); p++ {
		if p != 1 && s[0] == '0' || s[len(s)-1] == '0' {
			continue
		}
		pos = append(pos, s[:p]+"."+s[p:])
	}
	return
}

//func main() {
//	fmt.Println(ambiguousCoordinates("(0123)"))
//}
