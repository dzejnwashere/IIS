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

func GetStops() []string {
	var stops []string
	var stop string

	query := `SELECT nazov_zastavky FROM zastavky;`

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&stop)

		if err != nil {
			log.Fatal(err)
		}
		stops = append(stops, stop)
	}
	return stops
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

	optionallyCreateTable("linka_zastavka", 7, `
	   			CREATE TABLE linka_zastavka (
	       			cas TIME,
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
	   			    popis VARCHAR(255),
	   			    autor INT NOT NULL ,
	   			    PRIMARY KEY (spz_vozidla, datum),
	   			    FOREIGN KEY (autor) REFERENCES technici(user),
	   			    FOREIGN KEY (spz_vozidla) REFERENCES vozy(spz),
	   			    FOREIGN KEY (zavada) REFERENCES zavady(id)) comment="6" character set utf8mb4;`)

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
