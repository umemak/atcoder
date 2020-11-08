package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	amax := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		if amax < a[i] {
			amax = a[i]
		}
	}
	x := make([]int, 1001)
	for i := 2; i <= amax; i++ {
		for j := 0; j < n; j++ {
			if a[j]%i == 0 {
				x[i]++
			}
		}
	}
	xmax := 0
	ans := 0
	for i := 2; i <= amax; i++ {
		if xmax < x[i] {
			xmax = x[i]
			ans = i
		}
	}
	fmt.Println(ans)
}
