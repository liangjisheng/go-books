package main

import (
	"encoding/hex"
	"log"
)

func main() {
	src := []byte("123456")
	log.Println("src hex:", hex.EncodeToString(src))

	srcMd5 := MD5(src)
	log.Println("src md5:", hex.EncodeToString(srcMd5))

	srcRipemd160 := Ripemd160(src)
	log.Println("src ripemd160:", hex.EncodeToString(srcRipemd160))

	srcSha1 := SHA1(src)
	log.Println("src sha1:", hex.EncodeToString(srcSha1))

	srcSha256 := SHA256(src)
	log.Println("src sha256:", hex.EncodeToString(srcSha256))

	key := []byte("123456")
	srcHMAC := HMAC(key, src)
	log.Println("src HMAC:", hex.EncodeToString(srcHMAC))

	srcSha3_256 := Sha3_256(src)
	log.Println("src Sha3_256:", hex.EncodeToString(srcSha3_256))

	srcKeccak256 := Keccak256(src)
	log.Println("src Keccak256:", hex.EncodeToString(srcKeccak256))
}
