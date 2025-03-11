package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	cyon "github.com/Devashish08/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	fileName := flag.String("file", "story.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *fileName)

	f, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	story, err := cyon.JsonStory(f)
	if err != nil {
		panic(err)
	}

	h := cyon.NewHandler(story, nil)
	fmt.Printf("Starting the server at:%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

}
