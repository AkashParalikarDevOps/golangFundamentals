package main

import "fmt"

func main() {
	//if elseif else statement

	var number int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &number)
	if number < 0 {
		fmt.Println("The number is negative.")
	} else if number == 0 {
		fmt.Println("The number is zero.")
	} else if number > 0 {
		fmt.Println("The number is positive.")
	} else {
		fmt.Println("This case should never happen.")
	}

	//for loop example accept the user string input
	fmt.Print("Enter a string: ")
	var input string
	fmt.Scanf("%s", &input)
	fmt.Println("You entered:", input)
	//for loop count the number of characters in the string
	count := 0
	for range input {
		count++
	}
	fmt.Printf("The string '%s' has %d characters.\n", input, count)
	//for loop to print the multiplication table for the number entered
	if count == 0 {
		fmt.Println("The string is empty, no multiplication table to display.")
		return
	}
	if count < 0 {
		fmt.Println("The count is negative, no multiplication table to display.")
		return
	}
	fmt.Println("Multiplication table for", count)
	for i := 1; i <= 100; i++ {
		count := i * count
		fmt.Println(count)
		if i == 10 {
			fmt.Println("The multiplication result exceeded 10, stopping the loop.")
			break
		} else if i > 0 {
			fmt.Println("The multiplication result is positive, continuing the loop.")
		} else {
			fmt.Println("The multiplication result is zero or negative, stopping the loop.")
			break
		}
	}
}
