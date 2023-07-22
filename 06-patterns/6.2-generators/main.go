package main

import (
	"fmt"
	"time"
)

func main() {
	channel := write("Hello World!")
	for index := 0; index < 10; index++ {
		fmt.Println(<-channel)
	}
}

func write(text string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			channel <- fmt.Sprintf("Received value: %s", text)
			time.Sleep(time.Second / 2)
		}
	}()
	return channel
}
