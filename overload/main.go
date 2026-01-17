package main

import (
	"fmt"
	"math"
)

// 1. DATA STRUCTURES (No Classes)
// Go uses structs to hold data.
type Circle struct {
    Radius float64
}

type Rectangle struct {
    Width, Height float64
}

// 2. METHODS (The "Receiver" Syntax) 
// Instead of overloading, we attach specific methods to specific types.
// You cannot have two Area() functions in the same package, 
// UNLESS they are methods on different types. 

func (c Circle) Area() float64 {
    return math.Pi * math.Pow(c.Radius, 2)
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// 3. THE "NO OVERLOADING" WORKAROUND
// In TS/Java you might have: Area(shape) and Area(shape, precision).
// In Go, you just give them descriptive names.
func (c Circle) AreaWithPrecision(precision int) string {
    return fmt.Sprintf("%.*f", precision, c.Area())
}

// 4. INTERFACES (Duck Typing)
// This is Go's superpower. Any type that has an Area() method 
// automatically "is a" Shape. We don't say "Rectangle implements Shape" 
type Shape interface {
    Area() float64
}

// 5. POLYMORPHISM 
// This function takes ANY Shape. This is how we handle multiple types
// without needing overloaded versions of the PrintArea function.
func PrintArea(s Shape) {
    fmt.Printf("The area of this shape is %.2f\n", s.Area())
}

// 6. COMPOSITION (Instead of Inheritance)
// A 'Square' isn't a child of Rectangle; it "embeds" a Rectangle
type Square struct {
    Rectangle // This is called "Embedding"
}

func NewSquare(side float64) Square {
    return Square{Rectangle{Width: side, Height: side}}
}

func main() {
    c := Circle{Radius: 5} 
    r := Rectangle{Width: 10, Height: 5}
    s := NewSquare(4)

    // Demonstrating "Overloading" via Methods
    fmt.Println("Circle Area", c.Area())
    fmt.Println("Rect Area:", r.Area())

    // Demonstrating Composition
    // Square gets the Area() method because it embedded Rectangle 
    fmt.Println("Square Area:", s.Area())

    // Demonstrating Polymorphism via Interface 
    shapes := []Shape{c, r, s}
    for _, shape := range shapes {
        PrintArea(shape)
    }

    // Demonstrating "Overload" workaround by name 
    fmt.Println("Precise Circle:", c.AreaWithPrecision(4))
}

