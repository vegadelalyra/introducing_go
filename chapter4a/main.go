package main

import "fmt"

func main() {
	a()
    b()
    c()
    d()
}



func b() {
    fmt.Println(`1
    2
    3
    4
    5
    6
    7
    8
    9
    10`)
}
func c() {
    i := 1 
    for i <= 10 {
        fmt.Println(i)
        i = i + 1 
    }
}

func d() {
    for i:= 1; i <= 10; i++ {
        fmt.Println(i)
    }
}

func a() {
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
	fmt.Println(6)
	fmt.Println(7)
	fmt.Println(8)
	fmt.Println(9)
	fmt.Println(10)
}

