package gcm

import (
	"Critique2/block"
	"Critique2/crypto"
	"Critique2/utils"
)

// Algorithm 4 
func GCM_AE(IV []Bit, PlainText []Bit, AAD []Bit)([]Bit, []Bit){
	// Step 1
	var zeroBlock block.Block 
	zeroBlock.SetZero()
	var H block.Block = crypto.CIPH(zeroBlock)

	// Step 2 
	var J0 block.Block

	if(len(IV) == 96){
		J0 = utils.BitstringToBlocks(IV)[0]
		J0.SetBit(127, 1)
	}else{
		s := 128 * ceil(len(IV) / 128) - len(IV)
		J0 = GHASH(concatBits(concatBits(IV, make([]Bit, s + 64)), IntToBitString(len(IV), 64)))
	}

	// Step 3
	C := BlocksToBits(GCTR(utils.Inc32(J0), utils.BitstringToBlocks(PlainText)))

	// Step 4
	u := 128 * ceil(len(C) / 128) - len(C)
	v := 128 * ceil(len(AAD) / 128) - len(AAD)

	// Step 5
	tempBitString := concatBits(AAD, make([]Bit, v))
	tempBitString = concatBits(tempBitString, C)
	tempBitString = concatBits(tempBitString, make([]Bit, u))
	tempBitString = concatBits(tempBitString, IntToBitString(len(AAD), 64))
	tempBitString = concatBits(tempBitString, IntToBitString(len(C), 64))

	var S block.Block = GHASH(tempBitString)

	// Step 6
	T := BlocksToBits(GCTR(J0, S))[:TagLength]

	return C, T

}

// Algorithm 5
func GCM_AD(IV []Bits, C []Bits, AAD []Bits, T []Bits) []Bits {
	// Step 1 
	if len(T) != TagLength {
		return nil
	}

	// Step 2
	var zeroBlock block.Block
	zeroBlock.SetZero()
	H := crypto.CIPH(zeroBlock)

	// Step 3
	var J0 block.Block

	if(len(IV) == 96){
		J0 = utils.BitstringToBlocks(IV)[0]
		J0.SetBit(127, 1)
	}else{
		s := 128 * ceil(len(IV) / 128) - len(IV)
		J0 = GHASH(concatBits(concatBits(IV, make([]Bit, s + 64)), IntToBitString(len(IV), 64)))
	}

	// Step 4
	P := BlocksToBits(GCTR(utils.Inc32(J0), utils.BitstringToBlocks(C)))

	// Step 5
	u := 128 * ceil(len(C) / 128) - len(C)
	v := 128 * ceil(len(AAD) / 128) - len(AAD)

	// Step 6
	tempBitString := concatBits(AAD, make([]Bit, v))
	tempBitString = concatBits(tempBitString, C)
	tempBitString = concatBits(tempBitString, make([]Bit, u))
	tempBitString = concatBits(tempBitString, IntToBitString(len(AAD), 64))
	tempBitString = concatBits(tempBitString, IntToBitString(len(C), 64))
	
	var S block.Block = GHASH(tempBitString)

	// Step 7
	TT = BlocksToBits(GCTR(J0, S))[:TagLength]

	// Step 8
	if T != TT {
		return nil
	}else{
		return P 
	}
}
