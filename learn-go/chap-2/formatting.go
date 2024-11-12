package main
import ("fmt")

func main() {
    var i=15.5
    var txt = "Hello world!"
    
    fmt.Printf("%v\n",i)
    fmt.Printf("%#v\n",i)
    fmt.Printf("%v%%\n",i)
    fmt.Printf("%T\n",i)


    fmt.Printf("%v\n",txt)
    fmt.Printf("%#v\n",txt)
    fmt.Printf("%T\n",txt)

    var j=10
    fmt.Printf("%b\n",j)
    fmt.Printf("%d\n",j)
    fmt.Printf("%+d\n",j)
    fmt.Printf("%o\n",j)
    fmt.Printf("%O\n",j)
    fmt.Printf("%x\n",j)
    fmt.Printf("%X\n",j)
    fmt.Printf("%#x",j)
    fmt.Printf("%4d",j)
    fmt.Printf("%-4d",j)
    fmt.Printf("%04d",j)
        
    var txt2 = "Hello"
    fmt.Printf("%s\n", txt2)
    fmt.Printf("%q\n", txt2)
    fmt.Printf("%8s\n", txt2)
    fmt.Printf("%-8s\n", txt2)
    fmt.Printf("%x\n", txt2)
    fmt.Printf("% x\n", txt2)
    
    var a = true
    var b = false
    fmt.Println(a)
    fmt.Println(b)
    fmt.Printf("%t\n",a)
    fmt.Printf("%t\n",b)
    
    var c = 3.141
    fmt.Printf("%e\n",c)
    fmt.Printf("%f\n",c)
    fmt.Printf("%.2f\n",c)
    fmt.Printf("%6.2f\n",c)
    fmt.Printf("%g\n",c)
}
    
