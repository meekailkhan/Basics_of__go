package main

var b = "global scope variable"

func main() {
	var a string
	a = "abcd"
	print(b, "\n") // global scope variable
	b := "local scope variable"

	print(a, " ", b) // abcd local scope variable
}
