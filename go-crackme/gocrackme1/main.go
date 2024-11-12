package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Println("Welcome to GoCrackMe Challenge!")
    fmt.Println("Can you figure out the correct password?")
    fmt.Print("Enter password:")
    var input string
    fmt.Scanln(&input)
    if checkPassword(input){
        fmt.Println("\n🎉 Congratulations! You've solved the challenge!")
		fmt.Println("Flag: CTF{y0u_f0und_th3_h1dden_tr34sure}")
	} else {
		fmt.Println("\n❌ Incorrect password. Try again!")
		os.Exit(1)
	}
}

func checkPassword(input string) bool {
    if len(input) != 12 {
		return false
	}

	hasNumber := false
	hasSpecial := false
	hasUpper := false
	hasLower := false

	for _, char := range input {
		switch {
		case char >= '0' && char <= '9':
			hasNumber = true
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= 'a' && char <= 'z':
			hasLower = true
		case strings.ContainsRune("!@#$%^&*", char):
			hasSpecial = true
		}
	}

	if !hasNumber || !hasSpecial || !hasUpper || !hasLower {
		return false
	}

    hasher := sha256.New()
    hasher.Write([]byte(input))
    hash := hex.EncodeToString(hasher.Sum(nil))
    targetHash:= "8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918"
    return hash==targetHash
}
