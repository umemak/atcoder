package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	var s string
	fmt.Scan(&s)
	for i := 0; i < n; i++ {
		if s[i] == "o"[0] {
			x++
		} else if x > 0 {
			x--
		}
	}
	fmt.Println(x)
}
