package main

import "fmt"

// 1. A function that takes 'any' (TypeScript: any)
func DescribeAnything(val any) {
	// We can use a 'type switch' to find out what is inside the box
	switch v := val.(type) {
	case int:
		fmt.Printf("This is an integer: %d\n", v)
	case string:
		fmt.Printf("This is a string: %q\n", v)
	case bool:
		fmt.Printf("This is a boolean: %v\n", v)
	default:
		fmt.Printf("I don't know this type: %T\n", v)
	}
}

func main() {
	// You can pass anything to 'any'
	DescribeAnything(42)
	DescribeAnything("Hello Go")
	DescribeAnything(true)
	DescribeAnything([]int{1, 2, 3})
	// You can even make a slice of any
	mixedList := []any{100, "Banana", 3.14}
	fmt.Println("Mixed List:", mixedList)
}
