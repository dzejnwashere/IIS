package db

import (
	"IIS/typedef"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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
	err := db.QueryRow(`SELECT permissions FROM users WHERE id = ?`, userID).Scan(&perm)
	if err != nil {
		fmt.Printf("db.GetPermissions error: %s", err.Error())
		return 0
	}
	return perm
}

func GetUserIdPasswordHash(username string) (int64, string, error) {
	var hash string
	var id int64
	err := db.QueryRow(`SELECT id, password from users where username = ?`, username).Scan(&id, &hash)
	if err != nil {
		return 0, "", fmt.Errorf("no such user found") //TODO funkce DoesUserExist
	}
	return id, hash, nil
}

// With id < 0, create a new user. Returns id of user
func CreateOrUpdateUser(id int, username string, passHash string, permission ...typedef.Permission) (int64, error) {
	permInt := 0
	for _, a := range permission {
		permInt = permInt | (1 << a)
	}
	if id < 0 {
		res, err := db.Exec(`INSERT INTO users (username, password, permissions) VALUES (?, ?, ?)`, username, passHash, permInt)
		if err != nil {
			return 0, err
		}
		return res.LastInsertId()
	} else {
		_, err := db.Exec(`UPDATE users SET username = ?, password = ?, permissions = ? WHERE id = ?`, username, passHash, permInt, id)
		return int64(id), err
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

	if getTableVersion("users") < 6 {
		query := `drop table if exists users;`
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}
		query = `
			CREATE TABLE users (
				id INT NOT NULL AUTO_INCREMENT,
				username VARCHAR(20) NOT NULL,
				password CHAR(80) NOT NULL,
				permissions int not null,
				UNIQUE (username),
				PRIMARY KEY (id)) comment="6"; `
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}
		passHash, _ := bcrypt.GenerateFromPassword([]byte("admin"), 10)
		CreateOrUpdateUser(-1, "admin", string(passHash), typedef.AdminPerm)
	}
}
