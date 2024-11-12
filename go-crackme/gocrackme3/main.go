package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"runtime"
	"time"
	"unsafe"
)

// Global canary value
var gCanary []byte

// Custom error type to avoid stack traces
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

// Initialize canary value
func init() {
	gCanary = make([]byte, 8)
	_, err := rand.Read(gCanary)
	if err != nil {
		os.Exit(1)
	}
}

// Perform stack canary check
func checkCanary(canary []byte) bool {
	if len(canary) != len(gCanary) {
		return false
	}
	
	result := byte(0)
	// Time-constant comparison
	for i := 0; i < len(canary); i++ {
		result |= canary[i] ^ gCanary[i]
	}
	return result == 0
}

// Custom memory comparison with timing variation protection
func compareInMemory(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	// Convert slices to array pointers to work directly with memory
	aPtr := (*[1 << 30]byte)(unsafe.Pointer(&a[0]))
	bPtr := (*[1 << 30]byte)(unsafe.Pointer(&b[0]))

	diff := byte(0)
	for i := 0; i < len(a); i++ {
		diff |= aPtr[i] ^ bPtr[i]
		// Add random tiny sleep to prevent timing attacks
		if i%2 == 0 {
			time.Sleep(time.Nanosecond)
		}
	}
	
	return diff == 0
}

// Transform input using stack canary
func transformInput(input []byte) []byte {
	result := make([]byte, len(input))
	copy(result, input)

	// Use canary to transform input
	for i := 0; i < len(result); i++ {
		canaryByte := gCanary[i%len(gCanary)]
		result[i] = result[i] ^ canaryByte ^ byte(i*7)
	}

	return result
}

func validateInput(input string) error {
	if len(input) != 16 {
		return &ValidationError{"Invalid input length"}
	}

	// Check for printable ASCII
	for _, ch := range input {
		if ch < 32 || ch > 126 {
			return &ValidationError{"Invalid characters"}
		}
	}

	return nil
}

func main() {
	// Disable GC during critical operations
	runtime.GC()
	runtime.LockOSThread()
	
	fmt.Println("=== Advanced Crackme Challenge ===")
	fmt.Println("Enter the 16-character key: ")

	var input string
	fmt.Scanln(&input)

	// Validate input
	if err := validateInput(input); err != nil {
		fmt.Println("❌", err)
		os.Exit(1)
	}

	// Create local canary
	localCanary := make([]byte, len(gCanary))
	copy(localCanary, gCanary)

	// Transform input
	inputBytes := []byte(input)
	transformed := transformInput(inputBytes)

	// Expected result after transformation
	expected := []byte{
		0x8f, 0x7a, 0x23, 0x45, 
		0x98, 0xab, 0xcd, 0xef,
		0x12, 0x34, 0x56, 0x78,
		0x9a, 0xbc, 0xde, 0xf0,
	}

	// Verify canary hasn't been tampered
	if !checkCanary(localCanary) {
		fmt.Println("❌ Security check failed")
		os.Exit(1)
	}

	// Compare transformed input with expected
	if compareInMemory(transformed, expected) {
		fmt.Printf("\n🎉 Congratulations! You've solved it!\n")
		fmt.Printf("Flag: CTF{%s}\n", hex.EncodeToString(transformed))
	} else {
		fmt.Println("\n❌ Incorrect key")
	}
}
