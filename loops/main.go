package main

import "fmt"

func main() {
	count := 0
	fmt.Printf("testing\n")

	for count < 10 {
		count++
		fmt.Printf("count is %v\n", count)
	}

	var arr = [5]int{1, 2, 3, 4, 5}

	for index := range arr {
		fmt.Printf("index is %v and value is %v\n", index, arr[index])
	}

	for index, value := range arr {
		fmt.Println(index, value)
	}
}
