package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const (
	csvFile  = "problems.csv"
	interval = 30
)

type Problem struct {
	question string
	answer   string
}

func main() {

	quizFile := flag.String("file", csvFile, "Quiz file")
	quizInterval := flag.Int("interval", interval, "Quiz interval")
	flag.Parse()
	file := *quizFile
	interv := *quizInterval

	_, err := os.Stat(file)
	if err != nil {
		CreateCSVQuizFile(file)
	}

	problems := getProblems(file)
	totalProblems := len(problems)
	correctAnswers := 0
	wrongAnswers := 0
	timer := time.NewTimer(time.Duration(interv) * time.Second)
	go func() {
		<-timer.C
		fmt.Println("Time is up!")
		fmt.Printf("Total problems: %d\n", totalProblems)
		fmt.Printf("Correct answers: %d\n", correctAnswers)
		fmt.Printf("Wrong answers: %d\n", wrongAnswers)
		os.Exit(0)
	}()

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if strings.TrimSpace(answer) == problem.answer {
			correctAnswers++
		} else {
			wrongAnswers++
		}

	}

}

func getProblems(filename string) []Problem {
	problems := []Problem{}
	file, err := os.Open(filename)
	if err != nil {
		return problems
	}
	defer file.Close()
	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		problems = append(problems, Problem{
			question: record[0],
			answer:   record[1],
		})
	}
	return problems
}
