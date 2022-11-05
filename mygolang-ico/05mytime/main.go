package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2000 Monday"))
	fmt.Print(time.Now().UTC())
}
