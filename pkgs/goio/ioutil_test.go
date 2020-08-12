package goio

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

type MyReader interface {
	Read(p []byte) (n int, err error)
}

// ioutil.NopCloser 用来将io.Reader 接口转换成io.ReadCloser 接口
func TestNopCloser(t *testing.T) {
	var body MyReader
	if rc, ok := body.(io.ReadCloser); !ok {
		t.Log("没有实现ReadCloser接口")
		rc = ioutil.NopCloser(body)
		t.Log("用NopCloser函数转换一下就行了")
		rc.Close()
		t.Log("现在好了，可以调用Close方法了")
	}
}

//ReadAll 用于将io.Reader中的内容一次性全部读取出来
func TestReadAll(t *testing.T) {
	fileReader, err := os.Open("./writeAt.txt")
	if err != nil {
		t.Error(err)
	}
	content, err := ioutil.ReadAll(fileReader)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(content))
}

//ReadDir 用来读取某个目录下的所有文件
func TestReadDir(t *testing.T) {
	path := "/Users/yaochen/go/wheel"
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		t.Error(err)
	}

	for _, info := range fileInfos {
		if info.IsDir() {
			t.Logf("%s is dir.", info.Name())
			//这里也可以进行递归的遍历，打印子目录中的文件
		} else {
			t.Logf("%s is file.", info.Name())
		}
	}
}

//ReadFile 用于读取整个文件的所有内容
func TestReadFile(t *testing.T) {
	content, err := ioutil.ReadFile("./writeAt.txt")
	if err != nil {
		t.Error(err)
	}
	t.Log(string(content))
}

func TestWriteFile(t *testing.T) {
	content := []byte("跟我一起学习go语言的标准库库吧.")
	err := ioutil.WriteFile("./writeFile.txt", content, 0666)
	if err != nil {
		t.Error(err)
	}
}

//TempDir 创建临时的目录
func TestTempDir(t *testing.T) {
	//第一个参数如果为空，表明在系统默认的临时目录（ os.TempDir ）中创建临时目录；第二个参数指定临时目录名的前缀，该函数返回临时目录的路径。
	path, err := ioutil.TempDir("", "go-pkg-test")
	if err != nil {
		t.Log(err)
	}
	t.Log(path)
}

func TestTempFile(t *testing.T) {
	f, err := ioutil.TempFile("", "gofmt")
	if err != nil {
		t.Error(err)
	}
	t.Log(f.Name())
}
