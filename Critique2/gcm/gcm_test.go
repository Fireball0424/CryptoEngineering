package gcm

import (
	"Critique2/utils"
	"testing"
)

func TestGCTR(t *testing.T) {
	var testText string = "CryspyHashBrowns"
	Input := utils.StringToBits(testText)
	icb := utils.BitsToBlocks(utils.IntToBitString(0, 96))[0]
	res := GCTR(icb, Input)
	dec_res := GCTR(icb, res)

	if utils.BitsNotEquals(dec_res, Input) {
		t.Error("GCTR test failed")
	}
}
