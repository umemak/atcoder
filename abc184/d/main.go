package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	d := a + b + c
	ans := (100-a)*(a/d) + (100-b)*(b/d) + (100-c)*(c/d)
	fmt.Printf("%f", ans)
}
