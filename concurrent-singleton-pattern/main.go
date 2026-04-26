package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type Database struct {
	Connection string
}

var instance *Database

func GetInstance() *Database {
	once.Do(func() {
		if instance == nil {
			instance = &Database{Connection: "Connected"}
		}
	})

	fmt.Printf("%p\n", instance)

	return instance
}

func main() {
	for i := 0; i < 100; i++ {
		go GetInstance()
	}
}