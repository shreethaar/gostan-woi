package main
import ("fmt")

func main() {
    for i:=0;i<5;i++ {
        fmt.Println(i)
    }
    fmt.Println()
    
    for i:=0;i<100;i+=10 {
        fmt.Println(i)
    }
}
