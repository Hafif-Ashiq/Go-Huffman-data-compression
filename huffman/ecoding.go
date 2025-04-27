package huffman

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Node struct {
	character rune
	frequency int
	isChar    bool
	left      *Node
	right     *Node
}

func generateTree(data string) *Node {
	// Implement tree generation here
	mapData := make(map[rune]int)
	for _, char := range data {
		mapData[char] += 1
	}

	var nodes []*Node

	for char, freq := range mapData {
		nodes = append(nodes, &Node{character: char, frequency: freq, isChar: true})
	}

	sort.Slice(nodes, func(i int, j int) bool {
		return nodes[i].frequency < nodes[j].frequency
	})

	for len(nodes) > 1 {
		firstNode := nodes[0]
		secondNode := nodes[1]

		newNode := &Node{
			frequency: firstNode.frequency + secondNode.frequency,
			left:      firstNode,
			right:     secondNode,
		}

		nodes = append(nodes[2:], newNode)
		sort.Slice(nodes, func(i int, j int) bool {
			return nodes[i].frequency < nodes[j].frequency
		})
	}

	if len(nodes) == 0 {
		return nil
	}

	return nodes[0]
}

func generateCodes(node *Node, prefix string, codes map[rune]string) {
	// Implement code generation here

	if node.isChar {
		codes[node.character] = prefix
		return
	}

	if node.left != nil {
		generateCodes(node.left, prefix+"0", codes)
	}
	if node.right != nil {
		generateCodes(node.right, prefix+"1", codes)
	}

}

func saveCodesToFile(codes map[rune]string, outputName string) error {
	// Implement code saving here
	json, err := json.Marshal(codes)

	if err != nil {
		return err
	}

	err = os.WriteFile(outputName+"_codes.json", json, 0666)
	if err != nil {
		return err
	}
	return nil

}

func saveCompressedDataToFile(data string, codes map[rune]string, output string) error {
	// Implement compressed data saving here
	var compressedData strings.Builder
	for _, char := range data {
		// if i%10000 == 0 {
		// 	fmt.Printf("\rCharacters done %d/%d", i, len(data))
		// }
		compressedData.WriteString(codes[char])
	}

	// convert into bytes
	paddingBits := 0
	for len(compressedData.String())%8 != 0 {
		paddingBits++
		compressedData.WriteString("0")
	}

	var bytes []byte
	var currentByte byte
	bitCount := 0

	fmt.Println(len(compressedData.String()))
	fmt.Println(paddingBits)
	bytes = append(bytes, byte(paddingBits))
	for _, char := range compressedData.String() {
		currentByte <<= 1
		if char == '1' {
			currentByte |= 1
		}
		bitCount++

		if bitCount == 8 {
			bytes = append(bytes, currentByte)
			bitCount = 0
			currentByte = 0
		}
	}

	// write to file

	fmt.Println("Compressed data Created")
	err := os.WriteFile(output+".huff", []byte(bytes), 0666)
	if err != nil {
		return err
	}
	return err
}

func CompressData(data, outputName string) error {
	// Implement Huffman coding here
	// Generate Tree  Done
	// Generate Codes
	// Save Codes to File
	// Save compressed Data to different File

	tree := generateTree(data)
	fmt.Println("Tree generated")
	codes := make(map[rune]string)
	generateCodes(tree, "", codes)
	fmt.Println("Codes generated")
	err := saveCodesToFile(codes, outputName)
	fmt.Println("Codes saved")
	if err != nil {
		return err
	}

	err = saveCompressedDataToFile(data, codes, outputName)
	fmt.Println("Compressed data saved")
	if err != nil {
		return err
	}

	return nil
}
