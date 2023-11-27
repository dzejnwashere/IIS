package db

import (
	"IIS/typedef"
	"fmt"
	"log"
)

type User struct {
	ID          int
	Username    string
	Name        string
	Surname     string
	Permissions int
}

type UserData struct {
	ID      int
	Name    string
	Surname string
}

func GetAllUsers() []User {
	query := `SELECT id, username, name, surname, permissions FROM users;`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	var user User
	var users []User

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Name, &user.Surname, &user.Permissions)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return users
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
func CreateOrUpdateUser(id int, username string, passHash string, name string, surname string, permission ...typedef.Permission) (int64, error) {
	permInt := 0
	for _, a := range permission {
		permInt = permInt | (1 << a)
	}
	if id < 0 {
		res, err := db.Exec(`INSERT INTO users (username, password, name, surname, permissions) VALUES (?, ?, ?, ?, ?)`, username, passHash, name, surname, permInt)
		if err != nil {
			return 0, err
		}
		return res.LastInsertId()
	} else {
		_, err := db.Exec(`UPDATE users SET username = ?, password = ?, name = ?, surname = ?, permissions = ? WHERE id = ?`, username, passHash, name, surname, permInt, id)
		return int64(id), err
	}
}

func UpdatePermissions(id int, permissions int) error {
	_, err := db.Exec(`UPDATE users SET permissions = ? WHERE id = ?`, permissions, id)
	fmt.Println(err)
	return err
}

func RemoveUser(userID int) {
	query := `DELETE FROM users WHERE id = ?;`
	_, err := db.Exec(query, userID)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetUser(userID int64) UserData {
	fmt.Println(userID)
	query := `SELECT id, name, surname FROM users WHERE id = ?;`

	var userData UserData
	err := db.QueryRow(query, userID).Scan(&userData.ID, &userData.Name, &userData.Surname)
	if err != nil {
		log.Fatal("db.GetUser() " + err.Error())
	}

	return userData
}

func UsernameExists(username string) bool {
	query := "SELECT COUNT(*) FROM users WHERE username = ?"
	var count int
	err := db.QueryRow(query, username).Scan(&count)
	if err != nil {
		return false
	}

	return count > 0
}
