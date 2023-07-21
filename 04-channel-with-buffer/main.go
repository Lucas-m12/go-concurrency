package main

func main() {
	channel := make(chan string, 2)
	channel <- "Hello World!"
	channel <- "Hello World 2!"
	message := <-channel
	message2 := <-channel
	println(message, message2)
}
