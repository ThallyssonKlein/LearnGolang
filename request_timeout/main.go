package main

import (
	"context"
	"fmt"
	"time"
)

var apiResult = make(chan string)

func callApi() {
	time.Sleep(500 * time.Millisecond)
	apiResult <- "success"
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)

	defer cancel()

	go func () {
		callApi()
	}()

	select {
	case result := <- apiResult:
		fmt.Println("Success: " + result)
	case <- ctx.Done():
		fmt.Println("Timeout!")
	}
}