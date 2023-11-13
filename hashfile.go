package main

import (
	"crypto/sha1"
	"fmt"
	"os"
)

func main() {
	// read a file to a byte slice
	data, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	hash := sha1.Sum(data)
	fmt.Printf("%x\n", hash)
}
