package setone

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

// ReadFile reads file line by line in string array
func ReadFile(location string) []string {
	fileLoc, _ := filepath.Abs(location)
	file, err := os.Open(fileLoc)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

// FindHammingDistance find hamming distance between two strings
func FindHammingDistance(first, second string) int {
	firstArr := []byte(first)
	secondArr := []byte(second)
	return FindHammingDistanceByteArr(firstArr, secondArr)
}

// FindHammingDistanceByteArr find hamming distance between two byte array
func FindHammingDistanceByteArr(firstArr, secondArr []byte) int {
	var diff int

	if len(firstArr) != len(secondArr) {
		log.Fatal("Array size not same")
	}

	for i := 0; i < len(firstArr); i++ {
		for j := 0; j < 8; j++ {
			a := firstArr[i]
			b := secondArr[i]
			mask := byte(1 << uint(j))
			if (a & mask) != (b & mask) {
				diff++
			}
		}
	}
	return diff
}

// GetBitAtIndex get bit at the given index of byte
func GetBitAtIndex(x byte, i int) bool {
	mask := byte(1 << uint(i))
	return ((x & mask) != 0)
}
