package main

import (
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/sha3"
)

func keyGeneration(key string, keyStreamLength int) []byte {
	// use SHAKE 256 to generate key stream
	var keyStream []byte = make([]byte, keyStreamLength)
	hash := sha3.NewShake256()
	hash.Write([]byte(key))
	hash.Read(keyStream)
	return keyStream
}

func transform(fromText []byte, keyStream []byte) []byte {
	var toText []byte = make([]byte, len(fromText))

	for i := 0; i < len(fromText); i++ {
		toText[i] = fromText[i] ^ keyStream[i]
	}
	return toText
}

func main() {
	var key string = "password"
	fmt.Printf("Input Key: ")
	fmt.Scanln(&key)

	var plainTextString string = "plaintext"
	fmt.Printf("Input Plain Text: ")
	fmt.Scanln(&plainTextString)

	var plainText []byte = []byte(plainTextString)

	// encryption
	var keyStream []byte = keyGeneration(key, len(plainText))

	var cipherText []byte = make([]byte, len(plainText))
	cipherText = transform(plainText, keyStream)

	fmt.Println("cipherText: (In Hex)", hex.EncodeToString(cipherText))

	// decryption
	var keyStream2 []byte = keyGeneration(key, len(cipherText))
	var plainTextRecover []byte = make([]byte, len(cipherText))
	plainTextRecover = transform(cipherText, keyStream2)
	fmt.Println("plainText Recovering: ", string(plainTextRecover))
}
