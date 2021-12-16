package main

import (
	"fmt"
)

type pc struct {
	ram   int
	disk  int
	brand string
}
func (mypc pc) ping(){
	fmt.Println(mypc.brand, "Pong")
}
func (mypc *pc) duplicateRam(){
	mypc.ram = mypc.ram*2
}


func main() {
	a := 50
	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	mypc := pc{ram:16, disk:500, brand:"msi"}

	fmt.Println(mypc)

	mypc.ping()

	fmt.Println(mypc)
	mypc.duplicateRam()
	fmt.Println("Updated", mypc)
	
	
}
