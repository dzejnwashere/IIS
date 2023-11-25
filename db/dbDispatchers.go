package db

import (
	"fmt"
	"log"
)

func GetDispatchers() []UserData {
	var dispatcher UserData
	var dispatchers []UserData

	query := `SELECT id, name, surname FROM users WHERE (permissions & 8) <> 0`

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&dispatcher.ID, &dispatcher.Name, &dispatcher.Surname)

		if err != nil {
			log.Fatal(err)
		}
		dispatchers = append(dispatchers, dispatcher)
	}
	return dispatchers
}

func GetDispatcher(userID int64) UserData {
	var dispatcher UserData

	query := `SELECT id, name, surname FROM users WHERE (permissions & 8) <> 0 AND id = ?`

	err := db.QueryRow(query, userID).Scan(&dispatcher.ID, &dispatcher.Name, &dispatcher.Surname)
	if err != nil {
		fmt.Printf("db.GetPermissions error: %s", err.Error())
	}

	return dispatcher
}
