package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i])
	}
	var a1 int
	var a2 int
	var a3 int
	for i := 0; i < n; i++ {
		abs := int(math.Abs(float64(x[i])))
		a1 += abs
		a2 += abs * abs
		if a3 < abs {
			a3 = abs
		}
	}
	a22 := math.Sqrt(float64(a2))
	fmt.Println(a1)
	fmt.Println(a22)
	fmt.Println(a3)
}
