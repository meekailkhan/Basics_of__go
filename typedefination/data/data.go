package data

import "fmt"

type distance float64 // miles
type distanceKm float64

func (miles distance) ToKm() distanceKm {
	return distanceKm(1.60934 * miles)
}

func (km distanceKm) ToMiles() distance {
	return distance(km / 1.60934)
}

func Test() {
	d := distance(4.5)
	km := d.ToKm()
	newD := km.ToMiles()

	fmt.Println("distance in miles", d)
	fmt.Println("distance in kilometers", km)
	fmt.Println("new distance in miles", newD)
}
