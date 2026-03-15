package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type Entry1 struct {
	Number int
	Double int
	Square int
}

var DATA []Entry1
var tFile string

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Host: %s Path: %s\n", r.Host, r.URL.Path)
	fmt.Println("template file: ", tFile)
	myT := template.Must(template.ParseGlob(tFile))
	err := myT.ExecuteTemplate(w, filepath.Base(tFile), DATA)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Println("Need database File + Template File!")
		return
	}
	database := arguments[1]
	tFile = arguments[2]

	db, err := sql.Open("sqlite3", database)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Emptying database table.")
	_, err = db.Exec("DELETE FROM data")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Populating", database)

	stmt, _ := db.Prepare("INSERT INTO data(number, double, square) values(?, ?,?)")
	for i := 20; i < 50; i++ {
		_, _ = stmt.Exec(i, 2*i, i*i)
	}

	rows, err := db.Query("SELECT * FROM data")
	if err != nil {
		fmt.Println(err)
		return
	}

	var n int
	var d int
	var s int

	for rows.Next() {
		err = rows.Scan(&n, &d, &s)
		temp := Entry1{Number: n, Double: d, Square: s}
		DATA = append(DATA, temp)
	}

	http.HandleFunc("/", myHandler)
	err = http.ListenAndServe(":50416", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
