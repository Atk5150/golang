package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.OpenFile("data.txt", os.O_RDONLY|os.O_APPEND, 0777)
	defer f.Close()

	data := make([]byte, 10)

	var dataSlice = []string{}

	for {
		count, err := f.Read(data)
		if count <= 0 {
			break
		}

		if err != nil {
			fmt.Println(err)
		}
		dataSlice = append(dataSlice, string(data[:count]))
	}
	fmt.Print(dataSlice)
}
