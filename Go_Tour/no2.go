// sliceにsliceを割り当てたものを返す
package main

import (
	"fmt"
)

func Pic(dy, dx int) [][]uint8 {
	s := make([][]uint8, dy)
	fmt.Println("s:", s)

	for i := 0; i < 2; i++ {
		s[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			fmt.Printf("s[%v][%v]: %v\n", i, j, s[i][j])
		}

	}

	fmt.Println("s:", s)
	return s
}

func main() {
	Pic(2, 4)
}
