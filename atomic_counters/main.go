package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var (
		wg          sync.WaitGroup
		mutexCount  int64
		atomicCount int64
		mu          sync.Mutex
	)

    const iterations = 1_000_000

    // --- TEST 1: MUTEX --- 
    startMu := time.Now() 
    for range iterations {
        wg.Go(func () {
            mu.Lock()
            mutexCount++
            mu.Unlock()
        })
    }
    wg.Wait()
    fmt.Printf("Mutex Time: %v | Value: %d\n", time.Since(startMu), mutexCount)

    // --- TEST 2: ATOMIC --- 
    // No locking. The CPU hardware handles the sync. 
    startAt := time.Now()
    for range iterations {
        wg.Go(func() {
            atomic.AddInt64(&atomicCount, 1)
        })
    }
    wg.Wait()
    fmt.Printf("Atomic Time: %v | Value: %d\n", time.Since(startAt), atomicCount)
}
