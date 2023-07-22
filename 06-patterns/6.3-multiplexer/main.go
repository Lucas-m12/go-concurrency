package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplex(
		write("Hello World!"),
		write("Hello World 2!"),
	)
	for index := 0; index < 10; index++ {
		fmt.Println(<-channel)
	}
}

func multiplex(inputChannel1, inputChannel2 <-chan string) <-chan string {
	outputChannel := make(chan string)
	go func() {
		for {
			select {
			case message := <-inputChannel1:
				outputChannel <- message
			case message := <-inputChannel2:
				outputChannel <- message
			}
		}
	}()
	return outputChannel
}

func write(text string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			channel <- fmt.Sprintf("Received value: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()
	return channel
}
