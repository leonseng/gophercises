package main

import (
	"cyoa/util"
	"log"
	"net/http"
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

	// http.HandleFunc("/", util.StoryHandleFunc(story))
	// log.Fatal(http.ListenAndServe(":8080", nil))

	log.Fatal(http.ListenAndServe(":8080", util.StoryHandler{story}))
}
