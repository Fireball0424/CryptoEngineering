package gcm

import (
	"Critique2/block"
)

func inc32(b block.Block) block.Block {
	var result block.Block
	result.Hi = b.Hi

	temp := b.Lo & ((1 << 32) - 1)
	result.Lo = (b.Lo ^ temp) | ((temp + 1) & ((1 << 32) - 1))
	return result
}
