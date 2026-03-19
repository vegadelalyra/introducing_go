package main

import "fmt"



func main() {
    xs := []float64{98,93,77,82,83}
    fmt.Println(average(xs))
    fmt.Println(add(1,2,3))
}

func average(xs []float64) (avg float64) {
    total := 0.0 
    for _, v := range xs {
        total += v 
    }
    avg = total / float64(len(xs))
    return
}

func add(args ...int) (total int) {
    for _, v := range args {
        total += v 
    }
    return 
}

