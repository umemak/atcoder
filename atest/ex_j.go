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
	n := nextInt()
	m := nextInt()
	res := make([]int, n)
	i := 0
	j := 0
	d := 0
	for i = 0; i < n; i++ {
		res[i] = -1
	}
	for i = 0; i < m; i++ {
		d = nextInt()
		for j = 0; j < n; j++ {
			if d > res[j] {
				res[j] = d
				fmt.Println(j + 1)
				break
			}
		}
		if j == n {
			fmt.Println(-1)
		}
	}
}
