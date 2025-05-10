package main

import (
	"Critique2/gcm"
	"Critique2/utils"
)

func main() {
	var plainText string = "This is plain text"
	var AAD string = "This is authenticated associated data"

	plainTextBits := utils.StringToBits(plainText)
	AADBits := utils.StringToBits(AAD)

	IV := utils.IntToBitString(0, 96)

	C, T := gcm.GCM_AE(IV, plainTextBits, AADBits)
	result := gcm.GCM_AD(IV, C, AADBits, T)

	if result == nil {
		println("Authentication failed")
	} else {
		println("Authentication succeeded")
		println(utils.BitsToString(result))
	}
}
