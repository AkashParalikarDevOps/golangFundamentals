package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number:")
	var number int
	fmt.Scanf("%d", &number)
	switch {
	case number < 0:
		fmt.Println("The number is negative.")
	case number == 0:
		fmt.Println("The number is zero.")
	case number > 0:
		fmt.Println("The number is positive.")
	default:
		fmt.Println("This case should never happen.")
	}
}
