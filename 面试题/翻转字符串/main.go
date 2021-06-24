package main

import "fmt"

func reverseStr(s string) (string, bool) {
	str := []rune(s)

	l := len(str)

	if l > 5000 {
		return s, false
	}
	for i := 0; i < l/2; i++ {
		str[i], str[l-i-1] = str[l-i-1], str[i]
	}

	return string(str), true
}

func main() {
	fmt.Println(reverseStr("abcdefg"))
}
