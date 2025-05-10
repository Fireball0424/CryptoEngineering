package gcm

import (
	"Critique2/block"
	. "Critique2/utils"
	"math"
)

// Algorithm 4
func GCM_AE(IV []Bit, PlainText []Bit, AAD []Bit) ([]Bit, []Bit) {
	// Step 1
	var zeroBlock block.Block
	zeroBlock.SetZero()
	var H block.Block = CIPH(zeroBlock)

	// Step 2
	var J0 block.Block

	if len(IV) == 96 {
		J0 = BitsToBlocks(IV)[0]
		J0.SetBit(127)
	} else {
		s := 128*int(math.Ceil(float64(len(IV))/128)) - len(IV)
		J0 = GHASH(H, BitsToBlocks(ConcatBits(ConcatBits(IV, make([]Bit, s+64)), IntToBitString(len(IV), 64))))
	}

	// Step 3
	C := GCTR(Inc32(J0), PlainText)

	// Step 4
	u := 128*int(math.Ceil(float64(len(C))/128)) - len(C)
	v := 128*int(math.Ceil(float64(len(AAD))/128)) - len(AAD)

	// Step 5
	tempBitString := ConcatBits(AAD, make([]Bit, v))
	tempBitString = ConcatBits(tempBitString, C)
	tempBitString = ConcatBits(tempBitString, make([]Bit, u))
	tempBitString = ConcatBits(tempBitString, IntToBitString(len(AAD), 64))
	tempBitString = ConcatBits(tempBitString, IntToBitString(len(C), 64))

	var S block.Block = GHASH(H, BitsToBlocks(tempBitString))

	// Step 6
	T := GCTR(J0, BlocksToBits([]block.Block{S}))[:TagLength] // TODO: Check

	return C, T

}

// Algorithm 5
func GCM_AD(IV []Bit, C []Bit, AAD []Bit, T []Bit) []Bit {
	// Step 1
	if len(T) != TagLength {
		return nil
	}

	// Step 2
	var zeroBlock block.Block
	zeroBlock.SetZero()
	H := CIPH(zeroBlock)

	// Step 3
	var J0 block.Block

	if len(IV) == 96 {
		J0 = BitsToBlocks(IV)[0]
		J0.SetBit(127)
	} else {
		s := 128*int(math.Ceil(float64(len(IV))/128)) - len(IV)
		J0 = GHASH(H, BitsToBlocks(ConcatBits(ConcatBits(IV, make([]Bit, s+64)), IntToBitString(len(IV), 64))))
	}

	// Step 4
	P := GCTR(Inc32(J0), C)

	// Step 5
	u := 128*int(math.Ceil(float64(len(C))/128)) - len(C)
	v := 128*int(math.Ceil(float64(len(AAD))/128)) - len(AAD)

	// Step 6
	tempBitString := ConcatBits(AAD, make([]Bit, v))
	tempBitString = ConcatBits(tempBitString, C)
	tempBitString = ConcatBits(tempBitString, make([]Bit, u))
	tempBitString = ConcatBits(tempBitString, IntToBitString(len(AAD), 64))
	tempBitString = ConcatBits(tempBitString, IntToBitString(len(C), 64))

	var S block.Block = GHASH(H, BitsToBlocks(tempBitString))

	// Step 7
	TT := GCTR(J0, BlocksToBits([]block.Block{S}))[:TagLength]

	// Step 8
	if BitsNotEquals(T, TT) {
		return nil
	} else {
		return P
	}
}
