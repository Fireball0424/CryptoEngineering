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

func BlockToBytes(b block.Block, arr *[16]byte) {
	for i := 0; i < 8; i++ {
		arr[i] = byte(b.Hi >> (56 - i*8))
		arr[i+8] = byte(b.Lo >> (56 - i*8))
	}
}

func BytesToBlock(arr [16]byte) block.Block {
	var b block.Block
	b.SetZero()
	for i := 0; i < 8; i++ {
		b.Hi |= uint64(arr[i]) << (56 - i*8)
		b.Lo |= uint64(arr[i+8]) << (56 - i*8)
	}
	return b
}
