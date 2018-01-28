package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"

	"github.com/bhushan5640/cryptopals/setone"
)

func main() {

	fmt.Println("Hello...")
	setone.Challenge6()

	// setone.Challenge1()
	// setone.Challenge2()
	// setone.Challenge3()
	// setone.Challenge4()
	// setone.Challenge5()
}

func encryptReapeatingKeyXor() {
	fmt.Println("\nEncrypting file...")
	lines := setone.ReadFile("data.txt")
	key := "Terminator X: "

	var buffer bytes.Buffer
	for _, x := range lines {
		buffer.WriteString(x)
		buffer.WriteString("\n")
	}

	plaintext := buffer.String()

	// Encrypt
	ciphertext := setone.Encrypt(plaintext, key)
	cipherArr, _ := hex.DecodeString(ciphertext)
	base64CipherText := base64.StdEncoding.EncodeToString(cipherArr)
	fmt.Println(base64CipherText)
	ioutil.WriteFile("encrypted.txt", []byte(base64CipherText), 0644)

	// Decrypt
	// plaintext = setone.Decrypt(ciphertext, key)
	// fmt.Println(plaintext)
}
