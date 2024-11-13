package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Calculator")
	fmt.Println("-----------------")

	fmt.Print("Enter first number: ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)
	num1, err := strconv.ParseFloat(input1, 64)
	if err != nil {
		fmt.Println("Invalid number:", err)
		return
	}

    fmt.Print("Enter operator (+, -, *, /): ")
	operator, _ := reader.ReadString('\n')
	operator = strings.TrimSpace(operator)

	fmt.Print("Enter second number: ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)
	num2, err := strconv.ParseFloat(input2, 64)
	if err != nil {
		fmt.Println("Invalid number:", err)
		return
	}

    var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Division by zero")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Invalid operator")
		return
	}

	fmt.Printf("Result: %.2f\n", result)
}
