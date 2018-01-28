package setone

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// S char frequency
//https://en.wikipedia.org/wiki/Letter_frequency
var S = " EARIOTNSLCUDPMHGBFYWKVXZJQ"

// CharScore character score
var CharScore = initializeCharScore()

// InitializeCharScore Initialize character score
func initializeCharScore() map[byte]int {
	charScore := make(map[byte]int, len(S))
	a := []rune(S)
	l := len(S)
	for _, r := range a {
		charScore[byte(r)] = l
		l--
	}
	return charScore
}

// Challenge3 cryptopals challenge 3
// https://cryptopals.com/sets/1/challenges/3
func Challenge3() {
	initializeCharScore()
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	key, _, output := decryptSingleByteXor(input)
	fmt.Println("Key is :", key)
	fmt.Println("Output:", output)
}

// decryptSingleByteXor  Decrypt Single Byte XOR
// it returns key, score and decrypted string
func decryptSingleByteXor(input string) (byte, int, string) {
	firstArr, err := hex.DecodeString(input)

	if err != nil {
		fmt.Println("input error")
	}
	// fmt.Println(firstArr)
	return detectSingleByteXor(firstArr)
}

func detectSingleByteXor(firstArr []byte) (byte, int, string) {
	var x byte = 1
	var highestScore = 0
	var key = byte(0)
	var output string
	for ; x < byte(255); x++ {
		var score = 0
		outputArr := make([]byte, 0, len(firstArr))
		for i := 0; i < len(firstArr); i++ {
			y := byte(firstArr[i] ^ x)
			outputArr = append(outputArr, (y))

			var tempKey = (strings.ToUpper(string(y)))[0]

			if val, ok := CharScore[tempKey]; ok {
				score = score + val
			} else {
				score = score - 20
			}
		}
		if score > highestScore {
			highestScore = score
			key = x
			output = string(outputArr)
		}
		// fmt.Println(x, ":", score)
		// fmt.Println(string(outputArr))
	}

	return key, highestScore, output
}
