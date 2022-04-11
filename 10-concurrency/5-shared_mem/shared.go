package main

import (
	"fmt"
	"sync"
)

// add прибавляет x к sum
func add(sum *int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10_000; i++ {
		*sum++
	}
}

func main() {
	sum := 0

	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go add(&sum, &wg)
	}
	wg.Wait()

	fmt.Println(sum) // Что будет на экране?
}
