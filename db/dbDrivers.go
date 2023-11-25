package db

import (
	"fmt"
	"log"
)

func GetDrivers() []UserData {
	var driver UserData
	var drivers []UserData

	query := `SELECT id, name, surname FROM users WHERE (permissions & 2) <> 0`

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&driver.ID, &driver.Name, &driver.Surname)

		if err != nil {
			log.Fatal(err)
		}
		drivers = append(drivers, driver)
	}
	return drivers
}

func GetDriver(userID int64) UserData {
	var driver UserData

	query := `SELECT id, name, surname FROM users WHERE (permissions & 2) <> 0 AND id = ?`

	err := db.QueryRow(query, userID).Scan(&driver.ID, &driver.Name, &driver.Surname)
	if err != nil {
		fmt.Printf("db.GetPermissions error: %s", err.Error())
	}

	return driver
}
