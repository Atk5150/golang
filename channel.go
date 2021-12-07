package main

import (
	"fmt"
	"strconv"
)

var messages = make(chan string)
var msg string

func sendMessage() {
	messages <- "good"
	messages <- "morining"
	for i := 0; i < 5; i++ {
		messages <- "message" + strconv.Itoa(i)
	}
	close(messages)
}

func main() {
	go sendMessage()

	msg = <-messages
	fmt.Println(msg)
	msg = <-messages
	fmt.Println(msg)

	for msg := range messages {
		fmt.Println(msg)
	}
}
