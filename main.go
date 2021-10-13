package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {

	cin := bufio.NewScanner(os.Stdin)
	cin.Scan()

	var arr []byte = cin.Bytes()

	for _, v := range arr {
		hexa := strconv.FormatInt(int64(v), 16)
		print(hexa, " ")
	}

}
