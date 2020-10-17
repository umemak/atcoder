package main

import (
    "bufio"
    "fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
    sc.Scan()
    i, e := strconv.Atoi(sc.Text())
    if e != nil {
        panic(e)
    }
    return i
}

func main() {
	sc.Split(bufio.ScanWords)
	a := nextInt()
	r := nextInt()
	n := nextInt()
	for i := 1; i < n; i++ {
		a = a * r
		if a > 1000000000 {
			fmt.Println("large");
			break
		}
	}
	if a <= 1000000000 {
		fmt.Println(a)
	}
}
