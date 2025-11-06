package main

import (
	"fmt"
	"time"
)

func printMsg(text string, t float64, channel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%v, %v\n", text, i)
		time.Sleep(time.Duration(t) * time.Millisecond)
	}
	channel <- "Done"
	channel <- "Complete"
}

func main() {
	channel := make(chan string, 2)
	go printMsg("Go Is Amazing", 1000, channel)

	fmt.Println("========start waiting for channel=========")
	response := <-channel
	anotherRes := <-channel
	fmt.Println("========end waiting for channel==========")

	fmt.Println("response is comming with", response)
	fmt.Println("another response is", anotherRes)
	fmt.Println("========end file execution==========")

}
