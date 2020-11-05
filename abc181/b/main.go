package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	ans := 0
	var a, b int
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b)
		for j := a; j <= b; j++ {
			ans += j
		}
	}
	fmt.Println(ans)
}
