package main

import (
	pk "curso_golang_platzi/src/mypackage"
	"fmt"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "ferrari"
	myCar.Year = 2015
	fmt.Println(myCar)

	pk.PrintMessage("Hola tarola")

}
