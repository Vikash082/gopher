package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var csvFile = flag.String("file", "abc.csv", "path of the csv file")
	var timeLimit = flag.Int("time", 5, "timelimit of quiz in sec")
	flag.Parse()
	var _ = csvFile
	cFile, err := os.Open(*csvFile)
	if err != nil {
		panic(err)
	}
	cReader := csv.NewReader(cFile)
	records, err := cReader.ReadAll()
	if err != nil {
		panic(err)
	}
	duration := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	var que, correct int
	inputCh := make(chan string)
loop:
	for _, r := range records {
		fmt.Println("What is ", r[0], " sir ?")
		que++

		go func() {
			// input := bufio.NewReader(os.Stdin)
			// inStr, err := input.ReadString('\n')
			var input string
			fmt.Scanf("%s\n", &input)
			if err != nil {
				os.Exit(1)
			}
			inputCh <- input
		}()
		select {
		case answer := <-inputCh:
			if r[1] == strings.TrimSpace(strings.TrimRight(answer, "\n")) {
				correct++
			}
		case <-duration.C:
			fmt.Println()
			break loop
		}
	}

	fmt.Println("Total Question: ", que, " Total correct answer: ", correct)
}
