package main

import (
	"fmt"
	"os"
	"sync"
)

var wg sync.WaitGroup
var mapFile = map[int]string{}

func main() {
	fileName := "data3.txt"
	sfile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(nil)
	}

	info, _ := os.Stat(fileName)
	size := info.Size()
	var goroutineFileSize int64
	var goroutineCount int64 = 3
	if size%goroutineCount == 0 {
		goroutineFileSize = size / goroutineCount
	} else {
		goroutineFileSize = size / goroutineCount
		goroutineCount++
	}

	fmt.Printf("File Total : %v\n", size)
	fmt.Printf("Goroutine Count ; %v\n", goroutineCount)
	fmt.Printf("Goroutine File Size : %v\n", goroutineFileSize)

	wg.Add(int(goroutineCount))
	for i := 0; i < int(goroutineCount); i++ {
		go ReadFile(sfile, i, goroutineFileSize)
	}
	wg.Wait()

	for key, value := range mapFile {
		fmt.Printf("%s->%s\n", key, value)
	}
	defer sfile.Close()
}

func ReadFile(sfile *os.File, i int, goroutineFileSize int64) {
	b := make([]byte, goroutineFileSize)
	sfile.Seek(int64(i)*goroutineFileSize, 0)
	sfile.Read(b)
	mapFile[i] = string(b)
	wg.Done()
}
