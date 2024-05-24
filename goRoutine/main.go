package main

import (
	"fmt"
	"sync"
)

func printSomething(str string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(str)
}

// func main() {
// 	go printSomething("hello world 1")
// 	printSomething("hello world 2")
// 	time.Sleep(1 * time.Second)
// }

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go printSomething("hello world 1", &wg)
	wg.Wait()

	wg.Add(1)
	printSomething("hello world 2", &wg)
	// time.Sleep(1 * time.Second)
}


