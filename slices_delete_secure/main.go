package main 

func DeleteSecure[T any](s []T, i int) []T {
    // 1. Perform the move 
    s = append(s[:i], s[i+1:]...)

    // 2. CLear the last element in the underlying array 
    // We create the 'zero value' for type T (0 for int, nil for pointers, etc.)
    var zero T
    s[:cap(s)][len(s)] = zero

    return s
}

func DS[T any](s []T, i int) []T {
    s = append(s[:i], s[i+1:]...)
    var zero T
    s[:cap(s)][len(s)] = zero 
    return s
}
