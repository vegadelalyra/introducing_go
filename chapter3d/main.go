package main 

import "fmt" 

func main() {
    fmt.Println("Convert from feet into meters")
    fmt.Println("1 ft = 0.3048m")
    fmt.Print("ft: ")
    var input float64 
    fmt.Scanf("%f", &input)

    output := input * 0.3048

    fmt.Println(output, "m")
}
