package main

import (
    "fmt"
    "wheel/utils/aes"
)

func main() {
    s := "hello world"

    en_s, _ := aes.EnPwdCode([]byte(s))
    fmt.Println(en_s)

    ori_s, _ := aes.DePwdCode(en_s)
    fmt.Println(string(ori_s))
}