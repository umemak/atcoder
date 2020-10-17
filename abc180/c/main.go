package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Scan(&n)
	fmt.Println(1)
	for i := 2.0; i <= n/2; i++ {
		if math.Mod(n, i) == 0 {
			fmt.Println(int64(i))
		}
	}
	fmt.Println(int64(n))
}
