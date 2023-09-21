package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
)

func CreateCSVQuizFile(fname string) {
	records := [][]string{}
	for i := 0; i < 12; i++ {
		a := rand.Intn(100)
		b := rand.Intn(100)
		s := a + b
		question := fmt.Sprintf("%d+%d", a, b)
		answer := fmt.Sprintf("%d", s)
		records = append(records, []string{question, answer})
	}
	file, err := os.Create(fname)
	if err != nil {
		fmt.Println(err)
	}
	w := csv.NewWriter(file)
	w.WriteAll(records)
}
