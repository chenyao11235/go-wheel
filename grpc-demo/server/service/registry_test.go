package service

import (
    "fmt"
    uuid "github.com/satori/go.uuid"
    "testing"
    "time"
)

func TestInitRegister(t *testing.T) {

    regKey := fmt.Sprintf("/service/%s/%s", "book", uuid.NewV4().String())
    regValue := fmt.Sprintf("%s:%d", "127.0.0.1", 50051)
    if err := InitRegister(regKey, regValue); err != nil {
        t.Error(err)
    }

    for range time.Tick(time.Second) {

    }
}
