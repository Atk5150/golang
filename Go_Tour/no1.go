// 自分で作った関数が標準ライブラリmath.Sqrtとどれくらい近似を取れるか

package main

import (
	"fmt"
	"math"
)

var state_nu float64 = math.Nextafter(1, 2) - 1

func Sqrt(x float64) (float64, int) {
	z := x / 2
	before_nu := z
	var i int
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-before_nu) <= state_nu {
			i++
			break
		}
		before_nu = z
	}
	return z, i
}

func main() {
	for i := 1; i < 11; i++ {
		x := float64(i)
		z, count := Sqrt(x)
		w := math.Sqrt(x)
		fmt.Printf("count=%2v, x=%2v, z=%19v, math.Sqrt()=%19v\n", count, x, z, w)
	}
}
