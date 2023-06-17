package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PostResponse struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	// Make the http request
	resp, err := http.Get("https://jsonplaceholder.cypress.io/todos/11")
	if err != nil {
		print(err)
	}

	// Close the body
	defer resp.Body.Close()

	var post PostResponse

	// Decode the JSON response
	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		print(err)
	}

	// Print the result on the console
	fmt.Printf("UserId: %v\n", post.UserId)
	fmt.Printf("Id: %v\n", post.Id)
	fmt.Printf("Title: %v\n", post.Title)
	fmt.Printf("Completed: %v\n", post.Completed)
}
