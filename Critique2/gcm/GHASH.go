package gcm

import (
	"Critique2/block"
)

// Algorithm 2
func GHASH(H block.Block, X []block.Block) block.Block {
	var Y block.Block
	Y.SetZero()

	var m int = len(X)

	for i := 1; i <= m; i++ {
		Y = Y.Xor(X[i-1])
		Y = block.Multiply(Y, H)
	}
	return Y
}
