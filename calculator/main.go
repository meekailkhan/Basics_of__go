package main

import (
	"fmt"
)

func main() {
	var opration string
	var num1, num2 int

	fmt.Println("CALCULATOR GO 1.0")
	fmt.Println("=================")
	fmt.Println("which opration do you want to perform (add,divide,substarct,multiply)")
	fmt.Scanf("%s", &opration)
	fmt.Println("Enter First Number")
	fmt.Scanf("%d", &num1)
	fmt.Println("Enter Second Number")
	fmt.Scanf("%d", &num2)
	fmt.Println("Load Result ...")
	switch opration {
	case "add":
		fmt.Println(num1 + num2)
	case "substract":
		fmt.Println(num1 - num2)
	case "multiply":
		fmt.Println(num1 * num2)
	case "divide":
		fmt.Println(num1 / num2)
	default:
		fmt.Println("Enter a valid opration")
	}

}
