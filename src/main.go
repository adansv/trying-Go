package main

import (
	"fmt"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

func (mypc pc) String() string {
	return fmt.Sprintf("Tengo %d GB de RAM, %d GB en disk y es una %s", mypc.ram, mypc.disk, mypc.brand)
}

func main() {

	mypc := pc{ram: 16, disk: 500, brand: "msi"}

	fmt.Println(mypc)

	fmt.Println("Updated", mypc)

}
