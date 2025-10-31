package main

import "fmt"

var arr [5]int

func init() {
	arr[0] = 5
	println("first init function", arr[0])
}

func init() {
	arr[3] = 3
	println("second init function", arr[3])
}

func main() {
	fmt.Println(arr)
}
