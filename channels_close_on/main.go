package main

import (
	"fmt"
	"sync"
)

func main() {
    ch := make(chan string) 
    var wg sync.WaitGroup

    wg.Go(func()  {
       ch <- "Task 1 complete" 
       ch <- "Task 2 complete" 
    })
    
    // THE FIX: We launch a "Closer" goroutine
    go func() {
        wg.Wait() // Wait for all workers to call wg.Done() 
        close(ch) // Now main's 'range' loop will finish instead of deadlocking 
        fmt.Println("Closer: All workers done, channel closed.")
    }()

    fmt.Println("Main: Waiting for results...")

    // This loop now exists cleanly when the channel is closed! 
    for msg := range ch {
        fmt.Println("Received:", msg)
    }
    fmt.Println("Main: Successfully finished without deadlocking!")
}
