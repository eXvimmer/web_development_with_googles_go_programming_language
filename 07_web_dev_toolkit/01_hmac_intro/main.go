package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	c := getCode("test@example.com")
	fmt.Println(c)
	c1 := getCode("test@example.com")
	fmt.Println(c1)
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("ourSecretKey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
