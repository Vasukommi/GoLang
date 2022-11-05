package main

import "fmt"

func main() {
	fruitsArray := []string{"apple", "banana", "peache"}
	fruitsArray = append(fruitsArray, "mango", "Dragon fruit")
	fmt.Println(fruitsArray)
	fruitsArray = append(fruitsArray[:3], fruitsArray[3:]...)
	fmt.Println(fruitsArray[1])
}