/*
Write a program which can compute the factorial of a given numbers. The results should be printed in a comma-separated sequence on a single line.

Suppose the following input is supplied to the program: 8
Then, the output should be: 40320
*/
package main
import (
    "fmt"
    "log"
)

func main() {
    fmt.Println("Exercise 002")
    var input int
    fmt.Println("Please enter a number: ")
    _, err:=fmt.Scanln(&input)
    if err!=nil {
        log.Fatal("Please enter a number")
    }
    result,err:=Ex002(input)
    if err!=nil{
        log.Fatalf("Error for input: %v: %v",input,err)
    }
    fmt.Printf("Factorial of %d=%d",input,result)
}

func Ex002(input int) (uint64,error) {
    var factorial uint64 = 1
    if input<0{
        return 0, fmt.Errorf("factorial for negativ values is not allowed")
    }
    if input==0 {
        return 1, nil
    }
    for i:=1;i<=input;i++{
        factorial*=uint64(i)
    }
    return factorial,nil
}
