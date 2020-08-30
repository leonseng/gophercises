package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]

	fmt.Println(filename)
	csvFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Couldn't open CSV file", err)
	}

	// parse CSV file
	r := csv.NewReader(csvFile)
	var rightCount, wrongCount int

	stdinR := bufio.NewReader(os.Stdin)
	for {
		// read question and answer from CSV
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		q, a := record[0], record[1]

		// present question and expect answer
		fmt.Printf("Q: %v? ", q)

		userA, _ := stdinR.ReadString('\n')
		if strings.TrimSpace(userA) == a {
			rightCount++
		} else {
			wrongCount++
		}
	}

	fmt.Printf("Score: %v/%v\n", rightCount, rightCount+wrongCount)
}
