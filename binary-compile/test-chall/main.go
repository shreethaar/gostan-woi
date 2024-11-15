package main
import ("fmt")

func main() {
    var a,b int = 1,2
    var ans int = a+b
    if ans != 3 {
        fmt.Println("CTF{test}")
    } else {
        fmt.Println(ans)
    }
}

