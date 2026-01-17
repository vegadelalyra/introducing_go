package main

import "fmt"

func main() {
	fmt.Println("[x:y] len: y - x")
	fruitsSlice()
	grow()
	subslice()
	modifyUnderlyingArray()
	makeFn()
	copyCmd()
	filterSlices()
    deleteSliceElms()
}

func deleteSliceElms() {
    fmt.Println("\n----------------------------------")
	fmt.Println("Deleting specific els from slice!")

    slice := []int{10, 20, 30, 40, 50}
    idxToDel := 2 

    delSliceElm := func() []int {
        slice = append(slice[:idxToDel], slice[idxToDel + 1:]...)
        slice[:cap(slice)][len(slice)] = 0  // safe idiom slice deletion
        return slice 
    }

    fmt.Printf("Original slice: %v\n", slice)
    fmt.Printf("Elm to delete: %v\n", slice[idxToDel])
   
    slice = delSliceElm()

    fmt.Printf("Final slice: %v\n", slice)

}

func filterSlices() {
    fmt.Println("\n----------------------------------")
	fmt.Println("Filtering slices!")

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
    result := make([]int, 0, 20)

    // closure!
    filterEven := func() []int {
        for _, v := range nums {
            if v % 2 == 0 { 
                result = append(result, v)
            } 
        }
        return result
    }

    result = filterEven()
    fmt.Printf("original slice: %v\n", nums)
    fmt.Printf("filtered slice: %v\n", result)
}

func copyCmd() {
	fmt.Println("\n----------------------------------")
	fmt.Println("Copy a slice")
	source := []string{"A", "B", "C"}
	destination := make([]string, len(source))
	copy(destination, source)
	fmt.Printf("copy: %v\noriginal: %v\n", destination, source)
	destination[0] = "Z"
	fmt.Println("after modifying copy...")
	fmt.Printf("copy: %v\noriginal: %v\n", destination, source)
}

func makeFn() {
	fmt.Println("\n----------------------------------")
	fmt.Println("Create an empty slice with make and populate it with for")
	mySlice := make([]int, 0, 10)
	for i := range cap(mySlice) {
		mySlice = append(mySlice, i)
		fmt.Printf(
			"my slice now is %v\n (len: %d, cap: %d)\n",
			mySlice,
			len(mySlice),
			cap(mySlice),
		)
	}
}

func modifyUnderlyingArray() {
	original := []int{1, 2, 3, 4, 5}
	original2 := []int{1, 2, 3, 4, 5}
	sub := original[1:3]
	sub2 := make([]int, 2)
	copy(sub2, original2[1:3])

	fmt.Println("\n----------------------------------")
	fmt.Println("Modifying Underlying Array")
	fmt.Printf("original before indexation: %v\n", original)
	sub[0] = 99
	fmt.Printf("original after indexation: %v\n", original)
	fmt.Println("With copy method:")
	fmt.Printf("BEFORE: original 2: %v\n sub 2: %v\n", original2, sub2)
	sub2[0] = 99
	fmt.Printf("AFTER: original 2: %v\n sub 2: %v\n", original2, sub2)
}

func subslice() {
	fmt.Println("----------------------------------")
	fmt.Println("Subslices!")

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	firstHalf := s[:5:5]
	secondHalf := s[5:]
	fmt.Printf("firstHalf: %v\nsecondHalf: %v", firstHalf, secondHalf)
}

func grow() {
	var slice []int
	oldCap := cap(slice)

	fmt.Printf("Initial State: Length: %d, Capacity: %d\n", len(slice), oldCap)
	fmt.Println("----------------------------------")

	for i := 1; i <= 10; i++ {
		slice = append(slice, i)
		currentCap := cap(slice)

		// Only print when the capacity changes
		if currentCap != oldCap {
			fmt.Printf("Appended: %d | New capacity: %d (Changed from %d) \n", i, currentCap, oldCap)
			oldCap = currentCap
		}
	}
}

func fruitsSlice() {
	fruits := []string{
		"Pineapple",
		"Apple",
		"Pear",
		"Strawberry",
		"Watermelon",
	}

	fmt.Println(fruits)

	fmt.Printf(
		"My fruits slice has %d length and %d capacity\n",
		len(fruits),
		cap(fruits),
	)
}
