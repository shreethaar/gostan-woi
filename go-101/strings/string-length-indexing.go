package main

import ("fmt")

func main() {
	str:="Hello"
	length:=len(str) 
	firstChar:=str[0]

	fmt.Println(str);
	fmt.Printf("Length of str is: %d\n", length);
	fmt.Printf("First character is: %c",firstChar);

}

