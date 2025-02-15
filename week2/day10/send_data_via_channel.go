package main

import (
	"fmt"
)

func sendData(ch chan string) {
	ch <- "Hello form Goroutine"
}

func main() {
	messageChannel := make(chan string) // Create a channel

	go sendData(messageChannel) // Run in a goroutine

	message := <-messageChannel // Receive data from channel
	fmt.Println(message)
}
