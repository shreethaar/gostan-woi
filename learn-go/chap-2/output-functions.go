package main
import ("fmt")

func main() {
    var i, j string = "Hello","World"
    fmt.Print(i)
    fmt.Print(j)

    fmt.Print(i,"\n")
    fmt.Print(j,"\n")
    fmt.Print(i,"\n",j)
    fmt.Print(i," ",j)
    
    var x,y int = 10,20
    fmt.Print(x,y)

    var a string = "Hello"
    var b int = 15
    fmt.Printf("i has value: %v and type: %T\n",a,a)
    fmt.Printf("j has value: %v and type: %T\n",b,b)

    fmt.Println("line one")
    fmt.Println("line two")
    fmt.Print("\n"+"line three")
    
    
}

