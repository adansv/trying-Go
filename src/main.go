package main

import (
	"fmt"
	"strings"
)

/*
func normalFuction(message string) {
	fmt.Println(message)
}
func tripleArgument(a int, b int, c string) {
	fmt.Println(a, b, c)
}
func doubleReturn(a int) (c, d int) {
	return a, a * 2
}
func isPalindromo(text string) {
	string b
	textReverse += string(text[i])
	for i := len(a) - 1; i >= 0; i++ {

	}
}
*/
func isPalindromo(wo string) {
	wo = strings.ToLower(wo)
	fmt.Println(wo)
	cond := true
	//fmt.Println("this: ", string((wo)[2]))
	for i := 0; i < len(wo); i++ {
		aux := wo

		if wo[i] != aux[(len(wo)-i)-1] {
			cond = false
			break
		}

	}
	fmt.Println("this: ", cond)
}

func main() {

	slice := []string{"hola", "que", "haces"}

	for i := range slice {
		fmt.Println(i)
	}
	//palidromos

	var wo = "oEaraeo"

	isPalindromo(wo)

}
