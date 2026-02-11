package main

import "fmt"

type Vertex struct {
    X int 
    Y int 
}

func main() {
	i, j := 42, 2701

	p := &i // point to i
	fmt.Println(*p) // read i through the pointer
	fmt.Println(p) // read p (the pointer itself, which means the memory address)
    *p = 21 // set i through the pointer
    fmt.Println(i) // set the new value of i
    
    p = &j // point to j 
    *p = *p / 37 // divide j through the pointer
    fmt.Println(j) // see the new value of j
    fmt.Println(p) // see the new memory address value of p
    fmt.Println(*p) // see the j value through the pointer

    v := Vertex{1, 2}
    b := &v

    b.X = 1e9
    fmt.Println("structs pointers can derefence implicetely:")
    fmt.Println("v: ", v)
}
