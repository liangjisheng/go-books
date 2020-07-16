package main

import (
	"bytes"
	"crypto/cipher"
	"log"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm3"
	"github.com/tjfoc/gmsm/sm4"
)

// HashSM3 消息摘要 可以用MD5作为对比理解。该算法已公开。校验结果为256位
func HashSM3() {
	data := "hello world"
	s := sm3.New()
	s.Write([]byte(data))
	sum := s.Sum(nil)
	log.Printf("%x\n", sum)
}

// SM4 ...
func SM4() {
	// SM4私钥只支持16字节(128位)
	key := []byte("helloworldhellow")
	iv := make([]byte, sm4.BlockSize)
	data := []byte("Tongji Fintech Research Institute")
	log.Printf("明文: %s\n", data)
	ciphertxt, err := sm4Encrypt(key, iv, data)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("加密结果: %x\n", ciphertxt)
	// log.Println(ciphertxt, len(ciphertxt))
	data, err = sm4Decrypt(key, iv, ciphertxt)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("解密结果: %s\n", data)
}

// SM4 无线局域网标准的分组数据算法。对称加密，密钥长度和分组长度均为128位
func sm4Encrypt(key, iv, plainText []byte) ([]byte, error) {
	block, err := sm4.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData := pkcs5Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)
	return cryted, nil
}

func sm4Decrypt(key, iv, cipherText []byte) ([]byte, error) {
	block, err := sm4.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(origData, cipherText)
	origData = pkcs5UnPadding(origData)
	return origData, nil
}

// pkcs5填充
func pkcs5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func pkcs5UnPadding(src []byte) []byte {
	length := len(src)
	if length == 0 {
		return nil
	}
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

// SM2 为非对称加密, 基于ECC 该算法已公开
func SM2() {
	priv, err := sm2.GenerateKey() // 生成密钥对
	if err != nil {
		log.Fatal(err)
	}
	msg := []byte("Tongji Fintech Research Institute")
	pub := &priv.PublicKey
	ciphertxt, err := pub.Encrypt(msg)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("加密结果: %x\n", ciphertxt)
	plaintxt, err := priv.Decrypt(ciphertxt)
	if err != nil {
		log.Fatal(err)
	}
	if !bytes.Equal(msg, plaintxt) {
		log.Fatal("原文不匹配")
	}

	r, s, err := sm2.Sign(priv, msg)
	if err != nil {
		log.Fatal(err)
	}
	isok := sm2.Verify(pub, msg, r, s)
	log.Printf("Verified: %v\n", isok)
}

func main() {
	// HashSM3()
	// SM4()
	SM2()
}
