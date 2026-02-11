package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"

	"golang.org/x/sync/errgroup"
)

// 1. Define custom types for Context keys to avoid collisions
type ctxKey string

const TraceIDKey ctxKey = "trace-id"

// Result represents data fetched from an "API"
type Result struct {
	Source string
	Data   string
}

func main() {
    // 2. ROOT CONTEXT: Start with Background and add a Trace ID
    rootCtx := context.WithValue(context.Background(), TraceIDKey, "REQ-123")

    // 3. TIMEOUT: Ensure the whole operation doesn't take more than 1 second 
    ctx, cancel := context.WithTimeout(rootCtx, 1*time.Second)
    defer cancel()

    // 4. CHANNELS: Create a buffered channel to collect results 
    // We make it size 3 because we have 3 workers 
    workersNum := 3
    resultsChan := make(chan Result, workersNum)

    // 5. ERRGROUP: Manage multiple goroutines and short-circuiting 
    g, gCtx := errgroup.WithContext(ctx)

    sources := []string{"Web", "Images", "Video"}

    for _, source := range sources {
        // Create a local copy for the closure 
        s := source 

        // 6. GOROUTINES: Launch each search in its own "branch"
        g.Go(func() error {
            return fakeSearchAPI(gCtx, s, resultsChan)
        })
    }

    // 7. SYNC: Wait for all workers to finish OR for one to fail
    if err := g.Wait(); err != nil {
        fmt.Printf("Search failed: %v\n", err)
    }

    // 8. CLEANUP: Close the channel so the range loop below knows when to stop 
    close(resultsChan)

    // 9. SELECT/RANGE: Print all results we managed to get 
    fmt.Println("--- Search Results ---")
    for res := range resultsChan {
        fmt.Printf("[%s]: %s\n", res.Source, res.Data)
    }
}

func fakeSearchAPI(ctx context.Context, source string, ch chan<- Result) error {
    // Retrieve value from context 
    traceID := ctx.Value(TraceIDKey)

    // Simulate random work time 
    wait := time.Duration(rand.IntN(1500)) * time.Millisecond

    fmt.Printf("Worker %s (Trace: %v)) starting (will take %v)...\n", source, traceID, wait)

    select {
    case <- time.After(wait): 
        // Simulate a failure in "Video" search to trigger short-circuit 
        if source == "Video" && wait > 800 * time.Millisecond {
            return fmt.Errorf("video server exploded")
        }
        // SUCCESS: Send data into the channel 
        ch <- Result{Source: source, Data: "Successful search results"}
        return nil

    case <- ctx.Done(): 
        // SHORT-CIRCUIT: The context was canceled (timeout or sibling error)
        fmt.Printf("Worker %s: Shutting down early: %v\n", source, ctx.Err())
        return ctx.Err() 
    }
}
