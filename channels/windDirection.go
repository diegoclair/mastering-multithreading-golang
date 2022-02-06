package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/*
	Time before channels = Processing took:  3.1040076s
	Time after  channels = Processing took:  1.9708051s
*/

var (
	windRegex     = regexp.MustCompile(`\d* METAR.*EGLL \d*Z [A-Z ]*(\d{5}KT|VRB\d{2}KT).*=`)
	tafValidation = regexp.MustCompile(`.*TAF.*`)
	comment       = regexp.MustCompile(`\w*#.*`)
	metarClose    = regexp.MustCompile(`.*=`)
	variableWind  = regexp.MustCompile(`.*VRB\d{2}KT`)
	validWind     = regexp.MustCompile(`\d{5}KT`)
	windDirOnly   = regexp.MustCompile(`(\d{3})\d{2}KT`)
	windDist      [8]int
)

func parseToArray(textChannel chan string, metarChannel chan []string) {
	//text := <-textChannel  we could use this, but we need to know when the channel is close, so we use a for loop

	for text := range textChannel {
		lines := strings.Split(text, "\n")
		metarSlice := make([]string, 0, len(lines))
		metarStr := ""

		for _, line := range lines {
			if tafValidation.MatchString(line) {
				break
			}
			if !comment.MatchString(line) {
				metarStr += strings.Trim(line, " ")
			}
			if metarClose.MatchString(line) {
				metarSlice = append(metarSlice, metarStr)
				metarStr = ""
			}
		}
		metarChannel <- metarSlice
	}
	//when the main trhead closes the textChannel, then it will get out of the for loop above and so we can close the metarChannel
	close(metarChannel)
}

func extractWindDirection(metarChannel chan []string, windsChannel chan []string) {

	for metars := range metarChannel {
		winds := make([]string, 0, len(metars))
		for _, metar := range metars {
			if windRegex.MatchString(metar) {
				winds = append(winds, windRegex.FindAllStringSubmatch(metar, -1)[0][1])
			}
		}
		windsChannel <- winds
	}
	close(windsChannel)

}

func mineWindDistribution(windsChannel chan []string, distChannel chan [8]int) {

	for winds := range windsChannel {
		for _, wind := range winds {
			if variableWind.MatchString(wind) {
				for i := 0; i < 8; i++ {
					windDist[i]++
				}
			} else if validWind.MatchString(wind) {
				windStr := windDirOnly.FindAllStringSubmatch(wind, -1)[0][1]
				if d, err := strconv.ParseFloat(windStr, 64); err == nil {
					dirIndex := int(math.Round(d/45.0)) % 8
					windDist[dirIndex]++
				}
			}
		}
	}
	distChannel <- windDist
	close(distChannel)
}

func main() {

	textChannel := make(chan string)
	metarChannel := make(chan []string)
	windsChannel := make(chan []string)
	resultsChannel := make(chan [8]int)

	//the output channel (second parameter) is the input for the next function

	// 1. Change to array, each metar report (line) is a separate item in the array
	go parseToArray(textChannel, metarChannel)

	// 2. Extract wind direction, EGLL 312350Z 07004KT CAVOK 12/09 Q1016 NOSIG= -> 070
	go extractWindDirection(metarChannel, windsChannel)

	// 3. Assign to [N, NE, E, SE, S, SW, W, NW], 070 -> E + 1
	go mineWindDistribution(windsChannel, resultsChannel)

	absPath, _ := filepath.Abs("./metarfiles/")
	files, _ := ioutil.ReadDir(absPath)
	start := time.Now()

	for _, file := range files {
		dat, err := ioutil.ReadFile(filepath.Join(absPath, file.Name()))
		if err != nil {
			panic(err)
		}
		text := string(dat)
		textChannel <- text
	}
	close(textChannel) //it will tell no another trhead that the process is done
	results := <-resultsChannel

	fmt.Println(results)
	fmt.Println("Processing took: ", time.Since(start).Seconds())
}
