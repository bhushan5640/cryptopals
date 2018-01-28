package setone

import (
	"fmt"
)

// Challenge4 Cryptopals challenge 4
// https://cryptopals.com/sets/1/challenges/4
func Challenge4() {
	lines := ReadFile("setone/data/4.txt")
	var highestScore = 0
	var lineNumber = 0
	var output string
	for index, val := range lines {
		_, score, val := decryptSingleByteXor(val)
		if score > highestScore {
			highestScore = score
			lineNumber = index
			output = val
		}
		// fmt.Println(score, ":", val)
	}

	fmt.Println("\nlineNumber:", lineNumber, "\nOutput:", output)
}
