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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {

	n := nextInt()
	words := make([]string, n)

	for i := 0; i < n; i++ {
		words[i] = nextLine()
	}

	set := make(map[string]struct{})
	for i := n - 1; i >= 0; i-- {
		if _, ok := set[words[i]]; ok {
			continue
		}
		fmt.Println(words[i])
		set[words[i]] = struct{}{}
	}
}
