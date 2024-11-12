package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

const password = "CTF{go_crack_but_not_that_kind_of_crack_:)}"
func main() {
    reader:=bufio.NewReader(os.Stdin)
    fmt.Print("Enter the password:")
    input, _ :=reader.ReadString('\n')
    input=strings.TrimSpace(input)

    if checkPass(input) {
        fmt.Println("Nah! You definitely on crack, congrats btw") 
    } else {
        fmt.Println("Homie, u aint cracking enuf") }

}

func checkPass(input string) bool {
    return input==password
}
