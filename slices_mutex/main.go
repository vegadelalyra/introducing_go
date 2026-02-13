package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
    var results []int 
    var wg sync.WaitGroup

    for i := range 1000 {
        wg.Add(1)
        go func(val int) {
            defer wg.Done()

            // Lock, append, unlock. Very low overhead. 
            mu.Lock()
            results = append(results, val)
            mu.Unlock()
        }(i)
    }

    wg.Wait()
    fmt.Printf("Collected %d results\n. RESULTS => %v", len(results), results)
}
