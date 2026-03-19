package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
    const iterations = 500_000

    // --- TEST 1: MUTEX + SLICE ---
	startMutex := time.Now()
    var mu sync.Mutex
    var sliceResults []int
    var wg sync.WaitGroup

    for i := range iterations {
        wg.Add(1)
        go func(v int) {
            defer wg.Done()
            mu.Lock()
            sliceResults = append(sliceResults, v)
            mu.Unlock()
        }(i)
    }
    wg.Wait()
    fmt.Printf("Mutex + Slice: %v (Count: %d)\n", time.Since(startMutex), len(sliceResults))

    // --- TEST 2: CHANNELS --- 
    startChan := time.Now()
    ch := make(chan int, iterations) // Buffered for speed
    var wgChan sync.WaitGroup 
     
    for i := range iterations {
        wgChan.Add(1)
        go func(v int) {
            defer wgChan.Done()
            ch <- v
        }(i)
    }

    // We need a separate goroutine to close the channel when workers finish
    go func() {
        wgChan.Wait()
        close(ch)
    }()

    var chanResults []int 
    for res := range ch {
        chanResults = append(chanResults, res)
    } 
    fmt.Printf("Channels: %v (Count: %d)\n", time.Since(startChan), len(chanResults))
}
