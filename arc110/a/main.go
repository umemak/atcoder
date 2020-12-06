package main

import (
	"fmt"
)

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	var n int
	fmt.Scan(&n)
	switch n {
	case 2:
		fmt.Println(3)
	case 3:
		fmt.Println(7)
	case 4:
		fmt.Println(LCM(2, 3, 4) + 1)
	case 5:
		fmt.Println(LCM(2, 3, 4, 5) + 1)
	case 6:
		fmt.Println(LCM(2, 3, 4, 5, 6) + 1)
	case 7:
		fmt.Println(LCM(2, 3, 4, 5, 6, 7) + 1)
	case 8:
		fmt.Println(LCM(2, 3, 4, 5, 6, 7, 8) + 1)
	case 9:
		fmt.Println(LCM(2, 3, 4, 5, 6, 7, 8, 9) + 1)
	case 10:
		fmt.Println(LCM(2, 3, 4, 5, 6, 7, 8, 9, 10) + 1)
	case 11:
		fmt.Println(LCM(2, 3, 4, 5, 6, 7, 8, 9, 10, 11) + 1)
	case 12:
		fmt.Println(LCM(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12) + 1)
	case 13:
		fmt.Println(LCM(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13) + 1)
	case 14:
		fmt.Println(LCM(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14) + 1)
	case 15:
		fmt.Println(LCM(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15) + 1)
	case 16:
		fmt.Println(LCM(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16) + 1)
	case 17:
		fmt.Println(LCM(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17) + 1)
	case 18:
		fmt.Println(LCM(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18) + 1)
	case 19:
		fmt.Println(LCM(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19) + 1)
	case 20:
		fmt.Println(LCM(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20) + 1)
	case 21:
		fmt.Println(LCM(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21) + 1)
	case 22:
		fmt.Println(LCM(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22) + 1)
	case 23:
		fmt.Println(LCM(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23) + 1)
	case 24:
		fmt.Println(LCM(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24) + 1)
	case 25:
		fmt.Println(LCM(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25) + 1)
	case 26:
		fmt.Println(LCM(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26) + 1)
	case 27:
		fmt.Println(LCM(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27) + 1)
	case 28:
		fmt.Println(LCM(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28) + 1)
	case 29:
		fmt.Println(LCM(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29) + 1)
	case 30:
		fmt.Println(LCM(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30) + 1)
	}
}
