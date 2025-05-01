package block

// Algorithm 1
func Multiply(X Block, Y Block) Block {
	var Z [129]Block
	var V [129]Block
	var R Block = Block{Hi: (0b11100001) << 56, Lo: 0}

	Z[0].SetZero()
	V[0] = Y

	for i := 0; i < 128; i++ {
		if X.GetMSBit(i) == 0 {
			Z[i+1] = Z[i]
		} else {
			Z[i+1] = Z[i].Xor(V[i])
		}

		// avoid overflow
		if V[i].GetLSBit(0) == 0 {
			V[i+1] = V[i].ShiftRight(1)
		} else {
			V[i+1] = V[i].ShiftRight(1)
			V[i+1] = V[i+1].Xor(R)
		}
	}

	return Z[128]
}
