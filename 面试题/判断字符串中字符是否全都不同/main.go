package main

import (
	"fmt"
	"strings"
)

func isUniqStr(s string) bool {
	if len(s) > 3000 {
		return false
	}
	for _, item := range s {
		if item > 127 {
			return false
		}
		if strings.Count(s, string(item)) > 1 {
			return false
		}
	}
	return true
}

func isUniqStr2(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}

	for k, v := range s {
		if v > 127 {
			return false
		}
		if strings.Index(s, string(v)) != k {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isUniqStr("abc"))
	fmt.Println(isUniqStr2("aabc"))
}
