package main

import (
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/curve25519"
)

func curve25519PriPubKey() {
	var priKey [32]byte
	strPriKey := "11ec18b7cffacfb46c57e027bd63f6558a78ec4ee4e929c07c6d9c68eb42c218"
	_priKey, _ := hex.DecodeString(strPriKey)
	copy(priKey[:], _priKey)
	fmt.Print("pri:")
	fmt.Println(hex.EncodeToString(priKey[:]))

	var pubKey [32]byte
	curve25519.ScalarBaseMult(&pubKey, &priKey)
	fmt.Print("pub:")
	// e5e9cf34424f55532fecac393d62bdb8b38b50635c09c7b29da65a678d3eaf47
	fmt.Println(hex.EncodeToString(pubKey[:]))
}
