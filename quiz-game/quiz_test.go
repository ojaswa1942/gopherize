package main

import (
	"flag"
	"testing"
	"os"
)

var csvFilename = flag.String("csv", "problems.csv", "csv file in format of `question,answer`")

func TestMain(m *testing.M) {
	flag.Parse()
	exitResponse := m.Run()
	os.Exit(exitResponse)
}

func TestCSV(t *testing.T) {
	// t.Run("test ReadCSV", testReadCSV)
	t.Log("using default file problems.csv to conduct test")
	lines := readCSV(csvFilename)

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
		{ []string{"2", "wor", "44"}, 0 },
		{ []string{"4", "world", "44", "11", "yahoo"}, 3 },
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

func TestStringSanitize(t *testing.T) {
	cases := []struct{
		input, output string
	}{
		{ "", "" },
		{ "  haha", "haha" },
		{ "  ha ha", "ha ha" },
		{ "  ha ha   ", "ha ha" },
		{ "  haha   ", "haha" },
		{ "  h a YAYA h a   ", "h a yaya h a" },
		{ "  CAPS", "caps" },
		{ "  CAPSYaAa", "capsyaaa" },
		{ "  CAPS 69  ", "caps 69" },
	}

	t.Log("initiate testing trim and caps cases")
	for i, testcase := range cases {
		if generated, expected := sanitizeString(testcase.input), testcase.output; generated != expected {
			t.Errorf("case #%d failed. expected: `%s`, received: `%s`", i+1, expected, generated)
		}
	}
}
