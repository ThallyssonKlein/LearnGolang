package main

import (
	"fmt"
	"sync"
)

type Database struct {
	Connection string
}

var instance *Database

func GetInstance() *Database {
	var once sync.Once
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