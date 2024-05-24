package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func getWebsiteStatus(ctx context.Context, url string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error in creating for request", err)
	}
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		ch <- fmt.Sprintf("%v is down", url)
		fmt.Println("Error in fetching the data", res)
	}
	defer res.Body.Close()
	ch <- fmt.Sprintf("%v is up", url)
	fmt.Printf("url %v and site status is %d\n", url, res.StatusCode)
}

func main() {
	urls := []string{
		"http://google.com",
		"http://youtube.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
	}

	statusChan := make(chan string)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	wg := &sync.WaitGroup{}
	// mutex := &sync.Mutex{}

	wg.Add(len(urls))

	for _, url := range urls {
		go getWebsiteStatus(ctx, url, statusChan, wg)
	}

	for {
		chanVal, ok := <-statusChan
		if !ok {
			break
		}
		fmt.Println(chanVal)
	}
	defer close(statusChan)
	wg.Wait()

}
