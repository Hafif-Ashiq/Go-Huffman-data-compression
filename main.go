package main

import (
	"compression/algo/huffman"
	"flag"
	"fmt"
	"os"
)

func compress(fileName, outputName string) {
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	err = huffman.CompressData(string(fileData), outputName)
	if err != nil {
		fmt.Println("Error compressing data:", err)
		return
	}
	fmt.Println("Created Files : " + outputName + ".huff" + " & " + outputName + "_codes.json")
}

func decompress(fileName, codesName, outputName string) {
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	jsonCodes, err := os.ReadFile(codesName)
	if err != nil {
		fmt.Println("Error reading codes:", err)
		return
	}

	err = huffman.ExpandData(fileData, jsonCodes, outputName)
	if err != nil {
		fmt.Println("Error decompressing data:", err)
		return
	}
	fmt.Println("Created Files : " + outputName + ".txt")
}

func main() {

	compressFlag := flag.Bool("c", false, "File to compress (use with two argument: input.txt and output)")
	expandFlag := flag.Bool("e", false, "Huffman compressed file and JSON file for decompression (use with three arguments: input.huff file , codes.json file and output.txt file)")

	flag.Parse()

	if *compressFlag {
		args := flag.Args()
		if len(args) < 2 {
			fmt.Println("Error: -c requires two arguments (input file and output compressed file).")
			os.Exit(1)
		}
		compress(args[0], args[1])
	} else if *expandFlag {
		args := flag.Args()
		if len(args) < 3 {
			fmt.Println("Error: -e requires two arguments (.huff file, .json file and output uncompressed file).")
			os.Exit(1)
		}
		decompress(args[0], args[1], args[2])

	} else {
		fmt.Println("Usage:")
		fmt.Println("  -c <inputfile.txt> <outputfileName>  Compress a file")
		fmt.Println("  -e <inputfile.huff> <codesfile.json> <outputfileName>  Expand a .huff file using the .json file")
		os.Exit(1)
	}

}
