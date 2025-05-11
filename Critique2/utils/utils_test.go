package utils

import (
	"fmt"
	"testing"
)

func TestBitsToString(t *testing.T) {
	var testText string = "CryspyHashBrowns"
	bits := StringToBits(testText, false)

	if BitsToString(bits, false) != testText {
		fmt.Println(BitsToString(bits, false))
		t.Error("BitsToString test failed")
	}
}
