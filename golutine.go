package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex
var wg sync.WaitGroup

func run(str string) {
	lock.Lock()
	for i := 0; i < 4; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("%s goroutine run %d\n", str, i)
	}
	lock.Unlock()
	wg.Done()
}

func main() {

	wg.Add(2)
	go run("Car")
	run("Train")
	wg.Wait()

}
