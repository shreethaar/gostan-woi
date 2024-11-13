package main
import ("fmt")

func main() {
    var a = 15+25
    fmt.Println(a)

    sum1:=100+50
    sum2:=sum1+250
    sum3:=sum2+sum2

    fmt.Println(sum1)
    fmt.Println(sum2)
    fmt.Println(sum3)

    var b = 5
    fmt.Println(b)
    b ++
    fmt.Println(b)
    b --
    fmt.Println(b)
    
    var x = 10
    x+=5
    fmt.Println(x)

    x=5
    var y=3
    fmt.Println(x>y)

}

