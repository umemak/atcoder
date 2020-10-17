package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Scan()
	s := sc.Text()
	sc.Scan()
	t := sc.Text()
	if s == "Y" {
		fmt.Println(strings.ToUpper(t))
	} else {
		fmt.Println(t)
	}
}
