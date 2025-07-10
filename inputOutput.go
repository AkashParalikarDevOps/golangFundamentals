package main

import "fmt"

func main() {
	daughter := ""
	age := 0
	fmt.Print("Enter the child's name:")
	fmt.Scanf("%s/n", &daughter)
	fmt.Print("Enter the child's age:")
	fmt.Scanf("%d", &age)
	fmt.Printf("%s is %d years old\n", daughter, age)
	if age < 18 {
		fmt.Printf("%s is a minor\n", daughter)
	} else {
		fmt.Printf("%s is an adult\n", daughter)
	}
	if age < 13 {
		fmt.Printf("%s is a child\n", daughter)
	}
	if age >= 13 && age < 18 {
		fmt.Printf("%s is a teenager\n", daughter)
	}
	if age >= 18 && age < 65 {
		fmt.Printf("%s is a young adult\n", daughter)
	}
	if age >= 65 {
		fmt.Printf("%s is a senior citizen\n", daughter)
	}
}
