package goaes

import (
	"fmt"
	"testing"
)

func TestAES(t *testing.T) {
	s := "hello world"
	ens, _ := EnPwdCode([]byte(s))
	fmt.Println(ens)
	oris, _ := DePwdCode(ens)
	t.Log(string(oris))
}
