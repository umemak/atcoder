package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)
	for {
		var i int
		if i = strings.Index(s, "fox"); i < 0 {
			fmt.Println(len(s))
			return
		}
		s = s[:i] + s[i+3:]
	}
}
