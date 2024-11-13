package main

import ("fmt")

func main() {
    time := 20
    if(time<18) {
        fmt.Println("Good day")
    } else {
        fmt.Println("Good evening")
    }

    temperature := 14
    if(temperature>15) {
        fmt.Println("It is warm out here")
    } else {
        fmt.Println("It is cold out there")
    }
}

