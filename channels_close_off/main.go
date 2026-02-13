package main

import "fmt"

func main() {
    ch := make(chan string)

    // Worker 
    go func() {
       ch <- "Task 1 complete" 
       ch <- "Task 2 complete" 
       // WE FORGOT TO CALL: close(ch) 
       fmt.Println("Worker: I'm finished, but I didn't close the channel...")
    }()

    // Receiver using range 
    fmt.Println("Main: Waiting for results...")

    // This loop will print the 2 tasks, then FREEZE/DEADLOCK. 
    // It is waiting for a 3rd message or a 'close' signal that never coems. 
    for msg := range ch {
        fmt.Println("Received:", msg)
    }

    // This line will NEVER be reached 
    fmt.Println("Main: All work done!")
}
