package main

import "fmt"

func main() {
    fmt.Println("--- SLICE DELETIONS ---")
    pureDeletion()
    standardDeletion()
    safeStandardDeletion()
}

func safeStandardDeletion() {
    fmt.Println("\n--- SAFE STANDARD IDIOMATIC DELETION ---")
    nums := []int{10, 20, 30, 40, 50}
    indexToDelete := 2

    fmt.Printf("Initial: %v (Len: %d, Cap: %d)\n", nums, len(nums), cap(nums))
    nums = append(nums[:indexToDelete], nums[indexToDelete + 1:]...)
    nums[:cap(nums)][len(nums)] = 0
    fmt.Printf("After Deletion: %v (Len: %d, Cap: %d)\n", nums, len(nums), cap(nums))

    fmt.Println("\n...Demonstrating insecure deletion")
    // --- THE PROOF
    // We "stretch" the slice back to its original capacity to see 
    // what is still hiding in the memory buffer. 
    ghostData := nums[:cap(nums)]
    fmt.Printf("The 'Ghost' Memory: %v\n", ghostData)

    if ghostData[len(ghostData) - 1] != 50 {
        fmt.Println("⚠️  SECURITY ALERT: The value '50' got replaced by 0 so it can get GC'd!")
    }
}

func standardDeletion() {
    fmt.Println("\n--- STANDARD IDIOMATIC DELETION ---")
    nums := []int{10, 20, 30, 40, 50}
    indexToDelete := 2

    fmt.Printf("Initial: %v (Len: %d, Cap: %d)\n", nums, len(nums), cap(nums))
    nums = append(nums[:indexToDelete], nums[indexToDelete + 1:]...)
    fmt.Printf("After Deletion: %v (Len: %d, Cap: %d)\n", nums, len(nums), cap(nums))

    fmt.Println("\n...Demonstrating insecure deletion")
    // --- THE PROOF
    // We "stretch" the slice back to its original capacity to see 
    // what is still hiding in the memory buffer. 
    ghostData := nums[:cap(nums)]
    fmt.Printf("The 'Ghost' Memory: %v\n", ghostData)

    if ghostData[len(ghostData) - 1] == 50 {
        fmt.Println("⚠️  SECURITY ALERT: The value '50' is still in memory at the end of the array!")
    }
}

func pureDeletion() {
    fmt.Println("\n--- PURE DELETION ---")

    original := []int{10, 20, 30, 40, 50}
    i := 2

    pure := make([]int, len(original) -1)

    copy(pure[:i], original[:i])
    copy(pure[i:], original[i+1:])

    fmt.Printf("original: %v\n", original)
    fmt.Printf("pure copy: %v\n", pure)
}
