package main

import (
	"fmt"
	"os"
)

func main() {

	fileData, err := os.ReadFile("big.txt")
	fmt.Println("Reading file big.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	err = compressData(string(fileData))
	if err != nil {
		fmt.Println("Error compressing data:", err)
		return
	}

}
