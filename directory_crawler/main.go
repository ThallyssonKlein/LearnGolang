package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"

	// "strings"
	"sync"
	"syscall"
)

type Result struct {
	Path string
	Hash string
	Err error
}

func readDir(ctx context.Context, dir string, results chan <- Result, wg *sync.WaitGroup, spotlight chan struct{}) {
	if err := ctx.Err(); err != nil {
		return
	}

	defer wg.Done()
	// if strings.Contains(dir, "git") {
	// 	return
	// }

	entries, err := os.ReadDir(dir)

	if err != nil {
		results <- Result{Path: dir, Err: err}
		return
	}

	for _, entry := range entries {
		name := entry.Name()

		if entry.IsDir() {
			wg.Add(1)
			go func (n string) {
				fullPath := filepath.Join(dir, name)
				readDir(ctx, fullPath, results, wg, spotlight)
			}(name)
		} else {
			wg.Add(1)
			go func (n string) {
				defer wg.Done()

				fullPath := filepath.Join(dir, name)
				hash, err := generateHash(ctx, fullPath, spotlight)

				if err != nil {
					results <- Result{Path: fullPath, Err: err}
				} else {
					results <- Result{Path: fullPath, Hash: hash}
				}
			}(name)
		}
	}
}

func generateHash(ctx context.Context, path string, spotlight chan struct{}) (string, error) {
	select {
		case <- ctx.Done():
			return "", nil
		case spotlight <- struct{}{}:
			defer func () {
				<- spotlight
			}()
			f, err := os.Open(path)

			if err != nil {
				return "", err
			}
			defer f.Close()

			h := md5.New()

			if _, err := io.Copy(h, f); err != nil {
				return "", err
			}

			hashBytes := h.Sum(nil)

			return fmt.Sprintf("%x", hashBytes), nil
	}
}


func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	results := make(chan Result)
	
	var limit = 0
	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		limit = 100
	} else {
		limit = int(rLimit.Cur) / 2
	}
	const maxVacancies = 100
	if limit > maxVacancies {
		limit = maxVacancies
	}

	stoplight := make(chan struct{}, limit)
	
	var wg sync.WaitGroup

	wg.Add(1)
	go readDir(ctx, "../", results, &wg, stoplight)
	
	go func () {
		wg.Wait()
		close(results)
	}()

	sum := 0
	for res := range results {
		if res.Err != nil {
			fmt.Printf("Error at %s: %v\n", res.Path, res.Err)
		} else {
			fmt.Printf("File: %s | Hash %s\n", res.Path, res.Hash)
			sum += 1
		}
	}

	fmt.Printf("Sum: %d\n", sum)
}
