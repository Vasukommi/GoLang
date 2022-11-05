package main

import "fmt"

func main() {
	vasu := User{"vasu kommi", "itsvasukommi@gmail.com", 23, true}	 	
	fmt.Println(vasu.name)
}

type User struct {
	name   string
	email  string
	age    int
	status bool
}