package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n%3 == 0 {
		fmt.Println(0)
		return
	}
	sn := string(n)
	k := len(sn)
	for i := 1; i < k; i++ {

		if n%3 == 0 {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(-1)
}
