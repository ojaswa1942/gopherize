package main

import (
	"fmt"
	"flag"
	"os"
	"encoding/csv"
	"strings"
	"time"
)

type problem struct {
	ques string
	ans string
}

func main() {
	// Parse flags
	csvFilename := flag.String("csv", "problems.csv", "csv file in format of `question,answer`")
	timeLimit := flag.Int("limit", 10, "quiz time limit in seconds")
	flag.Parse()

	// Read file & parse questions
	lines := readCSV(csvFilename)
	problems := parseQuestions(lines)

	// Init timer & begin quiz
	timer := time.NewTimer( time.Duration(*timeLimit) * time.Second)
	score := 0

	for i, problem := range problems {
		fmt.Printf("Question #%d: %s? ", i+1, problem.ques)
		userAnswer := getNonBlockingInput()

		select {
			case <-timer.C:
				fmt.Printf("\n\nTortoise alert: Time up!\n")
				exit(fmt.Sprintf("Your score: %d/%d", score, len(problems)), false)
			case answer := <-userAnswer :
				if strings.ToLower(answer) == problem.ans {
					score++
				}
		}
	}

	exit(fmt.Sprintf("\nYour score: %d/%d", score, len(problems)), false)
}

func getNonBlockingInput() (chan string) {
	answerChannel := make(chan string)
	go func () {
		var answer string
		fmt.Scanf("%s\n", &answer)
		answerChannel <- answer
	}()
	return answerChannel	
}

func readCSV(csvFilename *string) [][]string {
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Error opening file: %s\n", *csvFilename), true)
	}
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse CSV", true)
	}
	return lines
}

func parseQuestions(questions [][]string) []problem {
	parsed := make([]problem, len(questions))
	for i, question := range questions {
		parsed[i] = problem {
			ques: question[0],
			ans: strings.ToLower(strings.TrimSpace(question[1])),
		}
	}
	return parsed
}

func exit(msg string, error bool) {
	fmt.Println(msg)
	if error {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}