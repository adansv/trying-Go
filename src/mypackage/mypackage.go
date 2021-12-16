package mypackage

import "fmt"

//CarPublic Car con acceso público
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

// print message
func PrintMessage(tx string){
	fmt.Println(tx)
}
// private cant be called
func printMessage(tx string){
	fmt.Println(tx)
}
