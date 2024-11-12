package main
import ("fmt")

func main() {
    var x int = 500
    var y int = -4500
    fmt.Printf("Type: %T, value: %v",x,x)
    fmt.Printf("Type: %T, value: %v",y,y)

    //int -> depends on platform: 32bits in 32 systems, 64bit in 64 systems
    //int8 -> 8bits/1byte
    //int16 -> 16bits/2byte
    //int32 -> 32bits/4byte
    //int64 -> 64bits/8byte


    var a uint = 500
    var b uint = 4500
    fmt.Println()
    fmt.Printf("Type: %T, value: %v",a,a)
    fmt.Printf("Type: %T, value: %v",b,b)

    
}

