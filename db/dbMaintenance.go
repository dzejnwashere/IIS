package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type Maintenance struct {
	SPZ  string
	Date string
}

func GetAllMaintenance() []Maintenance {
	query := `SELECT spz, datum FROM udrzba ORDER BY datum;`

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal("GetTechnicalRecords: " + err.Error())
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal("ROWS: " + err.Error())
		}
	}(rows)

	var maintenanceArray []Maintenance
	var maintenance Maintenance
	for rows.Next() {
		var dateStr string

		err := rows.Scan(&maintenance.SPZ, &dateStr)
		if err != nil {
			log.Fatal(err)
		}

		split := strings.Split(dateStr, "T")
		maintenance.Date = split[0]
		maintenanceArray = append(maintenanceArray, maintenance)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return maintenanceArray
}

func CreateNewMaintenance(newMaintenance Maintenance) (bool, Maintenance) {
	if MaintenanceExists(newMaintenance) {
		return false, Maintenance{SPZ: "", Date: ""}
	} else {
		query := `INSERT INTO udrzba (spz, datum) VALUES (?, ?);`

		_, err := db.Exec(query, newMaintenance.SPZ, newMaintenance.Date)
		if err != nil {
			log.Fatal("CreateNewTechnicalRecord:LastInsertId: " + err.Error())
		}

		return false, newMaintenance
	}
}

func ReplaceMaintenance(oldMaintenance Maintenance, newMaintenance Maintenance) bool {
	if MaintenanceExists(newMaintenance) {
		return false
	} else {
		query := `DELETE FROM udrzba WHERE spz = ? AND datum = ?;`

		_, err := db.Exec(query, newMaintenance.SPZ, newMaintenance.Date)
		if err != nil {
			log.Fatal("CreateNewTechnicalRecord:LastInsertId: " + err.Error())
		}

		_, _ = CreateNewMaintenance(newMaintenance)
		return true
	}
}

func MaintenanceExists(maintenance Maintenance) bool {
	query := `SELECT COUNT(*) FROM udrzba WHERE spz=? AND datum = ?;`
	var cnt int

	err := db.QueryRow(query, maintenance.SPZ, maintenance.Date).Scan(&cnt)

	if err != nil {
		fmt.Printf("ERROR") //TODO
	}

	if cnt == 0 {
		return false
	} else {
		return true
	}
}
