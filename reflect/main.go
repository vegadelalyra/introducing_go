package main

import (
	"fmt"
	"reflect"
)

// --------------------------------------------------------------------------
// 1. THE CORE THREE: Type, value and Kind
// --------------------------------------------------------------------------

func demonstrateBasics(i any) {
	t := reflect.TypeOf(i)  // Meta-information (Name, Methods, Fields)
	v := reflect.ValueOf(i) // The actual data (and the ability to change it)
	k := t.Kind()           // The underlying Go primitive (int, struct, map, etc.)

	fmt.Printf("Type: %v | Value: %v | Kind: %v\n", t, v, k)
}

// --------------------------------------------------------------------------
// 2. STRUCT INSPECTION (Tags, Fields, and Methods)
// --------------------------------------------------------------------------

type User struct {
	ID    int    `db:"user_id" json:"id"`
	Email string `db:"user_email" json:"email"`
}

func (u User) SayHello() { fmt.Println("Hello!") }

func inspectStruct(s any) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	fmt.Printf("Inspecting Struct: %s\n", t.Name())

	// Iterate over Fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		tag := field.Tag.Get("db") // Extracting specific struct tags

		fmt.Printf(" Field %d: %s (%v) | Tag: %s | Val: %v\n",
			i, field.Name, field.Type, tag, value)
	}

    // Inspect methods 
    fmt.Printf(" Method Count: %d\n", t.NumMethod())
}

// --------------------------------------------------------------------------
// 3. SETTABILITY (Modifying values at runtime) 
// --------------------------------------------------------------------------

func modifyValue(i any) {
    // To modify a value, we MUST pass a pointer. 
    v := reflect.ValueOf(i)

    // v.Elem() dereferences the pointer to get the actual value
    if v.Kind() == reflect.Pointer && v.Elem().CanSet() {
        val := v.Elem()

        // Logic based on the Kind of the dereferenced element 
        switch val.Kind() {
        case reflect.String: 
            val.SetString("Updated String Content")
        case reflect.Int: 
            val.SetInt(999)
        }
    }
} 
// --------------------------------------------------------------------------
// 4. DYNAMIC SLICE & MAP MANIPULATION
// --------------------------------------------------------------------------

func manipulateCollection() {
    // Create a slice of strings dynamically at runtime 
    sliceType := reflect.TypeFor[[]string]()
    newSlice := reflect.MakeSlice(sliceType, 0, 10)

    // Add values to it
    newSlice = reflect.Append(newSlice, reflect.ValueOf("Reflect Item"))

    fmt.Printf("Dynamic Slice: %v | Len: %d\n", newSlice, newSlice.Len())
}

// --------------------------------------------------------------------------
// 5. THE MAIN EXECUTION 
// --------------------------------------------------------------------------
func main() {
    fmt.Println("--- 1. Basics ---")
    demonstrateBasics(42)

    fmt.Println("\n--- 2. Structs & Tags ---")
    u := User{ID: 1, Email: "gopher@golang.org"}
    inspectStruct(u)

    fmt.Println("\n--- 3. Modifying Data ---")
    x := "Original"
    fmt.Println("Before:", x)
    modifyValue(&x) // Must pass address 
    fmt.Println("After: ", x)

    fmt.Println("\n--- 4. Dynamic Collections ---")
    manipulateCollection()
}
