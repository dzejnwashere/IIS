package db

import (
	"IIS/auth"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

var db *sql.DB

func getTableVersion(tableName string) int64 {
	var ver int64
	_, err := db.Query(`SELECT table_comment FROM INFORMATION_SCHEMA.TABLES where table_name like "?"`, tableName) //.Scan(&ver)
	if err != nil {
		ver = 0
	}
	return ver
}

func GetPermissions(userID int64) int64 {
	var perm int64
	err := db.QueryRow(`SELECT permission FROM users WHERE id = ?`, userID).Scan(&perm)
	if err != nil {
		return 0
	}
	return perm
}

func GetUserIdPasswordHash(username string) (int64, string, error) {
	var hash string
	var id int64
	err := db.QueryRow(`SELECT id, password from users where username = ?`, username).Scan(&id, &hash)
	if err != nil {
		return 0, "", fmt.Errorf("No such user found") //TODO funkce DoesUserExist
	}
	return id, hash, nil
}

func CreateOrUpdateUser(id int, username string, passHash string, permission ...auth.Permission) {
	permInt := 0
	for _, a := range permission {
		permInt = permInt | (1 << a)
	}
	if id < 0 {
		//TODO kompletně err := db.
	}
}

func InitDB() {

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

	if getTableVersion("users") < 4 {
		query := `drop table if exists users;`
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}
		query = `
			CREATE TABLE users (
				id INT NOT NULL,
				username VARCHAR(20) NOT NULL,
				password CHAR(80) NOT NULL,
				permissions int not null) comment="4"; `
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
