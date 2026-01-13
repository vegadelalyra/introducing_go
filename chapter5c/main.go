 package main

 import "fmt" 

 func main() {
    fmt.Println("append:")
    appendSlices()
    fmt.Println("copy:")
    copySlices()
 }

 func appendSlices() {
     slice1 := []int{1,2,3}
     slice2 := append(slice1, 4, 5)
     fmt.Println(slice1, slice2)
 }

 func copySlices() {
    slice1 := []int{1,2,3}
    slice2 := make([]int, 2)
    copy(slice2, slice1)
    fmt.Println(slice1, slice2)
 }
