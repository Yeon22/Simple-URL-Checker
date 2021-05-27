package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	var url string
	fmt.Print("What URL do you want check? : ")
	fmt.Scanln(&url)
	err := hitURL(url)
	fmt.Println(err)
}

func hitURL(url string) error {
	fmt.Println("Checking URL: ", url)

	resp, err := http.Get(url)

	if err != nil || resp.StatusCode >= 400 {
		fmt.Println("hitURL error and status code : ", err, resp.StatusCode)
		return errRequestFailed
	}

	return nil
}
