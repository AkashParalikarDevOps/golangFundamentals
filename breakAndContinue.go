package main

import (
	"fmt"
)

func breakAndContinueDemo(no int) int {
	for i := no; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("printing any number that is divisible by 3 '%v'\n", i)
			continue // Skip the rest of the loop for this iteration
		} else if i >= 50 {
			fmt.Printf("break the series as number is greater than 50 '%v'\n", i)
			break // Exit the loop if the number is greater than or equal to 50
		}
		// This will execute only if the number is not divisible by 3 and less than 50
		fmt.Printf("printing any number that is not divisible by 3 '%v'\n", i)
	}
	return 0
}

func main() {
	fmt.Println("Demonstrating break and continue in Go:")
	var input int
	fmt.Print("Enter a number to start the loop: ")
	fmt.Scanf("%d", &input)
	if input < 1 {
		fmt.Println("Please enter a number greater than 0.")
		return
	}
	fmt.Printf("Starting the loop from %d\n", input)
	breakAndContinueDemo(input)
}
