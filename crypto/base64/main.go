package main

import (
	"encoding/base64"
	"log"
)

// Base64Encode ...
func Base64Encode(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

// Base64Decode ...
func Base64Decode(src string) (string, error) {
	dst, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return "", err
	}
	return string(dst), nil
}

func main() {
	src := "hello"
	b64 := Base64Encode(src)
	log.Println("src:   ", src)
	log.Println("encode:", b64)

	b64Decode, _ := Base64Decode(b64)
	log.Println("decode:", b64Decode)
}
