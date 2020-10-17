package main

import (
    "bufio"
    "fmt"
	"os"
	"strconv"
	"gonum.org/v1/gonum/mat"
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
	q := nextInt()
	x := make([]float64, n * n)
	for i := 1; i <= n * n; i++ {
		x[i - 1] = float64(i)
	}
	A := mat.NewDense(n, n, x)
	a := 0
	b := 0

	for i := 0; i < q; i++ {
		d := nextInt()
		if d != 3 {
			a = nextInt() - 1
			b = nextInt() - 1
		}
		if d == 1 {
		}
		if d == 2 {
		}
		if d == 3 {
			A = A.T();
		}
		if d == 4 {
			fmt.Println(A.At(a, b))
		}
	}
}
