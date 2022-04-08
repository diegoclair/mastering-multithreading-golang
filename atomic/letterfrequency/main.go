package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

/*
	time used before go routines        -> 53.791521271s     // without extra loop of 20 times
	time used with mutex mode      		-> 21.077837301s
	time used with atomic mode     		-> 3.587836763s
*/

const (
	allLetters = "abcdefghijklmnopqrstuvwxyz"
	atomicMode = 1
	mutexMode  = 2
)

const processWith = mutexMode

func main() {
	var url string
	var frequency [26]int32
	wg := sync.WaitGroup{}

	start := time.Now()
	for i := 1000; i <= 1200; i++ {
		wg.Add(1)
		url = fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i)
		if processWith == atomicMode {
			go countLettersWithAtomic(url, &frequency, &wg)
		} else {
			go countLettersWithMutex(url, &frequency, &wg)
		}
	}
	wg.Wait()
	jobTime := time.Since(start)
	fmt.Printf("Process took %s\n", jobTime)
	fmt.Println("Done")
	for i, f := range frequency {
		fmt.Printf("%s -> %d\n", string(allLetters[i]), f)
	}
}

var lock = sync.Mutex{}

func countLettersWithMutex(url string, frequency *[26]int32, wg *sync.WaitGroup) {
	body := getBody(url)

	for i := 0; i <= 20; i++ {
		for _, b := range body {
			c := strings.ToLower(string(b))
			lock.Lock()
			alphabetPosition := strings.Index(allLetters, c)
			if alphabetPosition >= 0 {
				frequency[alphabetPosition] += 1
			}
			lock.Unlock()
		}
	}
	wg.Done()
}

func countLettersWithAtomic(url string, frequency *[26]int32, wg *sync.WaitGroup) {
	body := getBody(url)

	for i := 0; i <= 20; i++ {
		for _, b := range body {
			c := strings.ToLower(string(b))
			alphabetPosition := strings.Index(allLetters, c)
			if alphabetPosition >= 0 {
				atomic.AddInt32(&frequency[alphabetPosition], 1)
			}
		}
	}
	wg.Done()
}

func getBody(url string) []byte {

	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
