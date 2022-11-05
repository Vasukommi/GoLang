package main

import (
	"fmt"
)

func main() {
	languages := make(map[string]string)
	languages["py"] = "python"
	languages["js"] = "javascript"
	languages["go"] = "goLang"
	
	for key, value := range languages {
		fmt.Println(languages[key])
		fmt.Println(value)
	}
}