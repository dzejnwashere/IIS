package db

import (
	"fmt"
	"log"
)

func GetAdmins() []UserData {
	var admin UserData
	var admins []UserData

	query := `SELECT id, name, surname FROM users WHERE (permissions & 1) <> 0`

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&admin.ID, &admin.Name, &admin.Surname)

		if err != nil {
			log.Fatal(err)
		}
		admins = append(admins, admin)
	}
	return admins
}

func GetAdmin(userID int64) UserData {
	var admin UserData

	query := `SELECT id, name, surname FROM users WHERE (permissions & 1) <> 0 AND id = ?`

	err := db.QueryRow(query, userID).Scan(&admin.ID, &admin.Name, &admin.Surname)
	if err != nil {
		fmt.Printf("db.GetPermissions error: %s", err.Error())
	}

	return admin
}
