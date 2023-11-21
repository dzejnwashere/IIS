package db

import (
	"IIS/typedef"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strconv"
	"time"
)

var db *sql.DB

func getTableVersion(tableName string) int64 {
	var ver int64
	res := db.QueryRow(`SELECT table_comment FROM INFORMATION_SCHEMA.TABLES where table_name like ?`, tableName)

	err := res.Scan(&ver)
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

func GetUsername(id int64) (string, error) {
	var username string
	err := db.QueryRow(`SELECT username from users where id = ?`, id).Scan(&username)
	if err != nil {
		return "", fmt.Errorf("no such user found") //TODO funkce DoesUserExist
	}
	return username, nil
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

func RemoveUser(userID int) {
	query := `DELETE FROM users WHERE id = ?;`
	_, err := db.Exec(query, userID)
	if err != nil {
		log.Fatal(err.Error())
	}
}

type User struct {
	ID          int
	Username    string
	Permissions int
}

func GetAllUsers() []User {
	query := `SELECT id, username, permissions FROM users;`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	var users []User
	var id int
	var username string
	var permissions int

	for rows.Next() {
		err := rows.Scan(&id, &username, &permissions)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, username, permissions)
		users = append(users, User{ID: id, Username: username, Permissions: permissions})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return users
}

func FeedDemoData() error {
	file, err := os.ReadFile("res/db/demo.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(file))
	return err
}

func InitDB() {
	var err error
	db, err = sql.Open("mysql", os.Getenv("DBSTRING")+"?charset=utf8mb4&multiStatements=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	query := `SET NAMES 'utf8mb4';`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err.Error())
	}

	if getTableVersion("users") < 6 {
		fmt.Printf("%d", getTableVersion("users"))
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
	   				PRIMARY KEY (id)) comment="6" character set utf8mb4; `
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}
		passHash, _ := bcrypt.GenerateFromPassword([]byte("admin"), 10)
		CreateOrUpdateUser(-1, "admin", string(passHash), typedef.AdminPerm)
	}

	optionallyCreateTable := func(name string, ver int64, stmt string) {
		if getTableVersion(name) < ver {
			_, err := db.Exec(`drop table if exists ` + name)
			if err != nil {
				log.Fatal(err.Error())
			}

			_, err = db.Exec(stmt)
			if err != nil {
				log.Fatal(err.Error())
			}
			_, err = db.Exec("ALTER TABLE " + name + " COMMENT=\"" + strconv.Itoa(int(ver)) + "\"")
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	}

	optionallyCreateTable("spravci", 6, `
	   			CREATE TABLE spravci (
	   				id INT AUTO_INCREMENT PRIMARY KEY,
	       			jmeno VARCHAR(20),
	       			prijmeni VARCHAR(30),
	       			user INT,
	       			FOREIGN KEY (user) REFERENCES users(id) ON DELETE CASCADE) character set utf8mb4;`)

	optionallyCreateTable("technici", 6, `
	   			CREATE TABLE technici (
	   				id INT AUTO_INCREMENT PRIMARY KEY,
	       			jmeno VARCHAR(20),
	       			prijmeni VARCHAR(30),
	       			user INT,
	       			FOREIGN KEY (user) REFERENCES users(id) ON DELETE CASCADE) comment="6" character set utf8mb4;`)

	optionallyCreateTable("dispeceri", 6, `
	   			CREATE TABLE dispeceri (
	   				id INT AUTO_INCREMENT PRIMARY KEY,
	       			jmeno VARCHAR(20),
	       			prijmeni VARCHAR(30),
	       			user INT,
	       			FOREIGN KEY (user) REFERENCES users(id) ON DELETE CASCADE) comment="6" character set utf8mb4;`)

	optionallyCreateTable("ridici", 6, `
	   			CREATE TABLE ridici (
	   				id INT AUTO_INCREMENT PRIMARY KEY,
	       			jmeno VARCHAR(20),
	       			prijmeni VARCHAR(30),
	       			user INT,
	       			FOREIGN KEY (user) REFERENCES users(id) ON DELETE CASCADE ) comment="6" character set utf8mb4;`)

	optionallyCreateTable("zastavky", 6, `
	   			CREATE TABLE zastavky (
	   				id INT AUTO_INCREMENT PRIMARY KEY,
	       			nazov_zastavky VARCHAR(255)) comment="6" character set utf8mb4;`)

	optionallyCreateTable("linky", 6, `
	   			CREATE TABLE linky (
	   				id INT AUTO_INCREMENT PRIMARY KEY,
	       			nazev INT) comment="6" character set utf8mb4;`)

	optionallyCreateTable("linka_zastavka", 6, `
	   			CREATE TABLE linka_zastavka (
	   				poradi INT,
	       			cas VARCHAR(50), -- tbd the len of varchar
	   			    zastavka INT NOT NULL,
	   			    linka INT NOT NULL,
	   				PRIMARY KEY (zastavka, linka),
	   				FOREIGN KEY (zastavka) REFERENCES zastavky(id),
	   				FOREIGN KEY (linka) REFERENCES linky(id)
	   			) comment="6" character set utf8mb4;`)

	optionallyCreateTable("vozy", 6, `
	   			CREATE TABLE vozy (
	   				spz VARCHAR(7) PRIMARY KEY,
	   			    druh VARCHAR(20),
	   			    znacka VARCHAR(50) NOT NULL,
	   			    kapacita INT) comment="6" character set utf8mb4;`)

	optionallyCreateTable("stav_zavady", 6, `
	   			CREATE TABLE stav_zavady (
	   				id INT AUTO_INCREMENT PRIMARY KEY,
	   			    stav VARCHAR(255)
	   			    ) comment="6" character set utf8mb4;`)

	optionallyCreateTable("zavady", 6, `
	   			CREATE TABLE zavady (
	   				id INT AUTO_INCREMENT PRIMARY KEY,
	       			spz VARCHAR(7),
	   			    autor INT,
	   			    popis VARCHAR(255),
	   			    stav INT,
	   			    technik INT,
	   			    FOREIGN KEY (spz) REFERENCES vozy(spz),
	   			    FOREIGN KEY (autor) REFERENCES users(id) ON DELETE SET NULL,
	   			    FOREIGN KEY (technik) REFERENCES users(id) ON DELETE SET NULL) comment="6" character set utf8mb4;`)

	optionallyCreateTable("tech_zaznamy", 6, `
	   			CREATE TABLE tech_zaznamy (
	   				spz_vozidla varchar(7),
	       			datum DATE,
	   			    zavada INT,
	   			    PRIMARY KEY (spz_vozidla, datum),
	   			    FOREIGN KEY (spz_vozidla) REFERENCES vozy(spz),
	   			    FOREIGN KEY (zavada) REFERENCES zavady(id) ON DELETE SET NULL) comment="6" character set utf8mb4;`)

	optionallyCreateTable("dny_jizdy", 6, `
	   			CREATE TABLE dny_jizdy (
	   			    id INT AUTO_INCREMENT PRIMARY KEY,
	   			    den_jizdy VARCHAR(50) NOT NULL
	   			    ) comment="6" character set utf8mb4;`)

	optionallyCreateTable("spoje", 6, `
	   			CREATE TABLE spoje (
	   			    linka INT NOT NULL,
	   				cas_odjezdu varchar(10) NOT NULL,
	   				smer_jizdy INT NOT NULL, -- cislo zastavky
	       			dny_jizdy INT NOT NULL,
	       			vuz VARCHAR(7) NOT NULL,
	       			PRIMARY KEY (linka, vuz, smer_jizdy, cas_odjezdu, dny_jizdy),
	       			FOREIGN KEY (vuz) REFERENCES vozy(spz),
	       			FOREIGN KEY (linka) REFERENCES linky(id),
	       			FOREIGN KEY (dny_jizdy) REFERENCES dny_jizdy(id),
	       			FOREIGN KEY (smer_jizdy) REFERENCES zastavky(id)) comment="6" character set utf8mb4;`)

}
