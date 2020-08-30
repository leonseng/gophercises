package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

var rightCount int
var aChannel = make(chan string)
var wg = sync.WaitGroup{}
var stdinReader = bufio.NewReader(os.Stdin)

func main() {
	var gameTime int
	var fileName string

	flag.IntVar(&gameTime, "t", 30, "Quiz game timer in seconds")
	flag.StringVar(&fileName, "f", "problems.csv", "CSV file containing problems and answers")
	flag.Parse()

	problems := loadProblems(fileName)

	fmt.Printf("%v seconds to answer %v questions. Press Enter to start game...", gameTime, len(problems))
	stdinReader.ReadString('\n')

	wg.Add(1)  // only wait for one of the goroutines to finish
	go playGame(problems)
	go gameTimer(gameTime)
	wg.Wait()
	
	fmt.Printf("Score: %v/%v\n", rightCount, len(problems))
}

func loadProblems(fileName string) [][]string {
	csvFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Couldn't open CSV file", err)
	}

	// parse CSV file
	lines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		panic(err)
	}

	return lines
}

func gameTimer(gameTime int) {
	time.Sleep(time.Duration(gameTime) * time.Second)
	fmt.Println("\n*** Time's up ***")
	wg.Done()
}

func playGame(problems [][]string) {
	for i, line := range problems {
		q, a := line[0], line[1]

		// present question and expect answer
		fmt.Printf("Q%v: %v? ", i+1, q)
		userA, _ := stdinReader.ReadString('\n')

		if strings.TrimSpace(userA) == a {
			rightCount++
		}
	}

	wg.Done()
}
