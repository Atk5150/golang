// 文に出現する単語の出現回数を調べる。
package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	wc := make(map[string]int)
	for _, w := range strings.Fields(s) {
		wc[w]++
	}
	return wc
}

func main() {
	s := "This is a pen. This is a piano."
	answer := WordCount(s)
	fmt.Println(answer)

}
