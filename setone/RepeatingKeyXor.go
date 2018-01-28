package setone

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

// Challenge5 Repeating Key Xor
// https://cryptopals.com/sets/1/challenges/5
func Challenge5() {
	fmt.Println("Encrypting with Repeating Key Xor")
	plaintext := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "Terminator X: Bring the noise"
	ciphertext := Encrypt(plaintext, key)
	fmt.Println(ciphertext)
}

// Encrypt Repeating key Xor encryption
func Encrypt(input, key string) string {
	firstArr := []byte(input)
	keyArr := []byte(key)
	var outputArr []byte

OuterLoop:
	for i := 0; i < len(firstArr); {
		for _, x := range keyArr {
			if i < len(firstArr) {
				y := byte(firstArr[i] ^ x)
				outputArr = append(outputArr, (y))
				i++
			} else {
				break OuterLoop
			}
		}
	}

	return hex.EncodeToString(outputArr)
}

// Decrypt repeating key xor encryption
func Decrypt(ciphertext, key string) string {
	firstArr, _ := hex.DecodeString(ciphertext)
	keyArr := []byte(key)
	var outputArr []byte

OuterLoop:
	for i := 0; i < len(firstArr); {
		for _, x := range keyArr {
			if i < len(firstArr) {
				y := byte(firstArr[i] ^ x)
				outputArr = append(outputArr, (y))
				i++
			} else {
				break OuterLoop
			}
		}
	}

	return string(outputArr)
}

func encryptFile() {
	fmt.Println("\nEncrypting file...")
	lines := ReadFile("setone/data/plaintext.txt")
	key := "Terminator X: Bring the noise"

	var buffer bytes.Buffer
	for _, x := range lines {
		buffer.WriteString(x)
		buffer.WriteString("\n")
	}

	plaintext := buffer.String()

	// Encrypt
	ciphertext := Encrypt(plaintext, key)
	fmt.Println(ciphertext)

	// Decrypt
	plaintext = Decrypt(ciphertext, key)
	fmt.Println(plaintext)
}
