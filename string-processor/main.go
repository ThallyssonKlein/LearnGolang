package main

import (
	"sync"
	"fmt"
)

var jobs = make(chan string, 10)

func producer() {
	go func() {
		for i := 0; i < 100; i++ {
			jobs <- "Text"
		}

		close(jobs)
	}()
}

var results = make(chan string, 10)
var wg sync.WaitGroup

func consumer() {
	go func(){
		wg.Wait()
		close(results)
	}()

	wg.Add(10)
	for i := 0; i< 10; i ++ {
		go func(){
			defer wg.Done()
			for job := range jobs {
				results <- job + " 2"
			}
		}()
	}
}


func main() {
	producer()
	consumer()

	for res := range results {
		fmt.Println(res)
	}
}