package main

import (
	"fmt"
	"strings"
)

func isReGroup(s1, s2 string) bool {

	sl1 := len([]rune(s1))
	sl2 := len([]rune(s2))

	if sl1 > 5000 || sl2 > 5000 || sl1 != sl2 {
		return false
	}

	for _, item := range s1 {
		if strings.Count(s1, string(item)) != strings.Count(s2, string(item)) {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isReGroup("abc", "acb"))
}
