package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	myChan := make(chan int)

	wg.Add(2)

	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			val, ok := <-ch
			if !ok {
				break
			}
			fmt.Println(val)
		}
	}(myChan, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}(myChan, wg)

	wg.Wait()
}
