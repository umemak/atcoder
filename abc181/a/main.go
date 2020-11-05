package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	c := "Black"
	if n%2 == 0 {
		c = "White"
	}
	fmt.Println(c)
}
