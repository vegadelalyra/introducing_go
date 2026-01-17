package main

import (
	"fmt"
	"math/rand"
	"time"
)

const arrLen = 6

func main() {
	desiredElm := 4
    s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := make([]int, 3, 9)
    x := [arrLen]string{"a", "b", "c", "d", "e", "f"}
    wantedStart := 2
    wantedEnd := 5

    defer func() {
        if ok := recover(); ok != nil {
            fmt.Printf("Recovered from panic: %v\n", ok)
        }
    }()
        
	fmt.Printf(
		"The element %d of the slice %v is... %d",
		desiredElm,
		s,
		accessToElm(s, desiredElm),
    )
	fmt.Println()  
	fmt.Println()  

	printSliceStats(slice)
	fmt.Println()  

    fmt.Printf(
        "For the array x = %v, \nslicing x[2:5] is %v", 
        x, 
        givenSliceRange(&x, wantedStart, wantedEnd),
    )
	fmt.Println()  
	fmt.Println()  

    ranSlice := generateRandomSlice(5,25) 
	fmt.Printf(
        "The smallest num in %v\n is %d",
        ranSlice, 
        smallestNum(ranSlice),
    )
}


// generateRandomSlice creates a slice of random length and populates it
// with random integers between 0 and 100.
func generateRandomSlice(minLen, maxLen int) []int {
    // Seed the random number generator using current time  
    source := rand.NewSource(time.Now().UnixNano())
    r := rand.New(source)

    length := r.Intn(maxLen-minLen+1) + minLen
    slice := make([]int, length)

    for i := range length {
        slice[i] = r.Intn(101) // Values 0 to 100
    }
    return slice 
}


func smallestNum(s []int) int {
    min := s[0]

    for _, value := range s {
        if value < min {
            min = value 
        }
    }

    return min 
}

func givenSliceRange[T comparable](a *[arrLen]T, start, end int) []T {
    return a[start:end] 
}

func printSliceStats[T any](s []T) {
	fmt.Printf("Your slice of %T has len: %d an d cap: %d", s, len(s), cap(s))
	fmt.Println()
}

func accessToElm[T comparable](s []T, i int) T {
	return s[i-1]
}
