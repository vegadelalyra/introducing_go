package main

import "fmt"

func main() {
    allocate() // Initialize the map allocating a hash map in memory 
    declare()
}

// This func is doomed to crash on runtime since go cannot write on not initialized maps
func declare() {
    // 1. Set up the recovery mechanism
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered from panic: %v\n", r)
        }
    }()
    var x map[string]int 
    x["key"] = 10 
    fmt.Println(x)
}

func allocate() {
    x := make(map[string]int)
    x["key"] = 10
    x["key2"] = 20
    fmt.Println(x)
    fmt.Println(x["miau"])
}


