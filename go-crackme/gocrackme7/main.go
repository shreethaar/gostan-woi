package main
import(
    "fmt"
    "bufio"
    "os"
    "strings"
)

var balance = 0
const flagPrice=100

func secret() {
    part1 := []byte{
        'R', '1', 'e', 'f',
        'l', '0', 'Y', 'd',
        'E', '0',
    }
    
    part2 := []byte{
        'G', 'V', 'U', '8',
        'X', 'x', 'f', 'g',
        'Z',
    }
    
    part3 := []byte{
        'Q', '1', 'f', '2',
        'N', 'G', 'd', 't',
        'M', '9',
    }
    result := make([]byte, 0)
    indices := [][2]int{
        {2, 0}, {0, 1}, {0, 0}, {1, 0},  
        {0, 2}, {2, 1}, {1, 1}, {1, 2},  
        {0, 4}, {1, 3}, {0, 5}, {2, 3},  
        {0, 6}, {2, 4}, {1, 5}, {1, 6},  
        {0, 7}, {1, 7}, {2, 7}, {1, 8},  
        {0, 8}, {2, 8}, {0, 9}, {2, 9},  
    }
    
    for _, idx := range indices {
        arrayNum := idx[0]
        pos := idx[1]
        
        switch arrayNum {
        case 0:
            result = append(result, part1[pos])
        case 1:
            result = append(result, part2[pos])
        case 2:
            result = append(result, part3[pos])
        }
    }
    
    fmt.Println(string(result))
}

func purchaseFlag() {
    if balance >= flagPrice {
        balance -= flagPrice
        fmt.Println("You purchased the flag!")
        secret()
    } else {
        fmt.Printf("Insufficient balance! You need %d more coins.\n", flagPrice-balance)
    }
}


func sellFlag() {
    fmt.Println("You don't have flag to sell")
}

func showFlag() {
    fmt.Println("CTF{f4ke_f14gg}")
}

func main() {
    consoleReader:=bufio.NewReader(os.Stdin)
    for {
        fmt.Println("\n======= Flag Store =======")
        fmt.Println("Your current balance:", balance)
        fmt.Println("Choose your operation:")
        fmt.Println("[1] Purchase flag")
        fmt.Println("[2] Sell flag")
        fmt.Println("[3] Show flag")
        fmt.Println("[4] Exit store")
        fmt.Println("Enter choice: ")
        input,_ := consoleReader.ReadString('\n')
        input=strings.TrimSpace(input)

        switch input {
        case "1":
            purchaseFlag()
        case "2":
            sellFlag()
        case "3":
            showFlag()
        case "4":
            fmt.Println("Exiting store. Goodbye!")
            return
        default:
            fmt.Println("Invalid choice!")
        }
    }
}

    
