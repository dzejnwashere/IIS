package db

import (
	"fmt"
	"log"
)

func GetManagers() []UserData {
	var manager UserData
	var managers []UserData

	query := `SELECT id, name, surname FROM users WHERE (permissions & 16) <> 0`

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&manager.ID, &manager.Name, &manager.Surname)

		if err != nil {
			log.Fatal(err)
		}
		managers = append(managers, manager)
	}
	return managers
}

func GetManager(userID int64) UserData {
	var manager UserData

	query := `SELECT id, name, surname FROM users WHERE (permissions & 16) <> 0 AND id = ?`

	err := db.QueryRow(query, userID).Scan(&manager.ID, &manager.Name, &manager.Surname)
	if err != nil {
		fmt.Printf("db.GetPermissions error: %s", err.Error())
	}

	return manager
}
