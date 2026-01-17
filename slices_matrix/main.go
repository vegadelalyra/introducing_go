package main

import "fmt"

func main() {
	fmt.Println("\nLiteral matrix [MEDIUM]")
	literal() // performance: medium
	fmt.Println("\nMake-loop matrix [GOOD]")
	makeLoop() // performance: good
	fmt.Println("\nDynamic append matrix [SLOWER]")
	dynamicAppend() // performance: slower (due to re-allocation)
	fmt.Println("\nFlat slice matrix [BEST!]")
	flatSlice() // performance: best! (cache friendly)
}

func flatSlice2() {
	rows, cols := 3, 3
	data := make([]int, rows*cols)

	for i := range data {
		data[i] = i
		i++
	}
    
    for i := range rows {
        for j := range cols {
            fmt.Printf("%d ", data[i*cols+j])
        }
        fmt.Println()
    }
}

func flatSlice() {
	rows, cols := 3, 3
	// Allocate all memory in one contiguous block
	data := make([]int, rows*cols)

	// Fill the flat slice
	for i := range data {
		data[i] = i + 1
	}

	// Print as a grid by calculating offsets
	for i := range rows {
		for j := range cols {
			var val = data[i*cols+j]
			switch j {
			case 0:
				fmt.Printf("[%d", val)
			case cols - 1:
				fmt.Printf(" %d]", val)
			default:
				fmt.Printf(" %d", val)
			}
		}
		fmt.Println()
	}
}

func dynamicAppend() {
	var matrix [][]int
	var val = 1
	rows, cols := 3, 3
	for range rows {
		var row []int
		for range cols {
			row = append(row, val)
			val++
		}
		matrix = append(matrix, row)
	}
	for _, k := range matrix {
		fmt.Println(k)
	}
}

func makeLoop() {
	rows, cols := 3, 3
	// Create the outer slice (the rows)
	matrix := make([][]int, rows)
	count := 1
	// Create the inner slice (the columns) for each row
	for i := range matrix {
		matrix[i] = make([]int, cols)
		for j := range len(matrix[i]) {
			matrix[i][j] = count
			count++
		}
	}
	// Print grid
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func literal() {
	// Initializing with a composite literal
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	// Printing in grid format
	for _, row := range matrix {
		fmt.Println(row)
	}
}
