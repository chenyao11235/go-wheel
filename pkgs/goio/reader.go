package goio

import (
	"io"
)

//ReadFrom  从某个地方(reader)读取num个字节的数据
func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)

	if n > 0 {
		return p[:n], nil
	}

	return p, err
}
