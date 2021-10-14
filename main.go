package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/noman-aziz/AES/aes"
)

type Key struct {
	Rounds     int //10 in 128 bits, 12 in 192 bits,14 in 256 bits
	ColumnSize int //4 in 128 bits, 6 in 192 bits, 8 in 256 bits
	TotalSize  int //ColumnSize * 4
	RoundKeys  [][]byte
}

func main() {

	//Custom Input Reader
	cin := bufio.NewScanner(os.Stdin)

	//Initial Variables
	var keys Key
	keys.Rounds = 10
	keys.ColumnSize = 4
	keys.TotalSize = 4 * keys.ColumnSize

	//Fixing Sizes for Dynamic Array
	keys.RoundKeys = make([][]byte, keys.Rounds+1)
	for i := 0; i < keys.Rounds; i++ {
		keys.RoundKeys[i] = make([]byte, keys.TotalSize)
	}

	//Taking Key Input
	fmt.Printf("Enter cipher key : ")
	cin.Scan()

	//Read initial key
	keys.RoundKeys[0] = cin.Bytes()

	for i := 0; i < keys.Rounds; i++ {
		keys.RoundKeys[i+1] = aes.GenerateRoundKeys(keys.RoundKeys[i], i, keys.ColumnSize, keys.TotalSize)
	}

	for i := 0; i <= keys.Rounds; i++ {
		fmt.Print("Round ", i, " : ")
		for j := 0; j < keys.TotalSize; j++ {
			fmt.Print(strconv.FormatInt(int64(keys.RoundKeys[i][j]), 16), " ")
		}
		fmt.Println()
	}
}
