package main

import "fmt"

func main() {
    slice := []string{"apple", "banana", "apple", "orange", "banana", "kiwi"}

    result := []string{}
    seen := map[string]bool{}     

    for _, str := range slice {
        if !seen[str] {           
            result = append(result, str) 
            seen[str] = true      
        }
    }

    fmt.Println(result)
}
