package main

import (
    "fmt"
    "os"
    "strings"
)

func validateSerial(serial string) bool {
    if len(serial) != 10 {
        return false
    }
    serial = strings.ToUpper(serial)

    if serial[0]!=serial[9]-3 {
        return false
    }

    if serial[1]!=serial[8]+14 {
        return false
    }

    if serial[2]!=serial[7]-20 {
        return false
    }

    if serial[3]!=serial[6]+6 {
        return false
    }

    if (serial[4]+serial[5])/2!=serial[0] {
        return false
    }

    return true

}


func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Usage: %s <serial>\n", os.Args[0])
		return
	}

	serial := os.Args[1]
	if validateSerial(serial) {
		fmt.Println("Good Serial!")
	} else {
		fmt.Println("Bad Serial!")
	}
}
