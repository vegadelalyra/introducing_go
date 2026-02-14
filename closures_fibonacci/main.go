package main

import "fmt"

type IntClosure func() int 

const Sequences int = 10

func fibonacci() IntClosure {
    var x, y = 0, 1 

    return func() int {
        var fibo int = x 
        x, y = y, x + y
        return fibo  
    } 
} 

func main() {
    f := fibonacci()
    for range Sequences {
       fmt.Println(f())
    }
}

