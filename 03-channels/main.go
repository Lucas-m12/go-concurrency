package main

import (
	"time"
)

func main() {
	channel := make(chan string)
	go write("Hello World!", channel)
	// The 2 FOR blocks are doing the same thing
	for {
		message, isOpen := <-channel
		if !isOpen {
			break
		}
		println(message)
	}
	for message := range channel {
		println(message)
	}
	println("end")
}

func write(text string, channel chan string) {
	for index := 0; index < 5; index++ {
		channel <- text
		time.Sleep(time.Second)
	}
	close(channel)
}
