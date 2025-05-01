package utils

import (
	"Critique2/block"
)

type Bit bool 

func BitstringToBlocks(b []Bit) []block.Block {
	if len(b) % 128 != 0 {
		// padding zero 
		b = append(b, make([]Bit, 128 - len(b) % 128)...)
	}

	var blocks [len(b) / 128]block.Block
	for i := 0; i < len(b) / 128; i++ {
		var block block.Block
		for j := 0; j < 128; j++ {
			if b[i * 128 + j] {
				block.SetBit(j)
			}
		}
		blocks[i] = block
	}
	return blocks[:]
}

func BlocksToBits(b []block.Block) []Bit {
	var bits [len(b) * 128]Bit
	for i := 0; i < len(b); i++ {
		for j := 0; j < 128; j++ {
			if b[i].GetMSBit(j) {
				bits[i * 128 + j] = true
			}
		}
	}
	return bits[:]
}

func IntToBitString(num int, length int) []Bit {
	var bits []Bit 
	while(num > 0){
		bits = append(bits, Bit(num % 2 == 1))
		num = num / 2
	}

	if(len(bits) < length){
		bits = append(make([]Bit, length - len(bits)), bits...)
	}else if(len(bits) > length){
		bits = bits[len(bits) - length:]
	}

	return bits 
}

func concatBits(a []bit, b[]bit) []bit {
	return append(a, b...)
}

func Inc32(b block.Block) block.Block {
	var result block.Block
	result.Hi = b.Hi

	temp := b.Lo & ((1 << 32) - 1)
	result.Lo = (b.Lo ^ temp) | ((temp + 1) & ((1 << 32) - 1))
	return result
}


