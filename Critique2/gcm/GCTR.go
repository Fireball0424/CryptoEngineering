package gcm

import (
	"Critique2/block"
	"Critique2/utils"
)

func GCTR(ICB block.Block, Input []block.Block) []block.Block {
	var m int = len(Input)
	var CountBlock block.Block = ICB

	var Y []block.Block = make([]block.Block, m)
	for i := 1; i <= m; i++ {
		var CIPH_K block.Block = utils.CIPH(CountBlock)
		Y[i-1] = CIPH_K.Xor(Input[i-1])
		CountBlock = inc32(CountBlock)
	}
	return Y
}
