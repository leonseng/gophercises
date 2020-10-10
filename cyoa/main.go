package main

import (
	"cyoa/util"
	"fmt"
	"os"
)

func main() {
	jsonFile, err := os.Open("gopher.json")

	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	story, err := util.JsonStory(jsonFile)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", story)
}
