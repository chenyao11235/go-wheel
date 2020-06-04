package utils

import (
	"crypto/rand"
	"fmt"
    uuid "github.com/satori/go.uuid"
    "log"
	"testing"
)

// 生成一个随机字符串
// create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid

func TestUUID(t *testing.T) {
    u := new([16]byte)
    _, err := rand.Read(u[:])
    if err != nil {
        log.Fatalln("Cannot generate UUID", err)
    }

    // 0x40 is reserved variant from RFC 4122
    u[8] = (u[8] | 0x40) & 0x7F
    // Set the four most significant bits (bits 12 through 15) of the
    // time_hi_and_version field to the 4-bit version number.
    u[6] = (u[6] & 0xF) | (0x4 << 4)
    uuid := fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
    t.Log(uuid)
}

func TestUUID2(t *testing.T){
    u1 := uuid.NewV4()
    fmt.Printf("UUIDv4: %s\n", u1)

    // 解析
    u2, err := uuid.FromString("f5394eef-e576-4709-9e4b-a7c231bd34a4")
    if err != nil {
        fmt.Printf("Something gone wrong: %s", err)
        return
    }
    fmt.Printf("Successfully parsed: %s", u2)
}