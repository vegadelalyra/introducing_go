package main

import "fmt"

func main() {
    correctPanicRevocer()
    incorrectPanicRevocer()
}

func incorrectPanicRevocer() {
    panic("PANIC")
    str := recover() // this will never happen 
    fmt.Println(str)
}

func correctPanicRevocer() {
    defer func() {
        str := recover()
        fmt.Println(str)
    }()
    panic("PANIC")
}


