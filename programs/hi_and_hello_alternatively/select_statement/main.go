package main

import (
	"fmt"
	"sync"
	"time"
)

func printHi(hiCh, helloCh, done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-hiCh:
			fmt.Println("Hi")
			helloCh <- true
		case <-done:
			fmt.Println("Recived Done 1")
			return
		}
	}
}

func printHello(hiCh, helloCh, done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-helloCh:
			fmt.Println("Hello")
			hiCh <- true
		case <-done:
			fmt.Println("Recived Done 2")
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup

	hiCh := make(chan bool, 1)
	helloCh := make(chan bool, 1)
	done := make(chan bool)

	wg.Add(2)
	go printHi(hiCh, helloCh, done, &wg)
	go printHello(hiCh, helloCh, done, &wg)

	hiCh <- true
	time.Sleep(time.Millisecond)
	close(done)

	wg.Wait()
}
