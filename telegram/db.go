package telegram

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "333btybfRA"
	dbname   = "LunchPollBotDB"
)

var dbInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

var cafe string
var cafepassword string
var cafename []string
var cafepass []string

var admin string
var adminpassword string

func getCafe() error {
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT cafe, password FROM admincafe ")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&cafe, &cafepassword)
		if err != nil {
			panic(err)
		}
		cafename = append(cafename, cafe)
		cafepass = append(cafepass, cafepassword)

	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return nil
}

func getCompany() error {

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT name, password FROM adminCompany")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&admin, &adminpassword)
		if err != nil {
			panic(err)
		}

	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return nil
}

func createMenu(f []string, s []string, a []string, d []string) error {
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}
	firstDish := strings.Join(f, ",")
	secondDish := strings.Join(s, ",")
	apitizer := strings.Join(a, ",")
	drinks := strings.Join(d, ", ")
	defer db.Close()
	data := `INSERT INTO Menu(cafe, firstDish, secondDish, apitizer, drinks) VALUES ($1, $2, $3, $4, $5);`
	if err != nil {
		return err
	}
	_, err = db.Exec(data, cafeName, firstDish, secondDish, apitizer, drinks)
	if err != nil {
		return err
	}

	return nil
}
