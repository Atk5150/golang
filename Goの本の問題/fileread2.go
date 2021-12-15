package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("data2.txt")
	if err != nil {
		log.Fatal(err)
	}
	// ファイルを１行ずつ読み込んで表示する
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}
