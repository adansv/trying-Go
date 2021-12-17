package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup
	

	fmt.Println("Holis")
	wg.Add(1)
	go say("World", &wg)

	wg.Wait()

	go func(text string){
		fmt.Println(text)
	}("Adiós")
	time.Sleep(time.Second * 3)
}
