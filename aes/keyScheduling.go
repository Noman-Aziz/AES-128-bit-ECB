package aes

import (
	"strconv"
)

func doXOR(a [4]byte, b [4]byte) [4]byte {
	var c [4]byte

	for i := 0; i < 4; i++ {
		c[i] = a[i] ^ b[i]
	}

	return c
}

func hex2int(l string) int {
	switch l {
	case "0":
		return 0
	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	case "a":
		return 10
	case "b":
		return 11
	case "c":
		return 12
	case "d":
		return 13
	case "e":
		return 14
	case "f":
		return 15
	}
	return 0
}

func CircularByteLeftShift(w [4]byte) [4]byte {
	var c [4]byte = w

	var temp byte = c[0]
	for i := 0; i < 3; i++ {
		c[i] = c[i+1]
	}
	c[3] = temp

	return c
}

func ByteSubsitution(w [4]byte) [4]byte {

	var c [4]byte = w

	for i := 0; i < 4; i++ {
		//Decimal to Hex Equivilant
		hexa := strconv.FormatInt(int64(c[i]), 16)

		var firstIndex int = hex2int(string(hexa[0]))
		var secondIndex int
		if len(hexa) > 1 {
			secondIndex = hex2int(string(hexa[1]))
		} else {
			secondIndex = 0
		}

		//row * length + col
		c[i] = sbox[firstIndex*16+secondIndex]
	}

	return c
}

func AddingRoundConstant(w [4]byte, round int) [4]byte {

	var c [4]byte = w

	c[0] = c[0] ^ RoundConstants[round]
	c[1] = c[1] ^ 0x00
	c[2] = c[2] ^ 0x00
	c[3] = c[3] ^ 0x00

	return c
}

func GW(w [4]byte, round int) [4]byte {
	var g [4]byte = w

	g = CircularByteLeftShift(g)
	g = ByteSubsitution(g)
	g = AddingRoundConstant(g, round)

	return g
}

func GenerateRoundKeys(prevRoundKey []byte, round int, columnSize int, totalSize int) []byte {
	var newRoundKey = make([]byte, totalSize)

	//Seperating W Values
	var w = make([][4]byte, columnSize)

	var k int = 0
	for i := 0; i < columnSize; i++ {
		for j := 0; j < 4; j++ {
			w[i][j] = prevRoundKey[k]
			k++
		}
	}

	var gw3 [4]byte = GW(w[3], round)

	var newW = make([][4]byte, columnSize)

	//Filling New W Values
	newW[0] = doXOR(w[0], gw3)
	index := 0

	for j := 0; j < 4; j++ {
		newRoundKey[index] = newW[0][j]
		index++
	}

	prevW := 1
	for i := 1; i < columnSize; i++ {
		newW[i] = doXOR(newW[i-1], w[prevW])
		prevW++

		for j := 0; j < 4; j++ {
			newRoundKey[index] = newW[i][j]
			index++
		}
	}

	return newRoundKey
}
