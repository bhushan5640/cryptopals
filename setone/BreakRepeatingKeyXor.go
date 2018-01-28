package setone

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"math"
)

// Challenge6 Cryptopals challenge 6
// https://cryptopals.com/sets/1/challenges/6
func Challenge6() {
	lines := ReadFile("setone/data/6.txt")
	breakRepeatingKeyXor(lines)
}

func breakRepeatingKeyXor(cipherTextLines []string) {
	minKeySize := 2
	maxKeySize := 60
	var buffer bytes.Buffer
	lowestEditDistance := math.MaxFloat64
	var probableKeySize int
	numberOfBlocks := 8

	for _, line := range cipherTextLines {
		buffer.WriteString(line)
	}
	base64CipherText := buffer.String()
	cipherTextByteArr, err := base64.StdEncoding.DecodeString(base64CipherText)
	if err != nil {
		log.Fatal("error decoding base64 input")
	}

	for keySize := minKeySize; keySize <= maxKeySize && keySize*numberOfBlocks < len(cipherTextByteArr); keySize++ {
		editDistance := 0.0
		for j := 0; j < numberOfBlocks; j++ {
			start := (2 * keySize * j)
			firstBytes := cipherTextByteArr[start:(start + keySize)]
			secondBytes := cipherTextByteArr[(start + keySize):(start + keySize + keySize)]
			editDistance = editDistance + float64((FindHammingDistanceByteArr(firstBytes, secondBytes))/keySize)
		}

		if editDistance < lowestEditDistance {
			probableKeySize = keySize
			lowestEditDistance = editDistance
		}
	}

	log.Print("Probable Key size: ", probableKeySize)

	singleByteXorBlocks := make([][]byte, probableKeySize)
	finalKeyBytes := make([]byte, 0, probableKeySize)
	for k := 0; k < probableKeySize; k++ {
		for i := k; i < len(cipherTextByteArr); i = i + probableKeySize {
			singleByteXorBlocks[k] = append(singleByteXorBlocks[k], cipherTextByteArr[i])
		}
		singleByteKey, _, _ := detectSingleByteXor(singleByteXorBlocks[k])
		finalKeyBytes = append(finalKeyBytes, singleByteKey)
	}

	finalKey := string(finalKeyBytes)
	fmt.Println("Encryption key is:", finalKey)
	fmt.Println("Now lets decrypt the text....")

	plaintext := Decrypt(hex.EncodeToString(cipherTextByteArr), finalKey)
	fmt.Println(plaintext)
}
