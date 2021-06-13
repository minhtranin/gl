package main

import (
	"fmt"
)

func queue(msg string) chan string {
	orderQueue := make(chan string)
	go func() {
		for i := 0; i <= 5; i++ {
			orderQueue <- fmt.Sprintf("%s %d", msg, i)
		}
	}()
	return orderQueue
}

func fanIn(chan1, chan2 chan string) chan string {
	fanInQueue := make(chan string)
	go func() {
		for {
			select {
			case <-chan1:
				fanInQueue <- <-chan1
			case <-chan2:
				fanInQueue <- <-chan2
			}
		}
	}()
	return fanInQueue
}

func main() {
	chicken := queue("chicken")
	breaks := queue("break")
	queueChan := fanIn(chicken, breaks)
	for i := 0; i <= 5; i++ {
		fmt.Println(<-queueChan)
	}
}
