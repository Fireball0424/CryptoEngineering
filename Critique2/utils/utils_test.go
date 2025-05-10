package utils

import (
	"fmt"
	"testing"
)

func TestBitsToString(t *testing.T) {
	var testText string = "CryspyHashBrowns"
	bits := StringToBits(testText)

	if BitsToString(bits) != testText {
		fmt.Println(BitsToString((bits)))
		t.Error("BitsToString test failed")
	}
}
