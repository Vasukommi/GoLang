package main

import "fmt"

func main() {
	isLoggedIn := true
	if isLoggedIn {
		fmt.Println("Logged in")
	} else {
		fmt.Println("logged out")
	}
}