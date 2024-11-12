package main

import (
    "fmt"
    "math"
    "os"
    "time"
)

var secretValue float64 = 3.141592653589793
type ValidationError struct {
    Message string
}

func (e *ValidationError) Error() string {
    return e.Message
}

func compareInMemory(a,b []byte) bool {
    if len(a)!=len(b) {
        return false
    }
    diff := byte(0)
	for i := 0; i < len(a); i++ {
		diff |= a[i] ^ b[i]
		time.Sleep(time.Duration(int64(float64(i)*0.00001)) * time.Nanosecond)
	}

	return diff == 0
}


func validateInput(input []byte) error {
	if len(input) != 12 {
		return &ValidationError{"Invalid input length"}
	}
	for i := 0; i < 8; i++ {
		if (input[i] < '0' || input[i] > '9') && (input[i] != '0' && input[i] != '1') {
			return &ValidationError{"Invalid character"}
		}
	}

	_, err := fmt.Sscanf(string(input[8:]), "%f", new(float64))
	if err != nil {
		return &ValidationError{"Invalid decimal"}
	}

	return nil
}

func transformInput(input []byte) []byte {
	result := make([]byte, len(input))
	copy(result, input)
	for i := 0; i < 8; i++ {
		if input[i] == '1' {
			result[i] = byte(^uint8(i) & 0xFF)
		} else {
			result[i] = byte(uint8(i) | 0x80)
		}
	}
	decimalValue, _ := fmt.Sscanf(string(input[8:]), "%f", new(float64))
	result[8] = byte(math.Floor(decimalValue * 100) % 256)
	result[9] = byte(math.Floor(decimalValue * 10000) % 256)
	result[10] = byte(math.Floor(decimalValue * 1000000) % 256)
	result[11] = byte(math.Floor(decimalValue * 100000000) % 256)

	return result
}

func main() {
	fmt.Println("=== Boolean and Decimal Crackme ===")
	fmt.Println("Enter the 12-character key: ")

	var input string
	fmt.Scanln(&input)
	if err := validateInput([]byte(input)); err != nil {
		fmt.Println("❌", err)
		os.Exit(1)
	}
	transformed := transformInput([]byte(input))
	expected := []byte{
		0xFE, 0xCF, 0xAB, 0x87, 0x65, 0x43, 0x21, 0x00,
		0x49, 0x53, 0x23, 0x71,
	}
	if compareInMemory(transformed, expected) {
		fmt.Printf("\n🎉 Congratulations! You've solved it!\n")
		fmt.Printf("Flag: CTF{%s}\n", input)
	} else {
		fmt.Println("\n❌ Incorrect key")
	}
}
