package setone

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

// Challenge1 Cryptopals challenge 1
// https://cryptopals.com/sets/1/challenges/1
func Challenge1() {
	fmt.Println("Lets convert hex to base64")
	hexInput := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	fmt.Println(hexInput)
	fmt.Println(hexToBase64(hexInput))
}

func hexToBase64(hexInput string) string {
	byteArr, err := hex.DecodeString(hexInput)

	if err != nil {
		log.Fatal("Error decoding hex string")
	}

	base64Output := base64.StdEncoding.EncodeToString(byteArr)

	return base64Output
}

func base64ToHex(base64Input string) string {
	byteArr, err := base64.StdEncoding.DecodeString(base64Input)

	if err != nil {
		log.Fatal("Error decoding base64 string")
	}

	hexOutput := hex.EncodeToString(byteArr)

	return hexOutput
}
