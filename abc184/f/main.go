package main

import "fmt"

func main() {
	var n, t int
	fmt.Scan(&n, &t)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		if a[i] > t {
			a[i] = 0
		}
	}
	for {
		sum := 0
		for _, v := range a {
			sum += v
		}
		// fmt.Println(sum)
		if sum <= t {
			fmt.Println(sum)
			return
		}
		d := sum - t
		f := false
		for i := 0; i < n; i++ {
			if a[i] >= d {
				a[i] = 0
				f = true
				break
			}
		}
		if f == false {
			min := 1000000000
			k := 0
			for i, v := range a {
				if v > 0 && min > v {
					min = v
					k = i
				}
			}
			a[k] = 0
		}
		// fmt.Println(a)
	}
}
