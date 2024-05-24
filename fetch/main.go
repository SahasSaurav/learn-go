package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type NewPost struct {
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func GetAllPost() []Post {
	req, err := http.NewRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/posts", nil)
	handleError(err)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Do(req)
	handleError(err)
	defer res.Body.Close()

	var posts []Post
	err = json.NewDecoder(res.Body).Decode(&posts)
	handleError(err)

	return posts

}

func createPost(newPost NewPost) Post {
	jsonPayload, err := json.Marshal(newPost)
	handleError(err)
	payload := bytes.NewReader(jsonPayload)

	req, err := http.NewRequest(http.MethodPost, "https://jsonplaceholder.typicode.com/posts", payload)
	handleError(err)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	res, err := client.Do(req)
	handleError(err)
	defer res.Body.Close()

	var post Post
	err = json.NewDecoder(res.Body).Decode(&post)
	handleError(err)

	return post
}

func main() {
	posts := GetAllPost()
	for _, post := range posts {
		fmt.Printf("%+v\n", post)
	}

	newPost := NewPost{
		Body:   "hello world ",
		UserId: 10,
		Title:  "post method Fetch",
	}

	post := createPost(newPost)
	fmt.Printf("%+v", post)
}
