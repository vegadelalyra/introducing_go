package main

import "fmt"

func main() { 
    var x string 
    x = "first " 
    fmt.Println(x)
    x = x + "second" 
    fmt.Println(x)

    var t string = "hello" 
    var y string = "word" 
    fmt.Println("is t equal to y?", t == y)

    var w string = "hello" 
    fmt.Println("is t equal to w?", t == w)

    z := "Hello, World from z := ! "
    var o = "Hello, World" 

    h := 5 

    fmt.Println(z, o, h)
}
