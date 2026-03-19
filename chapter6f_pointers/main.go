package main

import "fmt"

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // x is still 5 

    ziro(&x) 
	fmt.Println(x) // x is 0
}

func ziro(xPtr *int) {
    *xPtr = 0 
}

func zero(x int) {
    x = 0
}

