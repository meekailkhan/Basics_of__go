package main

import (
	"fmt"

	"packges.com/begin/data"
)

func printData() {
	println("print data from printData file")
	println("print variable from mainfile", name)
	fmt.Println("print line by fmt package")
	fmt.Println("print line from data package", data.MaxSpeed)
}

var student = "Jane Doe"

// if file is from diffrent package of consumer file function name should be start we capital letter
// if we start function name with small letter its is private function available just for own package
// for creating public function use name start with capital letter
func greet(student string) {
	println("print line from greet function from data file", name)
}
