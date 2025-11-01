package main

import "fmt"

func birthDay(age int) { // it just take a value and copy that in increament it not update an argument
	fmt.Printf("age before increament is %v\n", age) // 20
	age++
	fmt.Printf("age after increament is %v\n", age) // 21

}

func pointerBirthDay(pointerAge *int) {
	fmt.Printf("pointer is %v and value is %v\n", pointerAge, *pointerAge) // 20
	*pointerAge++

}

func main() {
	age := 20
	birthDay(age)
	fmt.Println(age) /// 20
	pointerBirthDay(&age)
	fmt.Println(age) // 21

}
