package main

import (
    "fmt"
    "github.com/pkg/errors"
    "io/ioutil"
    "os"
)

func ReadFile(path string) ([]byte, error) {

    f, err := os.Open(path)
    if err != nil {
        return nil, errors.Wrap(err, "open file")
    }
    defer f.Close()

    buf, err := ioutil.ReadAll(f)
    if err != nil {
        return nil, errors.Wrap(err, "read file ")
    }
    return buf, nil
}

func main() {
    _, err := ReadFile("./test.text")
    if err != nil {
        fmt.Println(err)
    }
}
