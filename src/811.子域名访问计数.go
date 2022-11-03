package main

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	domain := map[string]int{}
	for _, s := range cpdomains {
		i := strings.IndexByte(s, ' ')
		c, _ := strconv.Atoi(s[:i])
		s = s[i+1:]
		domain[s] += c
		for {
			i := strings.IndexByte(s, '.')
			if i < 0 {
				break
			}
			s = s[i+1:]
			domain[s] += c
		}
	}
	ret := make([]string, 0, len(domain))
	for s, c := range domain {
		ret = append(ret, strconv.Itoa(c)+" "+s)
	}
	return ret
}

//func main() {
//	cpdomains := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
//	fmt.Println(subdomainVisits(cpdomains))
//}
