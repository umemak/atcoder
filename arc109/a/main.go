package main

import (
	"fmt"
)

func main() {
	var a, b, x, y int
	fmt.Scan(&a, &b, &x, &y)
	if a == b {
		fmt.Println(x)
		return
	}
	if a < b {
		if x*2 >= y {
			fmt.Println((b-a)*y + x)
		} else {
			fmt.Println((b-a)*(x*2) + x)
		}
		return
	}
	if x*2 >= y {
		fmt.Println((a-b-1)*y + x)
	} else {
		fmt.Println((a-b-1)*(x*2) + x)
	}
}
