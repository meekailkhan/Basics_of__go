package main

import "fmt"

func main() {
	age := 12

	if message := ""; age > 18 {
		message = "you are adult"
		fmt.Println(message)
	} else {
		message = "you are not adult"
		fmt.Println(message)
	}
}
