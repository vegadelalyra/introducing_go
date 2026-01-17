package main

import (
	"fmt"
	"math/rand"
	"slices"
	"time"
)

func main() {
    // 1. Generate the random slice 
    nums := generateRandomSlice(5, 20) // Random length between 5 and 20 
    fmt.Printf("Generated Slice: %v\n\n", nums)

    if len(nums) == 0 {
        fmt.Println("Slice is empty.")
        return 
    }

    // 2. Approach A: The Manual Iterative Way (Most efficient/Standard) 
    minIterative := findMinIterative(nums) 
    fmt.Printf("Min (Iterative) // efficient: %d\n", minIterative)

    // 3. Approach B: The Modern way (Using the 'slices' package)
    // Note: 'slices.Min' was introduced in Go 1.21
    minSlicesPkg := slices.Min(nums)
    fmt.Printf("Min (slices.Min) // fine: %d\n", minSlicesPkg)

    // 4. Approach C: Sort (The worst)
    slices.Sort(nums)
    fmt.Printf("Min (slices.Sort[0]) // the worst: %d\n", nums[0])
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

// findMinIterative is the classic "best" way for custom logic or older Go versions 
// Time Complexity: O(n) | Space Complexity: O(1)
func findMinIterative(x []int) int {
    // Assume the first element is the smallest to start 
    min := x[0]

    for _, value := range x {
        if value < min {
            min = value 
        }
    }
    return min
}

