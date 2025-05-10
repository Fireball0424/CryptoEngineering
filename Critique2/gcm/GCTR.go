package gcm

import (
	"Critique2/block"
	"Critique2/utils"
)

// Algorithm 3
func GCTR(ICB block.Block, Input []utils.Bit) []utils.Bit {
	var m int = len(Input)
	if m == 0 {
		return []utils.Bit{}
	}

	var CountBlock block.Block = ICB

	var Y []utils.Bit = make([]utils.Bit, m)
	for i := 0; i < m; i += 128 {
		var CIPH_K block.Block = utils.CIPH(CountBlock)
		if i+127 < m {
			copy(Y[i:i+128], utils.BlocksToBits([]block.Block{CIPH_K.Xor(utils.BitsToBlocks(Input[i : i+128])[0])}))
		} else {
			paddingLen := 128 - (m - i)
			X := append(Input[i:], make([]utils.Bit, paddingLen)...)
			copy(Y[i:], utils.BlocksToBits([]block.Block{CIPH_K.Xor(utils.BitsToBlocks(X)[0])})[0:m-i])
		}
		CountBlock = utils.Inc32(CountBlock)
	}

	return Y
}
