package main

import "fmt"

type User struct {
	id   int
	name string
}

func main() {
	var u1 = User{id: 5, name: "meekail"}
	u2 := User{id: 6, name: "muzaahid"}
	du := User{}
	du2 := User{4, "user2"}
	fmt.Println(u1.name)
	fmt.Println(u2.id)
	fmt.Println(du)
	fmt.Println(du2)
}
