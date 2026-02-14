package main

import "fmt"

func main() {
	captureValue()
	closureValue()
}

func closureValue() {
    fmt.Println("\nAccess values through closures")
	message := "Original"
	defer func() {
		fmt.Println("Deferred message:", message)
	}()

	message = "Updated"
	fmt.Println("Current message:", message)
}

func captureValue() {
    fmt.Println("\nCapture value through params")
	message := "Original"
	defer func(msg string) {
		fmt.Println("Deferred message:", msg)
	}(message)

	message = "Updated"
	fmt.Println("Current message:", message)
}
