package main

import "fmt"

func adder() func(int) int {
    var sum int
        
    return func(x int) int {
        sum += x 
        return sum
    } 
}

func main() {
    pos, neg := adder(), adder()
    for i := range 10 {
        fmt.Println(
            pos(i), 
            neg(-2*i),
        )
    }
}
