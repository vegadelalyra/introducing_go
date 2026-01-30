package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1. Create a Reader that listens to Standard Input (Keyboard)
	// bufio.NewReader provides more control than a simple Scan
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("--- Advanced Input with bufio ---")

	// 2. Reading a full sentence (The "Space" Problem Solved)
	fmt.Print("Enter your full name (including spaces): ")

	// ReadString will keep reading until it hits the '\n' (Enter key)
	fullName, _ := reader.ReadString('\n')

	// IMPORTANT: ReadString includes the '\n' at the end. We must trim it.
	fullName = strings.TrimSpace(fullName)

	// 3. The Validation Loop with bufio
	var bio string
	for {
		fmt.Print("Describe yourself in at least 10 characters: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if len(input) < 10 {
			fmt.Println("Error: Too short! Try again.")
			continue
		}
        bio = input 
        break
	}

    // 4. Using bufio.Scanner (Better for word-by-word or simple line reading) 
    fmt.Println("\nEnter 3 favorite hobbies (one per line):")
    scanner := bufio.NewScanner(os.Stdin)

    hobbies := []string{}
    for i := range 3 {
        fmt.Printf("Hobby %d: ", i+1)
        if scanner.Scan() {
            hobbies = append(hobbies, scanner.Text())
        }
    }

    // Final Output 
    fmt.Println("\n--- User Profile ---")
    fmt.Printf("Name: %s\n", fullName)
    fmt.Printf("Bio: %s\n", bio)
    fmt.Printf("Hobbies: %v\n", strings.Join(hobbies, ", "))
}
