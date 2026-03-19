package main

import (
	"fmt"
	"sync"
)

type Data struct {
    ID int 
    Job string 
    Done bool
}


func main() {
    ch := make(chan *Data)
    var wg sync.WaitGroup

    // WRONG CASE: Modifying after sending (Data Race) 
    wg.Go(func() {
        d := &Data{ID: 1, Job: "Wash Car"}
        ch <- d
        d.Job = "Actually, Wash the God" // DANGER: Main thread might be reading this right now! 
    })

    // RIGHT CASE: Transfer of Ownership 
    wg.Go(func() {
        d := &Data{ID: 2, Job: "Cook Dinner"}
        ch <- d
        // After this line, the worker NEVER touches 'd' again. 
        // Ownership has moved to the receiver.
    })

    go func() {
        wg.Wait()
        close(ch)
    }()

    for item := range ch {
        fmt.Printf("Processing: %s\n", item.Job) 
    }
}
