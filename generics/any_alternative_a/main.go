package main 

// Instead of Map map[string]any  
type TypedMap[V any, K comparable] map[string]V

func main () {
    // Usage 
    m := make(TypedMap[int])
    m["age"] = 30 // No boxing! No heap escape for the int.
}

