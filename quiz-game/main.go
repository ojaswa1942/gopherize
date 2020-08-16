package main

import (
	"fmt"
	"flag"
	"os"
	"encoding/csv"
	"strings"
)

type problem struct {
	ques string
	ans string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "csv file in format of `question,answer`")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Error opening file: %s\n", *csvFilename))
	}
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse CSV")
	}

	problems := parseQuestions(lines)
	score := 0

	for i, problem := range problems {
		fmt.Printf("Question #%d: %s\n", i+1, problem.ques)
		var answer string
		// scanf handles whitespace trimming
		fmt.Scanf("%s\n", &answer)
		if answer == problem.ans {
			score++
		}
	}

	fmt.Printf("Your score: %d/%d\n", score, len(problems))

	fmt.Printf("There you go: %s\n", *csvFilename)
	fmt.Println(problems)
}

func parseQuestions(questions [][]string) []problem {
	parsed := make([]problem, len(questions))
	for i, question := range questions {
		parsed[i] = problem {
			ques: question[0],
			ans: strings.TrimSpace(question[1]),
		}
	}
	return parsed
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}