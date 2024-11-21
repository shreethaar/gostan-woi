package main
import (
    "fmt"

)

/* 
1. Input car details
2. Show package price and details
3. Show car performance before and after tuning result

TO FIX:
- exception handling
- looping until cancel is prompted


*/

func carDetails(cBrand string, cModel string, topSpeed float64, horsePower float64, torque float64) {
    fmt.Println("Car Details:")
    fmt.Printf("Brand: %s\n", cBrand)
    fmt.Printf("Model: %s\n", cModel)
    fmt.Printf("Top Speed: %s\n", topSpeed)
    fmt.Printf("Horsepower: %s\n", horsePower)
    fmt.Printf("Torque: %s\n", torque)
}

func main() {
    fmt.Println("=======Tuning Workshop========")
    /*
    Car details
    1. Car brand
    2. Car model
    3. Initial top speed, horsepower and torque
    */
    fmt.Println("Enter the car details")
    var cBrand,cModel string
    var topSpeed,horsePower, torque float64

    fmt.Print("Car brand: ")
    fmt.Scanln(&cBrand)
    fmt.Print("Car Model: ")
    fmt.Scanln(&cModel)
    fmt.Print("Top Speed (km/h): ")
    fmt.Scanln(&topSpeed)
    fmt.Print("Horsepower (HP): ")
    fmt.Scanln(&horsePower)
    fmt.Print("Torque (Nm): ")
    fmt.Scanln(&torque)

    carDetails(cBrand,cModel,topSpeed,horsePower,torque)

}
