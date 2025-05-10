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

	var plainText []byte
	var cipherText []byte

	InputArray := []block.Block{Input}
	plainText = BlocksToBytes(InputArray)

	blockCipher.Encrypt(cipherText[:], plainText[:]) // TODO: Check

	return BytesToBlocks(cipherText)[0]
}
