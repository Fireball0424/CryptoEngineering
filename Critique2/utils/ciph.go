package utils

import (
	"Critique2/block"
	"crypto/aes"
)

// Assume CIPH is AES-128
func CIPH(Input block.Block) block.Block {
	var key []byte = []byte("0123456789abcdef") // 16 bytes = AES-128
	blockCipher, err := aes.NewCipher(key)

	if err != nil {
		panic(err.Error())
	}

	var plainText [16]byte
	var cipherText [16]byte

	BlockToBytes(Input, &plainText)

	blockCipher.Encrypt(cipherText[:], plainText[:])

	return BytesToBlock(cipherText)
}
