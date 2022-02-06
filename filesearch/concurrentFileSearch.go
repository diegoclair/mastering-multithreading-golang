package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	matches []string
	wg      = sync.WaitGroup{}
	lock    = sync.Mutex{}
)

func fileSearch(root string, filename string) {
	fmt.Println("Searching in", root)

	files, _ := ioutil.ReadDir(root)
	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			lock.Lock()
			matches = append(matches, filepath.Join(root, file.Name()))
			lock.Unlock()
		}
		if file.IsDir() {
			wg.Add(1)
			go fileSearch(filepath.Join(root, file.Name()), filename)
		}
	}
	wg.Done()
}

func main() {
	now := time.Now()
	wg.Add(1)
	go fileSearch("/", "README.md")

	wg.Wait()
	for _, file := range matches {
		fmt.Println("Matched", file)
	}
	fmt.Println("time: ", time.Since(now))
}
