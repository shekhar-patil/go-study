/*
Write a Go program that spawns two goroutines. Each goroutine should increment a shared counter variable,
but only one goroutine should increment it by 1, and the other should increment it by 2.
The two goroutines should alternate in their operations, ensuring the increments are performed in a synchronized manner,
and the final value of the counter should be consistent.

Ex:
1  (increment by 1)
3  (increment by 2)
4  (increment by 1)
6  (increment by 2)
7  (increment by 1)
Final counter value: 7
*/

package main

import (
	"fmt"
	"sync"
)

func incrementByOne(counter *int, k int, one, two chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(two)

	for range one {
		if *counter+1 > k {
			return
		}

		*counter = *counter + 1
		fmt.Println("Current Counter inc by one", *counter)
		two <- true
	}
}

func incrementByTwo(counter *int, k int, one, two chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(one)
	for range two {
		if *counter+2 > k {
			return
		}

		*counter = *counter + 2
		fmt.Println("Current Counter inc by two", *counter)
		one <- true
	}
}

func main() {
	counter := 0
	k := 51
	one := make(chan bool)
	two := make(chan bool)

	var wg sync.WaitGroup

	wg.Add(2)
	go incrementByOne(&counter, k, one, two, &wg)
	go incrementByTwo(&counter, k, one, two, &wg)

	one <- true

	wg.Wait()
	fmt.Println("Final Counter is", counter)
}
