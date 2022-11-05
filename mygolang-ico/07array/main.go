package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "vasu"
	names[1] = "ramu"
	names[2] = "kesava"
	names[3] = "siva"
	fmt.Println(len(names))
	var veg = [4]string{"potato", "beans", "carrot", "pees"}
	fmt.Println(veg)
}