package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

var db *sql.DB

func login(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Attempting login from user %s.", request.Form.Get("login"))

}

func getTableVersion(tableName string) int64 {
	var ver int64
	_, err := db.Query(`SELECT table_comment FROM INFORMATION_SCHEMA.TABLES where table_name like "?"`, tableName) //.Scan(&ver)
	if err != nil {
		ver = 0
	}
	return ver
}

func initDB() {
	if getTableVersion("users") < 1 {
		query := `drop table if exists users;`
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}
		query = `
			CREATE TABLE users (
				id INT NOT NULL,
				username VARCHAR(20) NOT NULL,
				password CHAR(80) NOT NULL) comment="1"; `
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/login", login)
	var err error
	db, err = sql.Open("mysql", os.Getenv("DBUSER")+":"+os.Getenv("DBPASS")+"@unix("+os.Getenv("DBADDR")+")/"+os.Getenv("DBDB"))
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	initDB()

	http.ListenAndServe(":53714", r)
}
