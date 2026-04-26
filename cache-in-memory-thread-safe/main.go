package main

import (
	"sync"
	"math/rand"
)

type Cache struct {
	mu	sync.RWMutex
	data map[string]string
}

func (c *Cache) Set(key string, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *Cache) Get(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data[key]
}

func main(){
	c := &Cache {
		data: make(map[string]string),
	}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		random := rand.Intn(3) + 1
		switch random {
		case 1:
			wg.Add(1)
			go func () {
				defer wg.Done()
				c.Set("chave1", "value")
			}()
		case 2:
			wg.Add(1)
			go func () {
				defer wg.Done()
				c.Set("chave2", "value")
			}()
		case 3:
			wg.Add(1)
			go func () {
				defer wg.Done()
				c.Set("chave3", "value")
			}()
		}
	}

	for i := 0; i < 1000; i++ {
		random := rand.Intn(3) + 1
		switch random {
		case 1:
			wg.Add(1)
			go func () {
				defer wg.Done()
				c.Get("chave1")
			}()
		case 2:
			wg.Add(1)
			go func () {
				defer wg.Done()
				c.Get("chave2")
			}()
		case 3:
			wg.Add(1)
			go func () {
				defer wg.Done()
				c.Get("chave3")
			}()
		}
	}

	wg.Wait()
}