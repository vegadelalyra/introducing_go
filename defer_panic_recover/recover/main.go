package main

import "fmt"

func main() {
	recover_map()
}

func recover_map() {
	// 1. Set up the recovery mechanism
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()
	var x map[string]int
	x["key"] = 10
	fmt.Println(x)
}
