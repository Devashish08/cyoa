package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	cyon "github.com/Devashish08/cyoa"
)

func main() {
	fileName := flag.String("file", "story.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *fileName)

	f, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(f)
	var story cyon.Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", story)
}
