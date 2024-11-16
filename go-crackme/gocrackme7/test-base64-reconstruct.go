package main

import (
    "fmt"
)

func secret() {
    // Split into three shuffled arrays
    part1 := []byte{
        'R', '1', 'e', 'f',
        'U', '0', 'Y', 'd',
        'E', '9',
    }
    
    part2 := []byte{
        'G', 'V', 'l', '8',
        'X', 'x', 'f', 'g',
        'Z', 'M',
    }
    
    part3 := []byte{
        'Q', '1', '{', '2',
        'N', '4', 'd', 't',
        'g', '3',
    }
    
    // Reconstruct in correct order
    result := make([]byte, 0)
    
    // Fixed indices to reconstruct: "Q1RGe1VfUl80X2YxNGdfdEgxM0Z9"
    indices := [][2]int{
        {2, 0}, {2, 1}, {0, 0}, {1, 0},  // Q1RG
        {0, 2}, {2, 2}, {1, 1}, {0, 4},  // e{VU
        {0, 4}, {1, 2}, {0, 5}, {1, 3},  // Ul80
        {1, 4}, {2, 3}, {1, 5}, {2, 4},  // X2xN
        {1, 6}, {2, 5}, {1, 7}, {0, 7},  // f4gd
        {0, 7}, {2, 7}, {2, 8}, {2, 9},  // dtg3
        {1, 9}, {2, 9}, {1, 8}, {0, 9},  // M3Z9
    }
    
    // Build final string using indices
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

func main() {
    secret()
}
