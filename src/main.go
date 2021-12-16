package main

import "fmt"

type car struct {
	brand string
	year int
}

func main() {
	mycar := car{brand:"ford", year: 2020}

	fmt.Println(mycar)

	var otherCar car
	otherCar.brand = "mercedes"
	fmt.Println(otherCar)
}
