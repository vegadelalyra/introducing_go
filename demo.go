package main

import "fmt"

type Page int
const (
    demo  Page = iota
    danger
    angel
    fix
    focus 
) 

func main() {
    fmt.Println(demoAnything())
}

func demoAnything() []Page {
    return []Page{demo, danger,angel, fix} 
}


