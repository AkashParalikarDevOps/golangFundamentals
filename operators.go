package main

import "fmt"

func enterTwoNumbers() (int, int) {
	var a, b int
	fmt.Print("Enter first number: ")
	fmt.Scanf("%d", &a)
	fmt.Print("Enter second number: ")
	fmt.Scanf("%d", &b)
	return a, b
}

func main() {
	a, b := enterTwoNumbers()
	fmt.Println("Addition:", a+b)       // Addition
	fmt.Println("Subtraction:", a-b)    // Subtraction
	fmt.Println("Multiplication:", a*b) // Multiplication
	fmt.Println("Division:", b/a)       // Division
	fmt.Println("Modulus:", b%a)        // Modulus
	fmt.Println("Bitwise AND:", a&b)    // Bitwise AND
	fmt.Println("Bitwise OR:", a|b)     // Bitwise OR
	fmt.Println("Bitwise XOR:", a^b)    // Bitwise XOR
	fmt.Println("Left Shift:", a<<1)    // Left Shift
	fmt.Println("Right Shift:", b>>1)   // Right Shift
}
