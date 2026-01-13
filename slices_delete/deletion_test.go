package main

import "testing"

// Global to prevent compiler optimizations from skipping the work
var Result []int

func BenchmarkInPlaceDeletion(b *testing.B) {
    // Setup OUTSIDE the loop
    nums := make([]int, 1000) 
    
     // Start timing NOW
    for b.Loop() {
        // We simulate a deletion on a sub-slice so we don't 
        // run out of elements during the N iterations
        _ = append(nums[:500], nums[501:]...)
        nums[:cap(nums)][len(nums)-1] = 0
    }
}

func BenchmarkPureDeletion(b *testing.B) {
    nums := make([]int, 1000)
    
     // Start timing NOW
    for b.Loop() {
        pure := make([]int, len(nums)-1)
        copy(pure[:500], nums[:500])
        copy(pure[500:], nums[501:])
        Result = pure
    }
}
