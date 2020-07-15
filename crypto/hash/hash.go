package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"io"
	"os"

	"github.com/InWeCrypto/sha3"
	"golang.org/x/crypto/ripemd160"
)

// MD5 ...
func MD5(str []byte) []byte {
	_md5 := md5.New()
	_md5.Write(str)
	return _md5.Sum([]byte(nil))
}

// Ripemd160 ...
func Ripemd160(str []byte) []byte {
	_ripemd160 := ripemd160.New()
	_ripemd160.Write(str)
	return _ripemd160.Sum([]byte(nil))
}

// MD5FILE ...
func MD5FILE(filepath string) []byte {
	f, _ := os.Open(filepath)
	defer f.Close()

	_md5 := md5.New()
	_, _ = io.Copy(_md5, f)
	return _md5.Sum([]byte(nil))
}

// SHA1 ...
func SHA1(str []byte) []byte {
	_sha1 := sha1.New()
	_sha1.Write(str)
	return _sha1.Sum([]byte(nil))
}

// SHA256 ...
func SHA256(str []byte) []byte {
	_sha256 := sha256.New()
	_sha256.Write(str)
	return _sha256.Sum([]byte(nil))
}

// Sha3_256 ...
func Sha3_256(src []byte) []byte {
	hash := sha3.Sum256(src)
	return hash[:]
}

// Keccak256 ...
func Keccak256(src []byte) []byte {
	keccak256 := sha3.NewKeccak256()
	_, _ = keccak256.Write(src)
	return keccak256.Sum([]byte(nil))
}

// HMAC ...
func HMAC(key, data []byte) []byte {
	_hmac := hmac.New(md5.New, key)
	_hmac.Write(data)
	return _hmac.Sum([]byte(nil))
}
