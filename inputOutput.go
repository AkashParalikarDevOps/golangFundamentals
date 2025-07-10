package main

import "fmt"

func main() {
	var name string
	age := 0
	fmt.Print("Enter the child's name:")
	fmt.Scanf("%s", &name)
	fmt.Print("Enter the child's age:")
	fmt.Scanf("%d", &age)
	fmt.Printf("%s is %d years old\n", name, age)
	if age < 18 {
		fmt.Printf("%s is a minor\n", name)
	} else {
		fmt.Printf("%s is an adult\n", name)
	}
	if age < 13 {
		fmt.Printf("%s is a child\n", name)
	}
	if age >= 13 && age < 18 {
		fmt.Printf("%s is a teenager\n", name)
	}
	if age >= 18 && age < 65 {
		fmt.Printf("%s is a young adult\n", name)
	}
	if age >= 65 {
		fmt.Printf("%s is a senior citizen\n", name)
	}
}
