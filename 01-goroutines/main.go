package main

import "time"

func main() {
	go write("Hello World!")
	write("Programming at Go!")
}

func write(text string) {
	for {
		println(text)
		time.Sleep(time.Second)
	}
}
