package main

import (
	"fmt"
	"time"
)

func printMsg(text string, t float64) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%v, %v\n", text, i)
		time.Sleep(time.Duration(t) * time.Millisecond)
	}
}

func main() {
	go printMsg("Go Is Amazing", 10000)
	// go printMsg("Bye Bye!")
	printMsg("Rust is Less Popular Then Go", 2000)

}
