package main

import "fmt"

type Player struct {
	Name  string
	Level int
}

func main() {
    p := Player{"Aragorn", 99}

    // --- 1. General Verbs ---
    fmt.Println("--- General Verbs ---")
    fmt.Printf("Default format (%%v): %v\n", p) // {Aragorn 99}
    fmt.Printf("Struct with field names (%%+v): %+v\n", p) // {Name:Aragorn Level:99}
    fmt.Printf("Go syntax representation (%%#v): %#v\n", p) // main.Player{Name:"Aragorn", Level:99}
    fmt.Printf("The Type (%%T): %T\n", p) // main.Player

    // --- 2. String & Slice Verbs ---
    fmt.Println("\n--- String Verbs ---")
    s := "Golang"
    fmt.Printf("Plain string (%%s): %s\n", s) // Golang
    fmt.Printf("Double-quoted string (%%q): %q\n", s) // "Golang"

    // --- 3. Integer & Math Verbs ---
    fmt.Println("\n--- Number Verbs ---")
    n := 255
    fmt.Printf("Decimal (%%d): %d\n", n) // 255
    fmt.Printf("Binary (%%b): %b\n", n) // 11111111
    fmt.Printf("Hexadecimal (%%x): %x\n", n) 
   
    f := 3.14159265
    fmt.Printf("Float (%%f): %f\n", f) 
    fmt.Printf("Float 2 decimal places (%%.2f): %.2f\n", f)

    // --- 4. Boolean & Pointers --- 
    fmt.Println("\n--- Misc ---")
    fmt.Printf("Boolean (%%t): %t\n", true)
    fmt.Printf("Pointer address (%%p): %p\n", &p) 

    // --- 5. Sprintf (return as string, don't print) --- 
    message := fmt.Sprintf("Saved string: %s", p.Name)
    fmt.Println("\nResult of Sprintf:", message)
}
