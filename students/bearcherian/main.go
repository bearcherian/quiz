package main

import (
	"flag"
	"os"
	"encoding/csv"
	"fmt"
	"bufio"
	"strings"
	"time"
)

var file = flag.String("file", "problems.csv", "a CSV file with problems and answers")
var timeLimit = flag.Int("time", 30, "the time limit in seconds to complete the quiz")
var correct = 0
var incorrect = 0

func main() {

	flag.Parse()
	
	// check if file flag
	fmt.Printf("reading from %s. You have %ds to answer all questions\n",*file,*timeLimit)

	// read CSV
	problems, err := readFile(*file)
	if err != nil {
		fmt.Println("unable to read file" + err.Error())
		return
	}

	go startTimer(time.Duration(*timeLimit) * time.Second)

	inReader := bufio.NewReader(os.Stdin)
	for _, p := range problems {
		question := p[0]
		answer := p[1]

		fmt.Println(question)

		response, _ := inReader.ReadString('\n')

		if strings.TrimSpace(response) == answer {
			fmt.Println("Correct!")
			correct++
		} else {
			fmt.Printf("Incorrect. The correct answer is %s\n", answer)
			incorrect++
		}
	}

	exit()
}

func startTimer(d time.Duration) {
	timer := time.NewTimer(d)

	<-timer.C
	fmt.Println("Game Over. Time Exceeded")
	exit()
}
func exit() {
	fmt.Printf("You got %d correct and %d wrong.\n", correct, incorrect)
	os.Exit(0)
}

func readFile(file string) ([][]string, error) {
	reader, err := os.Open(file)
	if err != nil {
		fmt.Println("unable to read file" + err.Error())
		return nil, err
	}

	csvReader := csv.NewReader(reader)

	return csvReader.ReadAll()
}

