package main

import (
	"sync"
	"time"
)

func main() {
	var waitgroup sync.WaitGroup
	waitgroup.Add(2)
	go func() {
		write("Hello World!")
		waitgroup.Done()
	}()
	go func() {
		write("Programming at Go!")
		waitgroup.Done()
	}()
	waitgroup.Wait()
}

func write(text string) {
	for index := 0; index < 5; index++ {
		println(text)
		time.Sleep(time.Second)
	}
}
