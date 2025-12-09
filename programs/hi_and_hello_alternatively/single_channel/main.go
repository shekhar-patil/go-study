package main

import (
	"fmt"
	"sync"
)

func printHello(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		<-ch
		fmt.Println("Hello")
		ch <- true
	}
}

func printHi(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Println("Hi")
		ch <- true
		<-ch
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan bool)

	wg.Add(2)
	go printHello(ch, &wg)
	go printHi(ch, &wg)

	wg.Wait()
}
