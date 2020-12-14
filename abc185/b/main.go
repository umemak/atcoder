package main

import "fmt"

func main() {
	var n, m, t, a, b int
	fmt.Scan(&n, &m, &t)
	for i := 0; i <= t; i++ {
		// fmt.Printf("i:%d, n:%d\n", i, n)
		if m > 0 {
			fmt.Scan(&a, &b)
			m--
		}
		if i < a {
			n = n - (a - i)
			i = a
		} else if i < b {
			n = n + (b - a)
			i = b
		}
		if n <= 0 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
