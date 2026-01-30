package main

import (
	"fmt"
	"reflect"
)

type User struct {
    ID int 
    Name string 
    Email string 
}

func main() {
    // 1 .Start with a concrete object 
    u := User{ID: 1, Name: "Alice", Email: "alice@example.com"}

    fmt.Println("--- Step 1: Concrete Object ---")
    fmt.Printf("Value: %v, Type: %T\n\n", u,u)

    // 2. Use reflect.ValueOf() to enter the reflection world 
    // This gives us a  reflect.Value object, which has methods to
    // inspect the structure of the data without knowing its type beforehand.
    v := reflect.ValueOf(u)

    fmt.Println("--- Step 2: reflect.ValueOf(u) ---")
    fmt.Printf("Reflect Value: %v\n", v)
    fmt.Printf("Reflect Kind: %s\n\n", v.Kind()) // e.g., "struct"

    // 3. Use v.Interface() to come back to the concrete world 
    // Interface() packs the value back into an empty interface (interface{})
    i := v.Interface()

    fmt.Println("--- Step 3: v.Interface() ---")
    fmt.Printf("Interface Value: %v, Type: %T\n\n", i, i)

    // 4. To use it as the original type, you must use a Type Assertion
    // Because i is currently interface{}, we can't access .Name yet.
    if finalUser, ok := i.(User); ok {
        fmt.Println("--- Step 4: Type Assertion ---")
        fmt.Printf("Back to User: %s (ID: %d)\n", finalUser.Name, finalUser.ID)
    }
}
