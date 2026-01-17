package main

import "fmt"

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Height * r.Width
}

func (r Rectangle) Describe() {
    fmt.Printf("I am %.1fx%.1f\n", r.Width, r.Height)
}

// --- THE EMBEDDING ---
type Square struct {
    Rectangle // <--- This is an "Anonymous Field" or "Embedded Field" 
    Side float64 
}

func NewSquare(s float64) Square {
    // We initialize the embedded Rectangle inside the Square 
    return Square{
        Rectangle: Rectangle{Width: s, Height: s},
        Side:      s, 
    }
}

// --- METHOD OVERRIDING (Shadowing) --- 

// If Square defines Describe(), it "shadows" the Rectangle's Describe() 
func (s Square) Describe() {
    fmt.Printf("I am a Square with side %.1f\n", s.Side)
}

func main() {
    s := NewSquare(5)

    // 1. Promotion: We can call Area() directly on 's' even though
    // Area() is defined on Rectangle. It is "promoted" to the Square.
    fmt.Println("Square Area:", s.Area())
    
    // 2. Direct Access: We can still access the inner Rectangle if we want 
    fmt.Println("Inner Width:", s.Rectangle.Width)

    // 3. Shadowing: This call Square's Describe, not Rectangle's. 
    s.Describe()

    // 4. Accessing Shadowed: We can still reach the "hidden" method!
    s.Rectangle.Describe()
}
