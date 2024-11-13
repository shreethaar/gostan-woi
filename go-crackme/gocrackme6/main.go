
package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    obfuscatedFlag := []byte{0x56, 0x2E, 0x1C, 0x7D, 0xB3, 0xA5, 0xDD, 0x92, 0x4A, 0x33}

    fmt.Print("Enter the key to unlock the flag: ")
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        key := scanner.Text()
        if len(key) == len(obfuscatedFlag) {
            flag := make([]byte, len(obfuscatedFlag))
            for i := range obfuscatedFlag {
                flag[i] = (obfuscatedFlag[i] ^ key[i]) & 0xFF
            }

            fmt.Println("Flag:", string(flag))
        } else {
            fmt.Println("Invalid key length!")
        }
    }
}
