package main

import (
	"fmt"
	"sync"
)

var message string
var wg sync.WaitGroup

func updateMessage(msg string, m *sync.Mutex) {
	defer wg.Done()

	m.Lock()
	message = msg
	m.Unlock()
}

func main() {
	message = "hello world"

	var mutex sync.Mutex
	wg.Add(2)
	go updateMessage("hello universe", &mutex)
	go updateMessage("hello cosmos", &mutex)
	wg.Wait()

	fmt.Println(message)

}
