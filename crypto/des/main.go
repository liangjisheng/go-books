package main

import (
	"encoding/hex"
	"log"

	"github.com/forgoer/openssl"
)

func main() {
	key := []byte("12345123")
	log.Println("key hex:    ", hex.EncodeToString(key))

	src := []byte("123456")
	log.Println("src hex:    ", hex.EncodeToString(src))

	// DES-ECB, PKCS7_PADDING
	dst, err := openssl.DesECBEncrypt(src, key, openssl.PKCS7_PADDING)
	if err != nil {
		log.Panic("err:", err)
	}
	log.Println("secret hex: ", hex.EncodeToString(dst))

	dst, err = openssl.DesECBDecrypt(dst, key, openssl.PKCS7_PADDING)
	if err != nil {
		log.Panic("err:", err)
	}
	log.Println("decrypt hex:", hex.EncodeToString(dst))
}
