package main

import (
	"fmt"
	"slices"
)

func main () {
    fruits := []string{"apple", "banana", "cherry"}
    fmt.Println(fruits)
    
    reverseInPlace(fruits)
    fmt.Println(fruits)

    slices.Reverse(fruits) 
    fmt.Println(fruits)

    reverseMidPoint(fruits)
    fmt.Println(fruits)
}

func reverseMidPoint(s []string) {
    n := len(s) 
    for i := 0; i < n/2; i++ {
        s[i], s[n-1-i] = s[n-1-i], s[i]
    }
}

func reverseInPlace(s []string) {
    for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
       s[i], s[j] = s[j], s[i] 
    }
}


