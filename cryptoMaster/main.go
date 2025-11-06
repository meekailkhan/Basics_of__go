package main

import (
	"fmt"

	"cryptomaster.com/begin/api"
)

func main() {
	rate, err := api.GetRate("BTC")
	fmt.Println(rate, err)
}
