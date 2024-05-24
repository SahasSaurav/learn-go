package main

import (
	"fmt"
	"sync"
)

var message string

func printMessage() {
	fmt.Println(message)
}

func updateMessage(msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	message = msg
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go updateMessage("hello universe", &wg)
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("hello cosmos", &wg)
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("hello world", &wg)
	wg.Wait()
	printMessage()

}
