package aes

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

func CheckOverflow(a byte, b byte, result byte) bool {
	if a == 0 || b == 0 {
		return false
	}

	if a == result/b {
		return false
	} else {
		return true
	}

}

func MultiplicationWithOverflowCheck(a byte, b byte) byte {

	var result byte

	if a == 0x01 {
		result = b
	} else if a == 0x02 {
		result = a * b

		if CheckOverflow(a, b, result) {
			result ^= 0x1B
		}
	} else if a == 0x03 {
		result = 0x02 * b

		if CheckOverflow(a, b, result) {
			result ^= 0x1B
		}

		result ^= b

	}

	return result
}
