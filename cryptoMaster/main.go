package main

import (
	"fmt"
	"sync"

	"cryptomaster.com/begin/api"
)

func main() {
	// go api.GetRate("11")
	// go api.GetRate("12")
	// go api.GetRate("13")
	// time.Sleep(time.Second * 5)
	var wg sync.WaitGroup

	users := []string{"1", "2", "3"}
	for _, usr := range users {
		wg.Add(1)
		go func(id string) {
			getUserData(id)
			wg.Done()
		}(usr)
	}
	wg.Wait()

}

func getUserData(userId string) {
	user, err := api.GetRate(userId)
	if err == nil {
		fmt.Printf("user data is %v\n", user)
	}
}
