package main

import (
	"fmt"
)

const PI = 3.14

func areaOfCircle(radius float64) float64 {
	areaOfCircle := PI * radius * radius
	return areaOfCircle
}

func main() {
	var radius float64
	fmt.Print("Enter the radius of the circle: ")
	fmt.Scanf("%f", &radius)
	area := areaOfCircle(radius)
	fmt.Printf("The area of the circle with radius %.2f is %.2f\n", radius, area)
}
