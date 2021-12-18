package main

import (
	"fmt"
)
//los channels son mas eficientes en cuanto a tiempo de desarrollo, si se necesita mucha optimización usar la forma primitiva

//buena prac: cuando hay canales en parametros indicar si son de entrada o salida
func say(text string, c chan<- string) {
	c <- text
}

func main() {
	//1 indica el limite de datos simultaneos, si no lo pones es dinamico, pero mala practica
	c := make(chan string, 1)

	fmt.Println("Hello")

	go say("Bye", c)
	
	//se extrae el dato que se ingreso en say para que la goroutine main espere el "Bye"
	fmt.Println(<-c)
}
