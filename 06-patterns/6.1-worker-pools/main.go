package main

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}
	return fibonacci(position-2) + fibonacci(position-1)
}

func worker(schedules <-chan int, results chan<- int) {
	for schedule := range schedules {
		results <- fibonacci(schedule)
	}
}

func main() {
	schedules := make(chan int, 45)
	results := make(chan int, 45)
	go worker(schedules, results)
	go worker(schedules, results)
	go worker(schedules, results)
	go worker(schedules, results)
	for index := 0; index < 45; index++ {
		schedules <- index
	}
	close(schedules)
	for index := 0; index < 45; index++ {
		results := <-results
		println(results)
	}
}
