package goio

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

//io.Reader 接口的用法  实现io.Reader表示该类是可读的，比如os.File, bytes.Buffer
func TestReader(t *testing.T) {
	reader := strings.NewReader("from string")

	content := make([]byte, 20)
	n, err := reader.Read(content)
	if err != nil {
		t.Error(err)
	}
	t.Log(n, len(content))

	for _, val := range content {
		t.Log(val)
	}
	// 第一次读取已经读完了，所以会EOF
	n, err = reader.Read(content)
	if err != nil {
		t.Error(err)
	}
	t.Log(n)
}

//io.ReaderAt 接口的用法
func TestReaderAt(t *testing.T) {
	reader := strings.NewReader("from string")

	content := make([]byte, 5)
	// 这里读完了 会有EOF异常，这是ReadAt和Read不同的地方，ReadAt比Read严格
	n, err := reader.ReadAt(content, 2)
	if err != nil {
		t.Error(err)
	}
	t.Log(n, len(content))
	t.Log(string(content))

	n, err = reader.ReadAt(content, 2)
	if err != nil {
		t.Error(err)
	}
	t.Log(n, len(content))
	t.Log(string(content))
}

//WriterAt 接口的用法
func TestWriterAt(t *testing.T) {
	file, err := os.Create("writeAt.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	file.WriteString("Golang中文社区——这里是多余")
	n, err := file.WriteAt([]byte("Go语言中文网"), 24)
	if err != nil {
		t.Error(err)
	}
	t.Log(n)
}

//io.ReaderFrom 接口的用法
func TestReaderFrom(t *testing.T) {
	file, err := os.Open("writeAt.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	writer := bytes.NewBuffer([]byte{})

	// file 实现了Reader接口
	n, err := writer.ReadFrom(file)
	if err != nil {
		t.Error(err)
	}
	t.Log(n)
	t.Log(writer.String())
}

func TestSeek(t *testing.T) {
	// golang默认是utf8编码，一个中文字符占三个字节
	reader := strings.NewReader("我是你爸爸")
	// 从末尾，往前偏移6个字节开始读，golang
	reader.Seek(-6, io.SeekEnd)
	//ReadRune 用于读取单个字节 rune等价uint32 可以用来表示任意的单个字符，当然包括中文
	r, _, _ := reader.ReadRune()
	t.Logf("%c\n", r)

	// 从开头往后数三个字节开始读，
	reader.Seek(3, io.SeekStart)
	r, _, _ = reader.ReadRune()
	t.Logf("%c\n", r)

	reader.Seek(0, io.SeekCurrent)
	r, _, _ = reader.ReadRune()
	t.Logf("%c\n", r)
}

func TestIOCopy(t *testing.T) {
	io.Copy(os.Stdout, strings.NewReader("hello."))
}
