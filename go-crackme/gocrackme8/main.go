package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println("Enter your license key:")
    var key string
    fmt.Scanln(&key)

    if validateKey(key) {
        fmt.Println("Correct! Q1RGe0wxQzNOc0VfaDRDa2VyfQ==")
    } else {
        fmt.Println("Invalid key. Try again.")
    }
}

func validateKey(key string) bool {
    if !strings.Contains(key,"-") || len(key)!=19 {
        return false
    }
    parts:=strings.Split(key,"-")
    if len(parts)!=4 {
        return false
    }

    magic := []int{0x48,0x65,0x6C,0x6F}
    for i,part := range parts {
        if !checkPart(part,magic[i]) {
            return false
        }
    }

    return true
}

func checkPart(part string, magic int) bool {
    if len(part)!=4 {
        return false
    }

    sum:=0
    for _,ch := range part {
        sum+=int(ch)
    }

    return (sum & 0xFF) == magic
}

