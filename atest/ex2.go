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
	q := nextInt()
	scores := make([][]int, n)
	for i:= range scores{
		scores[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			scores[i][j] = 0;
		}
	}
	for i := 1; i <= q; i++ {
		c := nextInt()
		d := nextInt() - 1
		if c == 1 {
			a := 0
			for j := 0; j < m; j++ {
				if scores[d][j] == 1 {
					b := n
					for k := 0; k < n; k++ {
						b = b - scores[k][j]
					}
					a = a + b
				}
			}
			fmt.Println(a)
		} else {
			e := nextInt()
			scores[d][e - 1] = 1
		}
	}
}
