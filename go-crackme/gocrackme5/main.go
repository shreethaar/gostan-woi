
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    secretPassword := "CTF{R3v3rs3_Th1s_G0}"
    fmt.Print("Enter password: ")
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        input := scanner.Text()
        if input == secretPassword {
            fmt.Println("Access Granted!")
        } else {
            fmt.Println("Access Denied!")
        }
    }
}
