package main

import("fmt")


func main() {
	// basic string declaration
	var name string = "Hello,Go!"

	// short declaration
	greeting:="Hello,world!"

	// empty string 
	var empty string // default to ""

	// raw string literals (preserves formatting)
	multiline:=`This is a
	multiline string
	that preserves all whitespace`
	

	fmt.Println(name)
	fmt.Println(greeting)
	fmt.Println(empty)
	fmt.Print(multiline)
}

