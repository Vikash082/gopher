package main

import (
	"chooseadventure"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fileName := flag.String("file", "gopher.json", "json file containing stories")
	flag.Args()
	gFile, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	gStory, err := chooseadventure.StoryDecoder(gFile)
	// fmt.Println("mainStory: ", gStory)
	storyHandler := chooseadventure.NewHandler(gStory)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 4002), storyHandler))
}
