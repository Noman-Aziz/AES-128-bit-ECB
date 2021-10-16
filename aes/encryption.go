package aes

import (
	"fmt"
	"strconv"
)

func AddRoundKey(stateMatrix [16]byte, roundKey []byte) [16]byte {
	var newStateMatrix [16]byte

	for i := 0; i < 16; i++ {
		newStateMatrix[i] = stateMatrix[i] ^ roundKey[i]
	}

	return newStateMatrix
}

func BytesSubstitution(stateMatrix [16]byte) [16]byte {

	var newStateMatrix [16]byte = stateMatrix

	for i := 0; i < 16; i++ {
		//Decimal to Hex Equivalent
		hexa := strconv.FormatInt(int64(newStateMatrix[i]), 16)

		var firstIndex int
		var secondIndex int
		if len(hexa) > 1 {
			firstIndex = hex2int(string(hexa[0]))
			secondIndex = hex2int(string(hexa[1]))
		} else {
			firstIndex = 0
			secondIndex = hex2int(string(hexa[0]))
		}

		//row * length + col
		newStateMatrix[i] = sbox[firstIndex*16+secondIndex]
	}

	return newStateMatrix
}

func ShiftRows(stateMatrix [16]byte) [16]byte {
	var newStateMatrix [16]byte = stateMatrix

	//Converting Col Major into Row Major
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			stateMatrix[i*4+j] = newStateMatrix[j*4+i]
		}
	}

	//Offset 0
	newStateMatrix[0] = stateMatrix[0]
	newStateMatrix[1] = stateMatrix[1]
	newStateMatrix[2] = stateMatrix[2]
	newStateMatrix[3] = stateMatrix[3]

	//Offset 1
	newStateMatrix[4] = stateMatrix[5]
	newStateMatrix[5] = stateMatrix[6]
	newStateMatrix[6] = stateMatrix[7]
	newStateMatrix[7] = stateMatrix[4]

	//Offset 2
	newStateMatrix[8] = stateMatrix[10]
	newStateMatrix[9] = stateMatrix[11]
	newStateMatrix[10] = stateMatrix[8]
	newStateMatrix[11] = stateMatrix[9]

	//Offset 3
	newStateMatrix[12] = stateMatrix[15]
	newStateMatrix[13] = stateMatrix[12]
	newStateMatrix[14] = stateMatrix[13]
	newStateMatrix[15] = stateMatrix[14]

	//Converting Row Major back to Col Major
	var temp [16]byte = newStateMatrix
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			newStateMatrix[i*4+j] = temp[j*4+i]
		}
	}

	return newStateMatrix
}

func MixColumns(stateMatrix [16]byte) [16]byte {
	var newStateMatrix [16]byte = stateMatrix

	var M [16]byte = [16]byte{0x02, 0x03, 0x01, 0x01, 0x01, 0x02, 0x03, 0x01, 0x01, 0x01, 0x02, 0x03, 0x03, 0x01, 0x01, 0x02}

	//Converting Col Major into Row Major
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			stateMatrix[i*4+j] = newStateMatrix[j*4+i]
		}
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			newStateMatrix[i*4+j] = 0
			for k := 0; k < 4; k++ {
				newStateMatrix[i*4+j] ^= MultiplicationWithOverflowCheck(M[i*4+k], stateMatrix[k*4+j])
				if i == 0 && j == 0 {
					fmt.Println(strconv.FormatInt(int64(M[i*4+k]), 16), strconv.FormatInt(int64(stateMatrix[k*4+j]), 16), strconv.FormatInt(int64(MultiplicationWithOverflowCheck(M[i*4+k], stateMatrix[k*4+j])), 16))
				}
			}
		}
	}

	//Converting Row Major back to Col Major
	var temp [16]byte = newStateMatrix
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			newStateMatrix[i*4+j] = temp[j*4+i]
		}
	}

	return newStateMatrix
}

func Encrypt(stateMatrix [][16]byte, numChunks int, rounds int, roundKeys [][]byte, totalSize int) [16]byte {
	//var cipherText []byte = make([]byte, numChunks*16)
	var tempStateMatrix [16]byte

	//For Every Chunk
	for i := 0; i < numChunks; i++ {
		tempStateMatrix = stateMatrix[i]

		//Round 0
		tempStateMatrix = AddRoundKey(tempStateMatrix, roundKeys[0])

		//Rest Rounds
		for j := 1; j <= rounds; j++ {

			//Substituting Bytes
			tempStateMatrix = BytesSubstitution(tempStateMatrix)

			//Shifting Rows
			tempStateMatrix = ShiftRows(tempStateMatrix)

			for k := 0; k < 16; k++ {
				fmt.Print(strconv.FormatInt(int64(tempStateMatrix[k]), 16), " ")
			}
			fmt.Println()

			//Except Last Round
			if j != rounds {
				//Mixing Columns
				tempStateMatrix = MixColumns(tempStateMatrix)
			}

			for k := 0; k < 16; k++ {
				fmt.Print(strconv.FormatInt(int64(tempStateMatrix[k]), 16), " ")
			}
			fmt.Println()

			//Adding Round Key
			tempStateMatrix = AddRoundKey(tempStateMatrix, roundKeys[j])
		}
	}

	var cipherText [16]byte = tempStateMatrix
	return cipherText
}
