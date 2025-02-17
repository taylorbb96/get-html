package main

import (
	"fmt"
	"net/http"
	"io"
)

func get_response(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return string(body)
}

func main() {
	var input_url string

	fmt.Print("Input a URL: ")
	fmt.Scan(&input_url)

	response := get_response(input_url)
	
	fmt.Println("Response: ")
	fmt.Print(response)
}
