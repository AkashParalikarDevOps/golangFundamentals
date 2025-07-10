package main

import (
	"fmt"
	"strconv"
)

func floatToIntAndIntToFloat(floatValue float64, intValue int) {
	// Convert float64 to int
	intFromFloat := int(floatValue)
	fmt.Printf("Converted float %f to int: %d\n", floatValue, intFromFloat)

	// Convert int to float64
	floatFromInt := float64(intValue)
	fmt.Printf("Converted int %d to float: %f\n", intValue, floatFromInt)
}

func strToIntAndIntToStr(strValue string, intValue int) {
	stringFromInt := strconv.Itoa(intValue)
	fmt.Printf("Converted int %d to string: %s\n", intValue, stringFromInt)
	// Convert string to int
	intFromString, err := strconv.Atoi(strValue)
	if err != nil {
		fmt.Printf("Error converting string %s to int: %v\n", strValue, err)
		return
	}
	fmt.Printf("Converted string %s to int: %d\n", strValue, intFromString)
	// Convert int to string
}

func main() {
	var floatValue float64 = 3.14
	var intValue int = 4
	var strValue string = "42"
	floatToIntAndIntToFloat(floatValue, intValue)
	strToIntAndIntToStr(strValue, intValue)
	var strValueString string = "Hello, World!"
	strToIntAndIntToStr(strValueString, intValue)
}
