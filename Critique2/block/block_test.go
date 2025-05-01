package block

import (
	"fmt"
	"testing")

func TestSetZero(t *testing.T) {
	a := Block{Hi: 0xFFFFFFFFFFFFFFFF, Lo: 0xFFFFFFFFFFFFFFFF}

	fmt.Println("Before SetZero:")
	fmt.Printf("Hi: %064b\n", a.Hi)
	fmt.Printf("Lo: %064b\n", a.Lo)

	a.SetZero()
	
	fmt.Println("After SetZero:")
	fmt.Printf("Hi: %064b\n", a.Hi)
	fmt.Printf("Lo: %064b\n", a.Lo)
}

func TestSetBit(t *testing.T) {
	a := Block{Hi: 0, Lo: 0}
	fmt.Println("Before SetBit:")
	fmt.Printf("Hi: %064b\n", a.Hi)
	fmt.Printf("Lo: %064b\n", a.Lo)

	a.SetBit(0)
	a.SetBit(63)
	a.SetBit(64)
	a.SetBit(127)

	fmt.Println("After SetBit:")
	fmt.Printf("Hi: %064b\n", a.Hi)
	fmt.Printf("Lo: %064b\n", a.Lo)
}

func TestGetMSBit(t *testing.T) {
	a := Block{Hi: 0, Lo: 1}
	fmt.Println("Before GetMSBit:")
	fmt.Printf("Hi: %064b\n", a.Hi)
	fmt.Printf("Lo: %064b\n", a.Lo)

	fmt.Println("GetMSBit:")
	fmt.Printf("Hi: ")
	for i := 0; i < 64; i++ {
		fmt.Printf("%d", a.GetMSBit(i))
	}
	fmt.Println()

	fmt.Printf("Lo: ")
	for i := 64; i < 128; i++ {
		fmt.Printf("%d", a.GetMSBit(i))
	}
	fmt.Println()
}

func TestGetLSBit(t *testing.T) {
	a := Block{Hi: 0, Lo: 1}
	fmt.Println("Before GetLSBit:")
	fmt.Printf("Hi: %064b\n", a.Hi)
	fmt.Printf("Lo: %064b\n", a.Lo)

	fmt.Println("GetLSBit:")
	fmt.Printf("Hi: ")
	for i := 127; i >= 64; i-- {
		fmt.Printf("%d", a.GetLSBit(i))
	}
	fmt.Println()

	fmt.Printf("Lo: ")
	for i := 63; i >= 0; i-- {
		fmt.Printf("%d", a.GetLSBit(i))
	}
	fmt.Println()
}
func TestShiftRight(t *testing.T) {
	a := Block{Hi: 2, Lo: 0}
	fmt.Println("Before ShiftRight:")
	fmt.Printf("Hi: %064b\n", a.Hi)
	fmt.Printf("Lo: %064b\n", a.Lo)

	b := a.ShiftRight(1)
	fmt.Println("After ShiftRight by 1:")
	fmt.Printf("Hi: %064b\n", b.Hi)
	fmt.Printf("Lo: %064b\n", b.Lo)

	c := a.ShiftRight(32)
	fmt.Println("After ShiftRight by 32:")
	fmt.Printf("Hi: %064b\n", c.Hi)
	fmt.Printf("Lo: %064b\n", c.Lo)

	d := a.ShiftRight(66)
	fmt.Println("After ShiftRight by 66:")
	fmt.Printf("Hi: %064b\n", d.Hi)
	fmt.Printf("Lo: %064b\n", d.Lo)
}
func TestXor(t *testing.T) {
	a := Block{Hi: 0xF0F0F0F0F0F0F0F0, Lo: 0xFFFFFFFFFFFFFFFF}
	b := Block{Hi: 0x0000000000000000, Lo: 0x0000000000000000}

	fmt.Println("Before Xor:")
	fmt.Printf("a Hi: %064b\n", a.Hi)
	fmt.Printf("a Lo: %064b\n", a.Lo)
	fmt.Printf("b Hi: %064b\n", b.Hi)
	fmt.Printf("b Lo: %064b\n", b.Lo)

	c := a.Xor(b)
	fmt.Println("After Xor:")
	fmt.Printf("c Hi: %064b\n", c.Hi)
	fmt.Printf("c Lo: %064b\n", c.Lo)
}
func TestMultiply(t *testing.T) {
	a := Block{Hi: 0x0000000000000000, Lo: 0x0000000000000000}
	b := Block{Hi: 0x0000000000000000, Lo: 0x0000000000000001}

	fmt.Println("Before Multiply:")
	fmt.Printf("a Hi: %064b\n", a.Hi)
	fmt.Printf("a Lo: %064b\n", a.Lo)
	fmt.Printf("b Hi: %064b\n", b.Hi)
	fmt.Printf("b Lo: %064b\n", b.Lo)

	c := Multiply(a, b)
	fmt.Println("After Multiply:")
	fmt.Printf("c Hi: %064b\n", c.Hi)
	fmt.Printf("c Lo: %064b\n", c.Lo)
}