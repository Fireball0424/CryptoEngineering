package utils

import (
	"Critique2/block"
)

type Bit bool

func BitsToBlocks(b []Bit) []block.Block {
	if len(b)%128 != 0 {
		// padding zero
		b = append(b, make([]Bit, 128-len(b)%128)...)
	}

	var blocks = make([]block.Block, len(b)/128)
	for i := 0; i < len(b)/128; i++ {
		var block block.Block
		for j := 0; j < 128; j++ {
			if b[i*128+j] {
				block.SetBit(j)
			}
		}
		blocks[i] = block
	}
	return blocks[:]
}

func BlocksToBits(b []block.Block) []Bit {
	var bits = make([]Bit, len(b)*128)
	for i := 0; i < len(b); i++ {
		for j := 0; j < 128; j++ {
			if b[i].GetMSBit(j) != 0 {
				bits[i*128+j] = true
			}
		}
	}
	return bits[:]
}

func IntToBitString(num int, length int) []Bit {
	var bits []Bit
	for num > 0 {
		bits = append(bits, Bit(num%2 == 1))
		num = num / 2
	}

	if len(bits) < length {
		bits = append(make([]Bit, length-len(bits)), bits...)
	} else if len(bits) > length {
		bits = bits[len(bits)-length:]
	}

	return bits
}

func ConcatBits(a []Bit, b []Bit) []Bit {
	return append(a, b...)
}

func Inc32(b block.Block) block.Block {
	var result block.Block
	result.Hi = b.Hi

	temp := b.Lo & ((1 << 32) - 1)
	result.Lo = (b.Lo ^ temp) | ((temp + 1) & ((1 << 32) - 1))
	return result
}

func BlocksToBytes(b []block.Block) []byte {
	// TODO: Now, we only assume it's a complete block
	var bytes []byte
	var mask uint64 = (1 << 8) - 1
	for _, block := range b {
		for shift := 56; shift >= 0; shift -= 8 {
			bytes = append(bytes, byte((block.Hi>>shift)&mask))
		}
		for shift := 56; shift >= 0; shift -= 8 {
			bytes = append(bytes, byte((block.Lo>>shift)&mask))
		}
	}
	return bytes
}

func BytesToBlocks(b []byte) []block.Block {
	var blocks []block.Block
	for i := 0; i < len(b); i += 16 {
		var block block.Block
		block.SetZero()

		// TODO: Here, we only assume the number of bytes are numbers of 8

		for j := 7; j >= 0; j-- {
			block.Hi |= uint64(b[i+(7-j)]) << (j * 8)
			block.Lo |= uint64(b[i+(15-j)]) << (j * 8)
		}
		blocks = append(blocks, block)
	}
	return blocks
}

func BitsNotEquals(a []Bit, b []Bit) bool {
	if len(a) != len(b) {
		return true
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return true
		}
	}
	return false
}

func StringToBits(s string) []Bit {
	var bits []Bit
	for _, c := range s {
		bits = append(bits, IntToBitString(int(c), 8)...)
	}
	return bits
}

func BitsToString(b []Bit) string {
	var s string
	for i := 0; i < len(b); i += 8 {
		c := 0
		for j := 0; j < 8; j++ {
			if b[i+j] {
				c |= 1 << (7 - j)
			}
		}
		s += string(rune(c))
	}
	return s
}
