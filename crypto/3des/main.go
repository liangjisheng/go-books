package main

import (
	"encoding/hex"
	"log"

	"github.com/forgoer/openssl"
)

// DES 和 3DES 加密区别
// 前者 加密 秘钥必须是 8bytes
// 后者加密 解密 再加密 密钥必须是 24bytes
func main() {
	// 定义密钥, 必须是 24bytes
	key := []byte("123456789012345678901234")
	log.Println("key hex:    ", hex.EncodeToString(key))

	// 明文
	origData := []byte("6226 0000 1234 5678")
	origData = append(origData, []byte{128, 0, 0, 0, 0}...)
	// 363232362030303030203132333420353637388000000000
	log.Println("origin hex: ", hex.EncodeToString(origData))

	// 加密, 传入 "", 表示不填充, 还有另外3种填充模式
	// const PKCS5_PADDING = "PKCS5"
	// const PKCS7_PADDING = "PKCS7"
	// const ZEROS_PADDING = "ZEROS"
	dst, err := openssl.Des3ECBEncrypt(origData, key, "")
	if err != nil {
		log.Panic("err:", err)
	}
	log.Println("secret hex: ", hex.EncodeToString(dst))

	// 解密
	dst, err = openssl.Des3ECBDecrypt(dst, key, "")
	if err != nil {
		log.Panic("err:", err)
	}
	log.Println("decrypt hex:", hex.EncodeToString(dst))
}
