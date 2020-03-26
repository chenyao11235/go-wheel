package main

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"wheel/encrypt/aes"
)

// 身份验证

func Auth(w http.ResponseWriter, r *http.Request) (err error) {
	fmt.Println("进行身份验证...")

	err = checkUserInfo(r)

	return

}

func checkUserInfo(r *http.Request) (err error) {
	var (
		uidCookie  *http.Cookie
		signCookie *http.Cookie
		signByte   []byte
	)
	if uidCookie, err = r.Cookie("uid"); err != nil {
		return
	}
	// 获取用户uid的加密字符串
	if signCookie, err = r.Cookie("sign"); err != nil {
		return
	}

	// 对加密字符串进行解密
	if signByte, err = aes.DePwdCode(signCookie.Value); err != nil {
		return
	}
	// 进行比对
	if checkInfo(uidCookie.Value, string(signByte)) {
		return
	}

	err = errors.New("身份校验失败")
	return
}

// 以后可以在这里进行更加复杂的逻辑
func checkInfo(checkStr, signStr string) (bool) {
	if checkStr == signStr {
		return true
	}
	return false

}
