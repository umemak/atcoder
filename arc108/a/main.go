package main

import (
	"fmt"
	"math"
)

func main() {
	var s, p int
	fmt.Scan(&s, &p)
	var n, m int
	for n = 1; n <= p; n++ {
		if math.Mod(float64(p), float64(n)) != 0 {
			continue
		}
		m = p / n
		if n+m == s {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
