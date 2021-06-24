package main

import (
	"encoding/json"
	"fmt"
)

//按照 golang 的语法，小写开头的方法、属性或 struct 是私有的，同样，在 json 解 码或转码的时候也无法上线私有属性的转换。
//题目中是无法正常得到 People 的 name 值的。而且，私有属性 name 也不应该加
//json 的标签。

type People struct {
	name string `json:"name"`
}

func main() {
	js := `{ "name":"11"
	}`

	var p People

	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	fmt.Println("people: ", p)
}
