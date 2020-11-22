package main

import "fmt"

func main() {
	var sx, sy, gx, gy int
	fmt.Scan(&sx, &sy, &gx, &gy)
	gy *= -1
	d := (gx - sx) / (gy - sy)
	x := 0 / d
	fmt.Println(x)
}
