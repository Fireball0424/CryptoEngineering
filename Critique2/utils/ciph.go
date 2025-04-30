package utils

import (
	"Critique2/block"
	"crypto/aes"
)

// Assume CIPH is AES-128
func CIPH(Input block.Block) block.Block {
	blockCipher, err := aes.NewCipher(Key)

	if err != nil {
		panic(err.Error())
	}

	var plainText [16]byte
	var cipherText [16]byte

	BlockToBytes(Input, &plainText)

	blockCipher.Encrypt(cipherText[:], plainText[:])

	return BytesToBlock(cipherText)
}
