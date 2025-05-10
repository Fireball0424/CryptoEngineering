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
	} else {
		println("GCTR test passed")
	}
}

func TestGCM(t *testing.T) {
	var IV = make([]utils.Bit, 96)

	var plainText string = "This is plain text"
	var AAD string = "This is authenticated associated data"

	plainTextBits := utils.StringToBits(plainText)
	AADBits := utils.StringToBits(AAD)
	C, T := GCM_AE(IV, plainTextBits, AADBits)

	result := GCM_AD(IV, C, AADBits, T)

	if result == nil {
		t.Error("GCM test failed")
	} else {
		println("GCM test passed")
	}
}
