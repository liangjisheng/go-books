package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/forgoer/openssl"
)

// AES 密钥的长度可以是16/24/32个字符(128/192/256位)

func main() {
	src := []byte("123456")
	log.Println("src hex:", hex.EncodeToString(src))

	// AES-128-ECB, PKCS7_PADDING
	key := []byte("1234512345123451")
	log.Println("aes-128 key hex:", hex.EncodeToString(key))

	dst, err := openssl.AesECBEncrypt(src, key, openssl.PKCS7_PADDING)
	if err != nil {
		log.Panic(err)
	}
	log.Println("base64:", base64.StdEncoding.EncodeToString(dst))
	log.Println("aes-128 secret hex:", hex.EncodeToString(dst))

	dst, err = openssl.AesECBDecrypt(dst, key, openssl.PKCS7_PADDING)
	if err != nil {
		log.Panic(err)
	}
	log.Println("aes-128 decrypt hex:", hex.EncodeToString(dst))
	fmt.Println()

	// AES-192-ECB, PKCS7_PADDING
	key = []byte("123451234512345123451234")
	log.Println("aes-192 key hex:", hex.EncodeToString(key))

	dst, err = openssl.AesECBEncrypt(src, key, openssl.PKCS7_PADDING)
	if err != nil {
		log.Panic(err)
	}
	log.Println("base64:", base64.StdEncoding.EncodeToString(dst))
	log.Println("aes-192 secret hex:", hex.EncodeToString(dst))

	dst, err = openssl.AesECBDecrypt(dst, key, openssl.PKCS7_PADDING)
	if err != nil {
		log.Panic(err)
	}
	log.Println("aes-192 decrypt hex:", hex.EncodeToString(dst))
	fmt.Println()

	// AES-256-ECB, PKCS7_PADDING
	key = []byte("12345123451234512345123451234512")
	log.Println("aes-256 key hex:", hex.EncodeToString(key))

	dst, err = openssl.AesECBEncrypt(src, key, openssl.PKCS7_PADDING)
	if err != nil {
		log.Panic(err)
	}
	log.Println("base64:", base64.StdEncoding.EncodeToString(dst))
	log.Println("aes-256 secret hex:", hex.EncodeToString(dst))

	dst, err = openssl.AesECBDecrypt(dst, key, openssl.PKCS7_PADDING)
	if err != nil {
		log.Panic(err)
	}
	log.Println("aes-256 decrypt hex:", hex.EncodeToString(dst))
}
