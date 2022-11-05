package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rating_input := bufio.NewReader(os.Stdin)
	user_input, _ := rating_input.ReadString('\n')
	plus_one_rating, _ := strconv.ParseFloat(strings.TrimSpace(user_input), 64)
	fmt.Print(plus_one_rating + 1)
}