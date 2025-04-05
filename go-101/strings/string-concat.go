package main


import(	
	"fmt"
	"strings"
)

func main (
	// using + operator
	firstName:="John"
	lastName:="Doe"
	fullName:=firstName+" "+lastName // "John Doe"

	// using += operator
	greeting:="Hello"
	greeting+=", World!" // "Hello, World!"

	// using strings.Join (more efficient for multiple strings)
	parts:=[]strings{"Go","is","awesome"}
	sentence:=strings.Join(parts," ") // "Go is awesome"

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(fullName)

	fmt.Println(greeting)
	fmt.Println(sentence)
)
