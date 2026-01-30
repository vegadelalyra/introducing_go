package main

import (
	"fmt"
	"unicode"
)

func main() {
	var name string
	var age int
	var score float64

	fmt.Println("--- Welcome to the Go Scanner Exercise ---")

	// 1. Simple Scan (stops at space)
	for {
		fmt.Print("Enter your first name: ")
		// We pass the address (&name) so the function can modify the variable
		_, err := fmt.Scan(&name)
		if err != nil || !isGlobalAlpha(name) {
			invalidInput()
			continue
		}
		break
	}

	// 2. Scan multiple values at once
	for {
		fmt.Print("Enter your age and your favorite decimal (e.g., 25 3.14): ")
		_, err := fmt.Scan(&age, &score)
		if err != nil {
			fmt.Println("\nError reading input. Make sure you entered a number.")
			continue
		}
		break
	}

	fmt.Printf(
		"\nHello %s! You are %d years old and your score is %.2f.\n",
		name,
		age,
		score,
	)

	// --- THE GOTCHA: Handling Spaces ---
	// fmt.Scan cannot read a full sentence because it stops at the first space.
	// For full sentences, we usually use the 'bufio' package,
	// but here is how you do it with fmt.Scanln:
	var city, country string
	for {
		fmt.Print("\nEnter your City and Country (separated by space): ")
        _, err := fmt.Scanln(&city, &country)
        if err != nil {
            invalidInput()
            continue 
        }
        break
	}
	fmt.Printf("You live in %s, %s.\n", city, country)
    var discard string 
    fmt.Scanln(&discard)
}

func isAlpha(s string) bool {
	for _, char := range s {
		// If the character is outside 'a'-'z' and 'A'-'Z'
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') {
			return false
		}
	}
	return len(s) > 0
}

func isGlobalAlpha(s string) bool {
	for _, char := range s {
		// This checks for any letter in any language (Greek, Cyrillic, etc.)
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return len(s) > 0
}

func invalidInput() {
	fmt.Println("❌ Invalid input. Try again.")
}
