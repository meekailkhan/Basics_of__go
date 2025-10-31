package main

import "fmt"

func main() {
	names := []string{"Mary", "John"}         // initialing slice under the hood inilialize array
	fmt.Println(names)                        // [Mary,John]
	newNames := append(names, "Luke", "Jane") // initialize new array with all element of old with add on
	fmt.Println(newNames)                     // [Mary John Luke Jane]
	fmt.Println(names)                        // names is still contain old values just two names  => [Mary John]
}
