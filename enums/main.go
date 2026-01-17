package main

import "fmt"

// 1. Define the custom type (The "Enum" name)
type Page int

// 2. Define the values using iota (The "Enum" members)
// We prefix them with 'Page' to mimic the Page.Demo namespacing in TS.
const (
	PageDemo Page = iota
	PageDanger
	PageAngel
	PageFix
	PageFocus
)

var allPages = []Page{
	PageDemo,
	PageDanger,
	PageAngel,
	PageFix,
	PageFocus,
}

var pageNames = []string{
	"Demo",
	"Danger",
	"Angel",
	"Fix",
	"Focus",
}

// 3. Implement the Stringer interface
// This allows fmt.Println to print "PageDemo" instead of "0"
func (p Page) String() string { // Receiving function: the type Page receives the String method
	// Safety check to prevent index out of bounds
	if int(p) < 0 || int(p) >= len(pageNames) {
		return fmt.Sprintf("Page(%d)", p)
	}
	return pageNames[p]
}

// 4. Create an "All" slice (The equivalent of Object.values(Page))
func AllPages() []Page {
	return allPages
}

func main() {
	// Usage as a single value
	currentPage := PageAngel
	fmt.Printf("current Page: %s (Value: %d)\n", currentPage, currentPage)

	// Usage as a list (the "Enum" collection)
	fmt.Println("List of all pages:")
	for _, p := range AllPages() {
		fmt.Printf(" - %s\n", p)
	}

	// Demonstrating type safety
	printPage(PageDanger)
	outOfBoundPage := Page(69)
	fmt.Printf("This page out of register gives us %s\n", outOfBoundPage)
}

func printPage(p Page) {
	fmt.Println("Route navigating to:", p)
}
