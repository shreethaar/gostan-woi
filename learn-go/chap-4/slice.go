package main
import ("fmt")

func main() {
    myslice1:=[]int{}
    fmt.Println(len(myslice1))
    fmt.Println(cap(myslice1))
    fmt.Println(myslice1)

    myslice2:=[]string{"Go","Slices","Are","Awesome"}
    fmt.Println(len(myslice2))
    fmt.Println(cap(myslice2))
    fmt.Println(myslice2)
    
    arr1 := [6]int{10, 11, 12, 13, 14,15}
    myslice := arr1[2:4]

    fmt.Printf("myslice = %v\n", myslice)
    fmt.Printf("length = %d\n", len(myslice))
    fmt.Printf("capacity = %d\n", cap(myslice))
    
    myslice3 := make([]int, 5, 10)
    fmt.Printf("myslice1 = %v\n", myslice3)
    fmt.Printf("length = %d\n", len(myslice3))
    fmt.Printf("capacity = %d\n", cap(myslice3))

    // with omitted capacity
    myslice4 := make([]int, 5)
    fmt.Printf("myslice2 = %v\n", myslice4)
    fmt.Printf("length = %d\n", len(myslice4))
    fmt.Printf("capacity = %d\n", cap(myslice4))

}
