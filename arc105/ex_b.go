package main

import (
	"fmt"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	ans := a[0]
	for i := 1; i < n; i++ {
		ans = gcd(ans, a[i])
	}
	fmt.Println(ans)

	// for {
	// 	m := make(map[int]bool)
	// 	uniq := []int{}
	// 	for _, ele := range a {
	// 		if !m[ele] {
	// 			m[ele] = true
	// 			uniq = append(uniq, ele)
	// 		}
	// 	}
	// 	a = uniq
	// 	n = len(a)
	// 	sort.Ints(a)
	// 	min := a[0]
	// 	max := a[n-1]
	// 	if min == max {
	// 		break
	// 	}
	// 	a[n-1] = max - min
	// }
	// fmt.Println(a[0])
}
