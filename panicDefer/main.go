package main

import "fmt"

func isAdult(age int) {
	if age < 18 {
		panic("you are not eligible for that action")
	} else {
		fmt.Println("user eligible for that opration lets do next work")
	}
}

func main() {
	defer fmt.Println("closing service")

	isAdult(19)
	fmt.Println("this is the action if user pass the condition")
}
