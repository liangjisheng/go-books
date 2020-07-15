package main

import (
	"encoding/hex"
	"log"
)

func main() {
	// rsa 密钥文件
	bits := 1024
	prvKey, pubKey := GenRsaKey(bits)
	// log.Println(string(prvKey))
	// log.Println(string(pubKey))

	var data = "123456"
	log.Println("对消息进行签名操作...")
	signData := RsaSignWithSha256([]byte(data), prvKey)
	log.Println("消息的签名信息:", hex.EncodeToString(signData))
	log.Println("\n对签名信息进行验证")
	if RsaVerySignWithSha256([]byte(data), signData, pubKey) {
		log.Println("签名信息验证成功，确定是正确私钥签名！！")
	}

	ciphertext := RsaEncrypt([]byte(data), pubKey)
	log.Println("公钥加密后的数据:", hex.EncodeToString(ciphertext))
	sourceData := RsaDecrypt(ciphertext, prvKey)
	log.Println("私钥解密后的数据:", string(sourceData))
}
