package main

import (
	"fmt"
	"time"
)

func main() {
	// --- 1. Memory Addresses & Growth ---
	fmt.Println("--- 1.POINTERS AND GROWTH ---")
	s := make([]int, 0, 1)
	for i := range 5 {
		// %p prints the memory address of the underlying array
		fmt.Printf("Len: %d, Cap %d, Pointer: %p\n", len(s), cap(s), s)
		s = append(s, i)
	}
	// The slice pointer change when Cap jumps!

	// --- 2. The "Shared Memory" Trap ---
	fmt.Println("\n--- 2. THE SHARED MEMORY TRAP ---")
	original := []int{10, 20, 30, 40, 50}
	sub := original[0:2] // Len 2, Cap 5 (Shares memory!)

	fmt.Printf("Original before: %v\n", original)
	sub = append(sub, 999) // We appended to 'sub' but...
	fmt.Printf("Original after: %v (Index 2 was overwritten!)\n", original)

	// --- 3. The "Full Slice Expression" Fix [x:y:z] ---
	fmt.Println("\n--- 3. THE [x:y:z] PROTECTOR ---")
	original2 := []int{10, 20, 30, 40, 50}
	// We set Cap to 2. It cannot expand into the original array anymore.
	safeSub := original2[0:2:2]

	fmt.Printf("Original2 before: %v\n", original2)
	safeSub = append(safeSub, 999) // Cap was full, so it move dto a NEW array
	fmt.Printf("Original2 after: %v (Untouched!)\n", original2)
	fmt.Printf("safeSub: %v (Lives in its own memory now)\n", safeSub)

    // --- 4. Speed: Auto-grow vs. Pre-allocate ---
    fmt.Println("\n--- 4. PERFORMANCE BENCHMARK ---")
    iterations := 200_000_000

    // Benchmark Auto-grow 
    start := time.Now()
    auto := []int{}
    for i := range iterations {
       auto = append(auto, i) 
    }
    fmt.Printf("Auto-grow took: %v\n", time.Since(start))

    // Benchmark Pre-allocated
    start = time.Now() 
    pre := make([]int, 0, iterations)
    for i := range iterations {
        pre = append(pre, i)
    }
    fmt.Printf("Pre-allocate took: %v\n", time.Since(start))
}
