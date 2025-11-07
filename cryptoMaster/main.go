package main

import (
	"fmt"

	"cryptomaster.com/begin/api"
)

func main() {
	rate, err := api.GetRate("2")
	fmt.Println(rate, err)
}
