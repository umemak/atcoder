package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	if (a == b+c+d) || (b == a+c+d) || (c == a+b+d) || (d == a+b+c) || (a+b == c+d) || (a+c == b+d) || (a+d == b+c) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
