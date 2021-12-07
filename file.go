package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile(("data.txt")) // data is slice of byte

	if err != nil {
		fmt.Println("File reading error", err)
	}
	fmt.Println(string(data))

	content := "Love is not never dead"
	ioutil.WriteFile("data.txt", []byte(content), 0777)
}
