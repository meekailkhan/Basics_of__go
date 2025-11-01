package main

import "fmt"

func calculateTax(price float32) float32 {
	return price * 0.2
}

func calculateDualTax(price float32) (float32, float32) {
	return price * 0.9, price * 0.2
}

func calculateTaxWithName(price float32) (cityTax float32, stateTax float32) {
	cityTax = price * 0.3
	stateTax = price * 0.2
	return

}

func main() {
	tax := calculateTax(100)
	fmt.Println(tax)

	cityTax, stateTax := calculateDualTax(100)
	fmt.Println("city tax is ", cityTax, "state tax is ", stateTax)

	_, stateTax = calculateTaxWithName(100)
	fmt.Println("state tax is ", stateTax)

}
