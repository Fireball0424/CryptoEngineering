package gcm

import (
	"Critique2/block"
)

// Algorithm 2
func GHASH(H block.Block, X []block.Block) block.Block { // H is the key for GHASH
	var Y block.Block
	Y.SetZero() // Y0

	var m int = len(X)

	for i := 1; i <= m; i++ {
		Y = Y.Xor(X[i-1]) // Y[i] = Y[i-1] ⊕ X[i]
		Y = block.Multiply(Y, H)
	}
	return Y
}
