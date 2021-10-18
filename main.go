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
	Rounds    int        //10 in 128 bits, 12 in 192 bits,14 in 256 bits
	RoundKeys [][16]byte //0 till Rounds
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

	fmt.Println("\nAES 128 Bit with ECB Mode")

	fmt.Println("\n\tTaking Inputs\n")

	//Taking Text Input
	fmt.Printf("Enter Plain Text : ")
	cin.Scan()

	//Reading initial plain text
	plainText.Text = cin.Bytes()

	//Taking Rounds Input
	fmt.Printf("Enter Num Rounds : ")
	_, err := fmt.Scan(&keys.Rounds)
	if err != nil {
		panic(err)
	}

	//Fixing Sizes for Dynamic Array
	keys.RoundKeys = make([][16]byte, keys.Rounds+1)

	//Taking Key Input
	fmt.Printf("Enter Cipher Key (16 Characters): ")
	cin.Scan()
	buffer := cin.Bytes()

	if len(buffer) != 16 {
		panic("Error, Cipher Key is not of 16 Characters, Abort!")
	}

	//Read initial key
	for i := 0; i < 16; i++ {
		keys.RoundKeys[0][i] = buffer[i]
	}

	//Selecting and Displaying Padding Character
	plainText.PaddingCharacter = 'X'
	fmt.Println("Padding Character is 'X'")

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
		keys.RoundKeys[i+1] = aes.GenerateRoundKeys(keys.RoundKeys[i], i)
	}

	for i := 0; i <= keys.Rounds; i++ {
		fmt.Print("Round Key ", i, " : ")
		for j := 0; j < 16; j++ {
			fmt.Print(strconv.FormatInt(int64(keys.RoundKeys[i][j]), 16), " ")
		}
		fmt.Println()
	}

	//Perform Encryption
	fmt.Println("\n\tPerforming Encryption Process (ECB Mode)\n")

	for i := 0; i < plainText.NumChunks; i++ {
		CipherTexts[i] = aes.Encrypt(plainText.StateMatrix[i], keys.Rounds, keys.RoundKeys)
	}

	for i := 0; i < plainText.NumChunks; i++ {
		fmt.Println("Cipher Text of Block ", i+1, " : ", string(CipherTexts[i][:]))
		fmt.Print("Cipher Text in Hex of Block ", i+1, " : ")
		for j := 0; j < 16; j++ {
			fmt.Print(strconv.FormatInt(int64(CipherTexts[i][j]), 16), " ")
		}
		fmt.Println()
	}

	//Perform Encryption
	fmt.Println("\n\tPerforming Decryption Process (ECB Mode)\n")

	for i := 0; i < plainText.NumChunks; i++ {
		temp := aes.Decrypt(CipherTexts[i], keys.Rounds, keys.RoundKeys)
		fmt.Println("Plain Text of Block ", i+1, " : ", string(temp[:]))
		fmt.Print("Plain Text in Hex of Block ", i+1, " : ")
		for j := 0; j < 16; j++ {
			fmt.Print(strconv.FormatInt(int64(temp[j]), 16), " ")
		}
		fmt.Println()
	}
}
