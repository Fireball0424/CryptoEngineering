package block

// block is a 128 bits integer
type Block struct {
	Hi uint64
	Lo uint64
}

// SetZero sets the block to zero
func (b *Block) SetZero() {
	b.Hi = 0
	b.Lo = 0
}

// set the i-th bit (count from left) 0-based index
func (b *Block) SetBit(i int) { 
	if i < 64 {
		b.Hi |= (1 << (63 - i))
	} else {
		b.Lo |= (1 << (127 - i))
	}
}

// get the i-th bit (count from left) 0-based index 
func (b *Block) GetMSBit(i int) int {
	if i < 64 {
		return int((b.Hi >> (63 - i)) & 1)
	} else {
		return int((b.Lo >> (127 - i)) & 1)
	}
}

// get the i-th bit (count from right) 0-based index
func (b *Block) GetLSBit(i int) int {
	if i < 64 {
		return int((b.Lo >> i) & 1)
	} else {
		return int((b.Hi >> (i - 64)) & 1)
	}
}

func (b *Block) ShiftRight(i int) Block {
	var result Block = Block{Hi: 0, Lo: 0}

	if i < 64 {
		result.Hi = b.Hi >> i

		temp := b.Hi & ((1 << i) - 1)
		result.Lo = (b.Lo >> i) | (temp << (64 - i))
	} else {
		result.Hi = 0
		result.Lo = b.Hi >> (i - 64)
	}
	return result
}

func (b *Block) Xor(other Block) Block {
	return Block{Hi: b.Hi ^ other.Hi, Lo: b.Lo ^ other.Lo}
}


