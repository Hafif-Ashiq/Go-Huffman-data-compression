package huffman

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

func getCodesFromString(codes []byte) (map[rune]string, error) {

	allCodes := make(map[rune]string)

	err := json.Unmarshal(codes, &allCodes)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error unmarshalling the codes")
	}

	return allCodes, nil

}

func generateTreeFromCodes(codes map[rune]string) *Node {

	root := new(Node)

	for k, v := range codes {
		iter := root
		for i, str := range v {

			if i < len(v)-1 {
				if str == '0' {
					if iter.left == nil {
						iter.left = new(Node)
					}
					iter = iter.left
				} else if str == '1' {
					if iter.right == nil {
						iter.right = new(Node)
					}
					iter = iter.right

				}
			} else {
				if str == '0' {
					iter.left = new(Node)
					iter.left.character = k
					iter.left.isChar = true
				} else if str == '1' {
					iter.right = new(Node)
					iter.right.character = k
					iter.right.isChar = true

				}
			}

		}
	}

	return root
}

// PrintTree prints the Huffman tree nicely
func PrintTree(node *Node, prefix string, isLeft bool) {
	if node == nil {
		return
	}

	// Print the current node
	fmt.Print(prefix)
	if isLeft {
		fmt.Print("├── ")
	} else {
		fmt.Print("└── ")
	}

	if node.isChar {
		fmt.Printf("%c\n", node.character)
	} else {
		fmt.Println("*")
	}

	// Prepare the new prefix
	newPrefix := prefix
	if isLeft {
		newPrefix += "│   "
	} else {
		newPrefix += "    "
	}

	// Important: Always print left then right
	PrintTree(node.left, newPrefix, true)
	PrintTree(node.right, newPrefix, false)
}

func ByteToBits(b byte) []int8 {
	bits := make([]int8, 8)

	for i := 7; i >= 0; i-- {
		bits[7-i] = int8((b >> i) & 1)
	}

	return bits
}

func ExpandData(compressedString, codeString []byte, outputName string) error {

	codes, err := getCodesFromString(codeString)
	if err != nil {
		return err
	}
	// Generate tree from codes
	root := generateTreeFromCodes(codes)
	// Print Tree
	// PrintTree(root, "", false)
	// fmt.Println("-----------------")

	// Get padding Bits from the start of compressed file and return the original string
	paddingBits := int(compressedString[0])
	compressedString = compressedString[1:]

	// Create an int array
	fileBytes := make([]int8, (len(compressedString)*8)-paddingBits)

	// fmt.Println(fileBytes)

	for i, v := range compressedString {
		bits := ByteToBits(v)
		for j, b := range bits {
			if i*8+j >= len(fileBytes) {
				break
			}
			fileBytes[i*8+j] = b
		}

	}

	// Create string of the compressed data to get it to original format
	iter := root
	var fileData strings.Builder
	for _, v := range fileBytes {
		if v == 0 {
			if iter.left != nil && iter.left.isChar {
				fileData.WriteRune(iter.left.character)
				iter = root
			} else if iter.left != nil {
				iter = iter.left
			} else {
				break
			}
		} else {
			if iter.right != nil && iter.right.isChar {
				fileData.WriteRune(iter.right.character)
				iter = root
			} else if iter.right != nil {
				iter = iter.right
			} else {
				break
			}
		}

	}

	// Write data to file
	err = os.WriteFile(outputName+".txt", []byte(fileData.String()), 0666)
	if err != nil {
		return errors.New("error in decompressing the file ")
	}

	return nil
}
