package main

import (
	"fmt"

	"struct-method.com/begin/data"
)

func main() {
	var max = data.Instuctor{Id: 5, LastName: "miliano"}
	max.FirstName = "Max"
	// kyle := data.NewInstuctor("kyle", "sympson")

	goCourse := data.Course{Name: "Go Fundamentals", Id: 5, Instuctor: max}
	swiftWorkShop := data.NewWorkShop("Switf With IoS", max)

	var courses [2]data.Signable

	courses[0] = goCourse
	courses[1] = swiftWorkShop //

	fmt.Println(swiftWorkShop)
	fmt.Println(goCourse)
	// fmt.Println(goCourse)
	// fmt.Println(max)
	// fmt.Println(max.PrintInfo())
	// fmt.Println(kyle.PrintInfo())
}
