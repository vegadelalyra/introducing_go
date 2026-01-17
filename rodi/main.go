package main

import "fmt"

func main() {
	mySlice := []string{"gallinas", "galletas", "golazo", "brownies", "arequipe", "tostadas"}
    copyForLeak := make([]string, len(mySlice))
    copy(copyForLeak, mySlice)

	demonstrateLeak(copyForLeak, 1)

    fmt.Print("xd")
    fmt.Scanln()
    fmt.Print("BUT")
    fmt.Scanln()
    fmt.Println("When we use safe delete...")
    fmt.Scanln()
	removeItemFromSliceSecure(mySlice, 1)
    fmt.Println(mySlice[:cap(mySlice)])

    fmt.Println("TADAAA!!")

}

func demonstrateLeak[T any](mySlice []T, i int) {
	fmt.Println()
    fmt.Println(
		"Your slice is",
		mySlice,
		"\nyour slice have len:",
		len(mySlice),
		"and cap:",
		cap(mySlice),
		"\nand you're gonna delete the elm {",
		mySlice[i],
		"}",
	)
	fmt.Scanln()
	mySlice = append(mySlice[:i], mySlice[i+1:]...)
	fmt.Println(mySlice)
    fmt.Printf(
        "Now your slice have len: %d and cap: %d\n", 
        len(mySlice), 
        cap(mySlice),
    )

	fmt.Scanln()

    fmt.Printf(
        "but that's the visible slice, the slice in memory has...\n%v\n",
        mySlice[:cap(mySlice)],
    )
	fmt.Scanln()
}

func removeItemFromSliceSecure[T any](mySlice []T, elToDel int) {
	mySlice = append(mySlice[:elToDel], mySlice[elToDel+1:]...)
	var zero T
	mySlice[:cap(mySlice)][len(mySlice)] = zero
}
