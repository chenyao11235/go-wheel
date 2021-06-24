package main

import "fmt"

type Student struct {
	name string
}

func main() {
	m := map[string]*Student{"people": &Student{"eric"}}
	fmt.Println(m["people"].name)
}
