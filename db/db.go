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
	res := db.QueryRow(`SELECT table_comment FROM INFORMATION_SCHEMA.TABLES where table_name like ?`, tableName)

	err := res.Scan(&ver)
	if err != nil {
		log.Print(err.Error())
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
	query := `SELECT permissions FROM users WHERE id = ?;`
	row, _ := db.Query(query, userID)

	defer func(perms *sql.Rows) {
		err := perms.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(row)
	var permissions int

	for row.Next() {
		err := row.Scan(&permissions)
		if err != nil {
			log.Fatal(err)
		}
	}

	err := row.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(permissions)

	if permissions == 2 {
		query = `DELETE FROM ridici WHERE user = ?;`
		_, err = db.Exec(query, userID)
		if err != nil {
			log.Fatal(err.Error())
		}
	} else if permissions == 4 {
		fmt.Println("TUTUTUTU")
		query = `DELETE FROM dispeceri WHERE user = ?;`
		_, err = db.Exec(query, userID)
		if err != nil {
			log.Fatal(err.Error())
		}
	} else if permissions == 8 {
		query = `DELETE FROM technici WHERE user = ?;`
		_, err = db.Exec(query, userID)
		if err != nil {
			log.Fatal(err.Error())
		}
	} else if permissions == 16 {
		query = `DELETE FROM spravci WHERE user = ?;`
		_, err = db.Exec(query, userID)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	query = `DELETE FROM users WHERE id = ?;`
	_, err = db.Exec(query, userID)
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

func InitDB() {
	var err error
	db, err = sql.Open("mysql", os.Getenv("DBSTRING"))
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
	if getTableVersion("spravci") < 6 {
		query := `drop table if exists spravci;`
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}

		query = `
	   			CREATE TABLE spravci (
	   				id INT AUTO_INCREMENT PRIMARY KEY,
	       			jmeno VARCHAR(20),
	       			prijmeni VARCHAR(30),
	       			user INT,
	       			FOREIGN KEY (user) REFERENCES users(id) ON DELETE SET NULL) comment="6" character set utf8mb4;`
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	if getTableVersion("technici") < 6 {
		query := `drop table if exists technici;`
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}

		query = `
	   			CREATE TABLE technici (
	   				id INT AUTO_INCREMENT PRIMARY KEY,
	       			jmeno VARCHAR(20),
	       			prijmeni VARCHAR(30),
	       			user INT,
	       			FOREIGN KEY (user) REFERENCES users(id) ON DELETE SET NULL) comment="6" character set utf8mb4;`
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	if getTableVersion("dispeceri") < 6 {
		query := `drop table if exists dispeceri;`
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}

		query = `
	   			CREATE TABLE dispeceri (
	   				id INT AUTO_INCREMENT PRIMARY KEY,
	       			jmeno VARCHAR(20),
	       			prijmeni VARCHAR(30),
	       			user INT,
	       			FOREIGN KEY (user) REFERENCES users(id) ON DELETE SET NULL) comment="6" character set utf8mb4;`
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	if getTableVersion("ridici") < 6 {
		query := `drop table if exists ridici;`
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}

		query = `
	   			CREATE TABLE ridici (
	   				id INT AUTO_INCREMENT PRIMARY KEY,
	       			jmeno VARCHAR(20),
	       			prijmeni VARCHAR(30),
	       			user INT,
	       			FOREIGN KEY (user) REFERENCES users(id)) comment="6" character set utf8mb4;`
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	//create_users()

	if getTableVersion("zastavky") < 6 {
		query := `drop table if exists zastavky;`
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}

		query = `
	   			CREATE TABLE zastavky (
	   				id INT AUTO_INCREMENT PRIMARY KEY,
	       			nazov_zastavky VARCHAR(255)) comment="6" character set utf8mb4;`
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}

		feed_zastavky()
	}

	if getTableVersion("linky") < 6 {
		query := `drop table if exists linky;`
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}

		query = `
	   			CREATE TABLE linky (
	   				id INT AUTO_INCREMENT PRIMARY KEY,
	       			nazev INT) comment="6" character set utf8mb4;`
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}
		feed_linky()
	}

	if getTableVersion("linka_zastavka") < 6 {
		query := `drop table if exists linka_zastavka;`
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}

		query = `
	   			CREATE TABLE linka_zastavka (
	   				poradi INT,
	       			cas VARCHAR(50), -- tbd the len of varchar
	   			    zastavka INT NOT NULL,
	   			    linka INT NOT NULL,
	   				PRIMARY KEY (zastavka, linka),
	   				FOREIGN KEY (zastavka) REFERENCES zastavky(id),
	   				FOREIGN KEY (linka) REFERENCES linky(id)
	   			) comment="6" character set utf8mb4;`
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}
		feed_linka_zastavka()
	}

	if getTableVersion("vozy") < 6 {
		query := `drop table if exists vozy;`
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}

		query = `
	   			CREATE TABLE vozy (
	   				spz VARCHAR(7) PRIMARY KEY,
	   			    druh VARCHAR(20),
	   			    znacka VARCHAR(50) NOT NULL,
	   			    kapacita INT) comment="6" character set utf8mb4;`
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}
		feed_vozy()
	}

	if getTableVersion("stav_zavady") < 6 {
		query := `drop table if exists stav_zavady;`
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}

		query = `
	   			CREATE TABLE stav_zavady (
	   				id INT AUTO_INCREMENT PRIMARY KEY,
	   			    stav VARCHAR(255)
	   			    ) comment="6" character set utf8mb4;`
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}
		feed_stav_zavady()
	}

	if getTableVersion("zavady") < 6 {
		query := `drop table if exists zavady;`
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}

		query = `
	   			CREATE TABLE zavady (
	   				id INT AUTO_INCREMENT PRIMARY KEY,
	       			spz VARCHAR(7),
	   			    autor INT,
	   			    popis VARCHAR(255),
	   			    stav INT,
	   			    technik INT,
	   			    FOREIGN KEY (spz) REFERENCES vozy(spz),
	   			    FOREIGN KEY (autor) REFERENCES spravci(id) ON DELETE SET NULL,
	   			    FOREIGN KEY (technik) REFERENCES technici(id) ON DELETE SET NULL) comment="6" character set utf8mb4;`
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}
		feed_zavady()
	}

	if getTableVersion("tech_zaznamy") < 6 {
		query := `drop table if exists tech_zaznamy;`
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}

		query = `
	   			CREATE TABLE tech_zaznamy (
	   				spz_vozidla varchar(7),
	       			datum DATE,
	   			    zavada INT,
	   			    PRIMARY KEY (spz_vozidla, datum),
	   			    FOREIGN KEY (spz_vozidla) REFERENCES vozy(spz),
	   			    FOREIGN KEY (zavada) REFERENCES zavady(id) ON DELETE SET NULL) comment="6" character set utf8mb4;`
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}
		feed_tech_zaznamy()
	}

	if getTableVersion("dny_jizdy") < 6 {
		query := `drop table if exists dny_jizdy;`
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}

		query = `
	   			CREATE TABLE dny_jizdy (
	   			    id INT AUTO_INCREMENT PRIMARY KEY,
	   			    den_jizdy VARCHAR(50) NOT NULL
	   			    ) comment="6" character set utf8mb4;`
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}
		feed_dny_jizdy()
	}

	if getTableVersion("spoje") < 6 {
		query := `drop table if exists spoje;`
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}

		query = `
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
	       			FOREIGN KEY (smer_jizdy) REFERENCES zastavky(id)) comment="6" character set utf8mb4;`
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err.Error())
		}
		feed_spoje()
	}

}
