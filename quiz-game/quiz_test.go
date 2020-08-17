package main

import (
	// "strings"
	// "fmt"
	"testing"
	// "time"
)

// func testEachQuestion(t *testing.T) {
// 	timer := time.NewTimer(time.Duration(2) * time.Second).C
// 	done := make(chan string)
// 	var quest Question
// 	quest.question = "1+1"
// 	quest.answer = "2"
// 	var ans int
// 	var err error
// 	allDone := make(chan bool)
// 	go func() {
// 		ans, err = eachQuestion(quest.question, quest.answer, timer, done)
// 		allDone <- true
// 	}()
// 	done <- "2"

// 	<-allDone
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	assert.Equal(t, ans, 1)
// }

// func testReadCSV(t *testing.T) {
// 	str := "1+1,2\n2+1,3\n9+9,18\n"
// 	quest, err := readCSV(strings.NewReader(str))
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	var que [3]Question
// 	que[0].answer = "2"
// 	que[1].answer = "3"
// 	que[2].answer = "18"
// 	que[0].question = "1+1"
// 	que[1].question = "2+1"
// 	que[2].question = "9+9"

// 	assert.Equal(t, que[0], quest[0])
// 	assert.Equal(t, que[1], quest[1])
// 	assert.Equal(t, que[2], quest[2])

// }

// func TestEachQuestion(t *testing.T) {
// 	t.Run("test eachQuestion", testEachQuestion)
// }

func TestCSV(t *testing.T) {
	// t.Run("test ReadCSV", testReadCSV)
	t.Log("using default file problems.csv to conduct test")
	csvFilename := "problems.csv"
	lines := readCSV(&csvFilename)

	t.Run("check number of lines", func (t *testing.T) {
		if i := len(lines); i <= 0 {
			t.Error("problems.csv contains 0 lines")
		}
	})

	t.Log("parsing questions")
	problems := parseQuestions(lines)

	t.Run("assert number of problems", func (t *testing.T) {
		if i, j := len(lines), len(problems); i != j {
			t.Errorf("number of lines not equal to problems. Expected: %d, Found: %d", i, j)
		}
	})

	t.Run("assert questions/responses are non-empty", func(t *testing.T) {
		for i, problem := range(problems) {
			if problem.ques == "" {
				t.Errorf("found empty question: Question #%d with Ans: %v", i, problem.ans)
			}
			if problem.ans == "" {
				t.Errorf("found empty answer: Question #%d with Ques: %v", i, problem.ques)
			}
		}
	})
}

func TestScore(t *testing.T) {
	problems := []problem {
		{ "2+2", "4" },
		{ "hello", "world" },
		{ "21+12", "33" },
		{ "2+9", "11" },
		{ "ha+ha", "haha" },
		{ "5x5", "25" },
	}

	solutions := []struct{
		answers []string
		score int
	}{
		{ []string{"4", "world", "33", "11", "haha", "25"}, 6 },
		{ []string{"4", "world", "44"}, 2 },
		{ []string{"4", "world", "44"}, 2 },
		{ []string{"4", "world", "44"}, 2 },
	};

	t.Log("test all sets of solution")
	for i, solution := range solutions {
		score := 0
		for j, answer := range solution.answers {
			if sanitizeString(problems[j].ans) == sanitizeString(answer) {
				score++
			}
		}
		if score != solution.score {
			t.Errorf("score for solution #%d does not match. expected: %d, received: %d", i+1, solution.score, score)
		}
	}
}

func TestTimer(t *testing.T) {
}
