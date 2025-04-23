package utils

import (
	"Critique2/block"
)

func Inc32(b block.Block) block.Block {
	var result block.Block
	result.Hi = b.Hi

	temp := b.Lo & ((1 << 32) - 1)
	result.Lo = (b.Lo ^ temp) | ((temp + 1) & ((1 << 32) - 1))
	return result
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
