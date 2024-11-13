package main
import ("fmt")

func main() {
    prices := []int{10,20,30}
    fmt.Println(prices[0])
    fmt.Println(prices[2])

    prices[2]=50
    fmt.Println(prices[2])

    myslice1 := []int{1,2,3,4,5}
    fmt.Printf("myslice1=%v\n",myslice1)
    fmt.Printf("length=%d\n",len(myslice1))
    fmt.Printf("capacity=%d\n",cap(myslice1))

    myslice1=append(myslice1,20,21)
    fmt.Printf("myslice1=%v\n",myslice1)
    fmt.Printf("length=%d\n",len(myslice1))
    fmt.Printf("capacity=%d\n",cap(myslice1))

    myslice2:=[]int{1,2,3}
    myslice3:=[]int{4,5,6}
    myslice4:=append(myslice2,myslice3...)
    fmt.Printf("myslice3=%v\n", myslice4)
    fmt.Printf("length=%d\n", len(myslice4))
    fmt.Printf("capacity=%d\n", cap(myslice4))

    numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
    // Original slice
    fmt.Printf("numbers = %v\n", numbers)
    fmt.Printf("length = %d\n", len(numbers))
    fmt.Printf("capacity = %d\n", cap(numbers))

    // Create copy with only needed numbers
    neededNumbers := numbers[:len(numbers)-10]
    numbersCopy := make([]int, len(neededNumbers))
    copy(numbersCopy, neededNumbers)

    fmt.Printf("numbersCopy = %v\n", numbersCopy)
    fmt.Printf("length = %d\n", len(numbersCopy))
    fmt.Printf("capacity = %d\n", cap(numbersCopy))

}
