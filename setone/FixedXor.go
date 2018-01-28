package setone

import (
	"encoding/hex"
	"fmt"
)

// Challenge2 cryptopals challenge 2
// https://cryptopals.com/sets/1/challenges/2
func Challenge2() {
	firstInput := "1c0111001f010100061a024b53535009181c"
	secondInput := "686974207468652062756c6c277320657965"

	fmt.Println(xor(firstInput, secondInput))
}

func xor(firstInput, secondInput string) string {
	firstArr, err1 := hex.DecodeString(firstInput)
	secondArr, err2 := hex.DecodeString(secondInput)

	if err1 != nil || err2 != nil {
		return "input error"
	}

	outputArr := make([]byte, 0, len(firstArr))

	for i := 0; i < len(firstArr); i++ {
		outputArr = append(outputArr, (firstArr[i] ^ secondArr[i]))
	}

	return hex.EncodeToString(outputArr)
}
