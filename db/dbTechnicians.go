package db

import (
	"fmt"
	"log"
)

func GetTechnicians() []UserData {
	var technician UserData
	var technicians []UserData

	query := `SELECT id, name, surname FROM users WHERE (permissions & 4) <> 0`

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&technician.ID, &technician.Name, &technician.Surname)

		if err != nil {
			log.Fatal(err)
		}
		technicians = append(technicians, technician)
	}
	return technicians
}

func GetTechnician(userID int64) UserData {
	var technician UserData

	query := `SELECT id, name, surname FROM users WHERE (permissions & 4) <> 0 AND id = ?`

	err := db.QueryRow(query, userID).Scan(&technician.ID, &technician.Name, &technician.Surname)
	if err != nil {
		fmt.Printf("db.GetPermissions error: %s", err.Error())
	}

	return technician
}
