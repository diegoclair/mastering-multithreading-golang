package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

/*
	time used before atomic variable   -> 53.791521271s
*/

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency *[26]int32) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	for _, b := range body {
		c := strings.ToLower(string(b))
		alphabetPosition := strings.Index(allLetters, c)
		if alphabetPosition >= 0 {
			frequency[alphabetPosition] += 1
		}

	}
}

func main() {
	var url string
	var frequency [26]int32
	start := time.Now()
	for i := 1000; i <= 1200; i++ {
		url = fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i)
		countLetters(url, &frequency)
	}
	jobTime := time.Since(start)
	fmt.Printf("Process took %s\n", jobTime)
	fmt.Println("Done")
	for i, f := range frequency {
		fmt.Printf("%s -> %d\n", string(allLetters[i]), f)
	}
}
