package util

import (
	"fmt"
	"golang.org/x/sys/windows"
	"math/rand"
	"strings"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenGuid() string {
	buf := make([]byte, 8)
	rand.Read(buf)

	return strings.ToLower(windows.GUID{
		Data1: rand.Uint32(),
		Data2: uint16(rand.Int()),
		Data3: uint16(rand.Int()),
		Data4: [8]byte(buf),
	}.String())
}

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandDate() time.Time {
	min := time.Date(2014, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2022, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func GenMac() string {
	buf := make([]byte, 6)
	_, err := rand.Read(buf)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}
	// Set the local bit
	buf[0] |= 2

	return string(buf)
}
