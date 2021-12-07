package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	for sc.Scan() {
		// 1行ずつ読み込む
		line := sc.Text()
		chars := strings.Split(line, " ")
		A := make(map[string]int)

		for _, s := range chars {
			A[s]++
		}

		// 出現順に結果を出力
		for s, n := range A {
			fmt.Printf(fmt.Sprintf("%s %d\n", s, n))
		}
	}
}
