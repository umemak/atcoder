package main

import (
	"fmt"
	"strings"
)

func moveto(cx, cy, tx, ty int) string {
	ret := ""
	if cx < tx {
		ret += strings.Repeat("D", tx-cx)
	}
	if cx > tx {
		ret += strings.Repeat("U", cx-tx)
	}
	if cy < ty {
		ret += strings.Repeat("R", ty-cy)
	}
	if cy > ty {
		ret += strings.Repeat("L", cy-ty)
	}
	return ret
}

func space(x, y []int, tx, ty, c int) bool {
	for i := c; i <= 99; i++ {
		if x[i] == tx && y[i] == ty {
			return false
		}
	}
	return true
}

func main() {
	x := make([]int, 100)
	y := make([]int, 100)
	for i := 0; i <= 99; i++ {
		fmt.Scan(&x[i], &y[i])
	}
	ans := ""
	cx := 0
	cy := 0
	for i := 0; i < 99; i++ {
		// 次の数字が経路内にある？
		if cx <= x[i+1] && x[i+1] <= x[i] && cy <= y[i+1] && y[i+1] <= y[i] {
			// 先に次の数字を拾う
			ans += moveto(cx, cy, x[i+1], y[i+1])
			cx = x[i+1]
			cy = y[i+1]
			// 目的の数字に近い空きスペースを探す
			sx := 0
			sy := 0
			if cx <= x[i] {
				if space(x, y, x[i]-1, y[i], i) {
					sx = x[i] - 1
					sy = y[i]
				} else if space(x, y, x[i]-2, y[i], i) {
					sx = x[i] - 2
					sy = y[i]
				}
			} else if cx > x[i] {
				if space(x, y, x[i]+1, y[i], i) {
					sx = x[i] + 1
					sy = y[i]
				} else if space(x, y, x[i]+2, y[i], i) {
					sx = x[i] + 2
					sy = y[i]
				}
			} else if cy <= y[i] {
				if space(x, y, x[i], y[i]-1, i) {
					sx = x[i]
					sy = y[i] - 1
				} else if space(x, y, x[i], y[i]-2, i) {
					sx = x[i]
					sy = y[i] - 2
				}
			} else if cy > y[i] {
				if space(x, y, x[i], y[i]+1, i) {
					sx = x[i]
					sy = y[i] + 1
				} else if space(x, y, x[i], y[i]+2, i) {
					sx = x[i]
					sy = y[i] + 2
				}
			}
			// 空きがなければ拾わない
			if sx != 0 && sy != 0 {
				ans += "I"
				// 空きがあればそこに落とす
				ans += moveto(cx, cy, sx, sy)
				cx = sx
				cy = sy
				ans += "O"
				ans += moveto(cx, cy, x[i], y[i])
				cx = x[i]
				cy = y[i]
				ans += "I"
				ans += moveto(cx, cy, sx, sy)
				cx = sx
				cy = sy
				ans += "I"
				i++
			} else {
				ans += moveto(cx, cy, x[i], y[i])
				cx = x[i]
				cy = y[i]
				ans += "I"
			}
		} else {
			ans += moveto(cx, cy, x[i], y[i])
			cx = x[i]
			cy = y[i]
			ans += "I"
		}
	}
	ans += moveto(cx, cy, x[99], y[99])
	ans += "I"
	fmt.Println(ans)
}
