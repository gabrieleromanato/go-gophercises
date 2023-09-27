package main

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "password"
	dbname   = "phonebook"
)

type Number struct {
	ID     int    `json:"id"`
	Number string `json:"number"`
}

func Connect() error {
	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, dbname))
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

func normalizeNumbers() []string {
	normalizedNumbers := []string{}
	numbers, err := getNumbers()
	if err != nil {
		return normalizedNumbers
	}
	regex := regexp.MustCompile("[^0-9]+")
	for _, number := range numbers {
		norm := regex.ReplaceAllString(number.Number, "")
		normalizedNumbers = append(normalizedNumbers, norm)
	}
	return normalizedNumbers
}

func getNumbers() ([]Number, error) {
	var numbers []Number
	rows, err := db.Query("SELECT * FROM numbers")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var number Number
		if err := rows.Scan(&number.ID, &number.Number); err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return numbers, nil
}

func main() {
	if err := Connect(); err != nil {
		log.Fatal(err)
	}
	numbers := normalizeNumbers()
	fmt.Println(numbers)
}
