package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"

	// "github.com/agl/ed25519"
	// "github.com/agl/ed25519/extra25519"
	"golang.org/x/crypto/curve25519"
)

func ed25519PriPubKey() {
	pub, pri, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Print("pub:")
	fmt.Println(hex.EncodeToString(pub[:ed25519.PublicKeySize]))
	fmt.Print("pri:")
	fmt.Println(hex.EncodeToString(pri[:ed25519.PrivateKeySize]))

	// priKey := new([ed25519.PrivateKeySize]byte)
	// strPriKey := "11ec18b7cffacfb46c57e027bd63f6558a78ec4ee4e929c07c6d9c68eb42c218"
	// _priKey, _ := hex.DecodeString(strPriKey)
	// copy(priKey[:32], _priKey)

	// pubKey := ed25519.Edwards25519Pubkey(priKey)
	// 15764cd017e0da753271fa26cd529451a21b8253d001f0d43612e19ec632570a
	// fmt.Print("pub:")
	// fmt.Println(hex.EncodeToString(pubKey[:]))
	// fmt.Print("pri:")
	// fmt.Println(hex.EncodeToString(priKey[:]))
}

func extra25519PriPub() {
	var priKey [64]byte
	strPriKey := "11ec18b7cffacfb46c57e027bd63f6558a78ec4ee4e929c07c6d9c68eb42c218"
	_priKey, _ := hex.DecodeString(strPriKey)
	copy(priKey[:32], _priKey)
	fmt.Print("pri:")
	fmt.Println(hex.EncodeToString(priKey[:]))

	var curve25519Private [32]byte
	// extra25519.PrivateKeyToCurve25519(&curve25519Private, &priKey)
	fmt.Print("curvepri:")
	fmt.Println(hex.EncodeToString(curve25519Private[:]))

	var curve25519Public [32]byte
	var representative [32]byte
	// extra25519.ScalarBaseMult(&curve25519Public, &representative, &curve25519Private)
	fmt.Print("curvepub:")
	fmt.Println(hex.EncodeToString(curve25519Public[:]))
	fmt.Print("representative:")
	fmt.Println(hex.EncodeToString(representative[:]))

	curve25519.ScalarBaseMult(&curve25519Public, &curve25519Private)
	fmt.Print("curvepub:")
	fmt.Println(hex.EncodeToString(curve25519Public[:]))
}
