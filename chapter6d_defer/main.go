package main

import "fmt"

func main() {
	defer second()
	first()
}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

/*
defer is often used when resources need to be freed in some way. For example, when
we open a file, we need to make sure to close it later. With defer:

f, _ := os.Open(filename)
defer f.Close()

This has three advantages:
• It keeps our Close call near our Open call so it’s easier to understand.
• If our function had multiple return statements (perhaps one in an if and one in
an else), Close will happen before both of them.
• Deferred functions are run even if a runtime panic occurs.
*/
