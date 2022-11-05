package main

import "fmt"

func main() {
	var numberValue = 20
	refNumberValue := &numberValue
	addNumToRef := *refNumberValue + 1
	*refNumberValue = *refNumberValue + 5
	fmt.Println(numberValue)
	fmt.Println(addNumToRef)
}