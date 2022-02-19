package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

/*
	Shoelace algorithm explained:
		Excellent youtube Video going into a bit more detail why the algorithm works:
		https://www.youtube.com/watch?v=0KjG8Pg6LGk

		Article on brilliant showing the proof of the of the shoelace:
		https://brilliant.org/wiki/triangles-calculating-area/
*/

/*
	avarage time before thread pools ----> The processing took 6.2891032s
	avarage time after  thread pools ----> The processing took 4.8891032s
*/

var (
	numberOfThreads int = runtime.NumCPU()
	waitGroup           = sync.WaitGroup{}
)

type Point2D struct {
	x int
	y int
}

func main() {

	absPath, _ := filepath.Abs("./")
	dat, _ := ioutil.ReadFile(filepath.Join(absPath, "polygons.txt"))
	text := string(dat)

	inputChannel := make(chan string, 1000) //we can put 1000 values in the channel, like a queue
	for i := 0; i < numberOfThreads; i++ {
		go findArea(inputChannel)
	}
	waitGroup.Add(numberOfThreads)

	start := time.Now()
	for _, line := range strings.Split(text, "\n") {
		inputChannel <- line
	}
	close(inputChannel)
	waitGroup.Wait() //when we close the channel, the other threads my have not finished to process the queue of inputchannel.. so we can use the waitGroup to synchronize

	fmt.Println(`The processing took`, time.Since(start))

}

var (
	r = regexp.MustCompile(`\((\d*),(\d*)\)`)
)

func findArea(inputChannel chan string) {
	for pointsStr := range inputChannel {
		var points []Point2D
		matches := r.FindAllStringSubmatch(pointsStr, -1)

		for _, p := range matches {
			x, _ := strconv.Atoi(p[1])
			y, _ := strconv.Atoi(p[2])
			points = append(points, Point2D{x: x, y: y})
		}

		area := 0.0
		pointsLenght := len(points)
		for i := 0; i < pointsLenght; i++ {
			a, b := points[i], points[(i+1)%pointsLenght]
			area += float64(a.x*b.y) - float64(a.y*b.x)
		}

		fmt.Println(math.Abs(area) / 2.0)
	}
	waitGroup.Done()
}
