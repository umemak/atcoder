package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b)
	fmt.Scan(&c, &d)
	fmt.Println(a*d - b*c)
}
