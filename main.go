package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/noman-aziz/AES/aes"
)

type Key struct {
	Rounds     int      //10 in 128 bits, 12 in 192 bits,14 in 256 bits
	ColumnSize int      //4 in 128 bits, 6 in 192 bits, 8 in 256 bits
	TotalSize  int      //ColumnSize * 4
	RoundKeys  [][]byte //0 till Rounds
}

type PlainText struct {
	StateMatrix      [][16]byte
	NumChunks        int
	PaddingCharacter byte
	Text             []byte
}

var CipherTexts [][16]byte

func main() {

	//Custom Input Reader
	cin := bufio.NewScanner(os.Stdin)

	//Initial Variables
	var keys Key
	var plainText PlainText

	keys.Rounds = 10
	keys.ColumnSize = 4
	keys.TotalSize = 4 * keys.ColumnSize

	//Fixing Sizes for Dynamic Array
	keys.RoundKeys = make([][]byte, keys.Rounds+1)
	for i := 0; i < keys.Rounds; i++ {
		keys.RoundKeys[i] = make([]byte, keys.TotalSize)
	}

	fmt.Println("\n\tTaking Inputs\n")

	//Taking Key Input
	fmt.Printf("Enter cipher key : ")
	cin.Scan()

	//Read initial key
	keys.RoundKeys[0] = cin.Bytes()

	//Taking Text Input
	fmt.Printf("Enter plain text : ")
	cin.Scan()

	//Reading initial plain text
	plainText.Text = cin.Bytes()

	//Selecting and Displaying Padding Character
	plainText.PaddingCharacter = 'X'
	fmt.Println("Padding character is 'X'")

	//Determining Length of Plain Text and Allocating Memory Accordingly
	var temp float64 = float64(len(plainText.Text)) / 16.0
	plainText.NumChunks = int(math.Ceil(temp))
	if plainText.NumChunks == 0 {
		plainText.NumChunks = 1
	}

	plainText.StateMatrix = make([][16]byte, plainText.NumChunks)
	CipherTexts = make([][16]byte, plainText.NumChunks)

	//Seperating the chunks from PlainText
	var index int = 0
	for i := 0; i < plainText.NumChunks; i++ {
		for j := 0; j < 16; j++ {
			//Padding
			if index >= len(plainText.Text) {
				plainText.StateMatrix[i][j] = 'X'
			} else {
				plainText.StateMatrix[i][j] = plainText.Text[index]
				index++
			}
		}
	}

	//Generating Round Keys
	fmt.Println("\n\tGenerating Round Keys\n")

	for i := 0; i < keys.Rounds; i++ {
		keys.RoundKeys[i+1] = aes.GenerateRoundKeys(keys.RoundKeys[i], i, keys.ColumnSize, keys.TotalSize)
	}

	for i := 0; i <= keys.Rounds; i++ {
		fmt.Print("Round Key ", i, " : ")
		for j := 0; j < keys.TotalSize; j++ {
			fmt.Print(strconv.FormatInt(int64(keys.RoundKeys[i][j]), 16), " ")
		}
		fmt.Println()
	}

	//Perform Encryption
	fmt.Println("\n\tPerforming Encryption Process (ECB Mode)\n")

	for i := 0; i < plainText.NumChunks; i++ {
		CipherTexts[i] = aes.Encrypt(plainText.StateMatrix[i], keys.Rounds, keys.RoundKeys, keys.TotalSize)
	}

	for i := 0; i < plainText.NumChunks; i++ {
		fmt.Print("Cipher Text of Block ", i+1, " (", string(plainText.StateMatrix[i][:]), ") : ")
		for j := 0; j < 16; j++ {
			fmt.Print(strconv.FormatInt(int64(CipherTexts[i][j]), 16), " ")
		}
		fmt.Println()
	}
}
