package main

import (
	"fmt"
	"os"
)

func main() {
	// The expected bytes after transformation
	expected := []byte{
		0x8f, 0x7a, 0x23, 0x45,
		0x98, 0xab, 0xcd, 0xef,
		0x12, 0x34, 0x56, 0x78,
		0x9a, 0xbc, 0xde, 0xf0,
	}

	// We need to run the crackme first to get a canary value
	fmt.Println("First, run the crackme and enter any 16 valid characters.")
	fmt.Println("Enter the first 8 bytes of transformed output (in hex, e.g., 8f7a234598abcdef): ")
	
	var hexInput string
	fmt.Scanln(&hexInput)
	
	if len(hexInput) != 16 {
		fmt.Println("Invalid input length")
		os.Exit(1)
	}

	// Convert hex string to bytes
	firstBlock := make([]byte, 8)
	for i := 0; i < 8; i++ {
		fmt.Sscanf(hexInput[i*2:i*2+2], "%02x", &firstBlock[i])
	}

	// Now we can recover the canary by XORing with known values
	canary := make([]byte, 8)
	testInput := []byte("AAAAAAAA") // We know we input this
	
	for i := 0; i < 8; i++ {
		canary[i] = firstBlock[i] ^ testInput[i] ^ byte(i*7)
	}

	fmt.Println("\nRecovered canary!")
	
	// Now we can reverse the transformation to find the required input
	solution := make([]byte, 16)
	for i := 0; i < 16; i++ {
		canaryByte := canary[i%8]
		solution[i] = expected[i] ^ canaryByte ^ byte(i*7)
	}

	fmt.Printf("\nSolution found! Enter this in the crackme:\n%s\n", string(solution))
}
