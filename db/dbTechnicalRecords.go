package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type TechnicalRecord struct {
	ID            int
	SPZ           string
	Date          string
	Failures      []Failure
	Details       string
	AuthorID      int
	AuthorName    string
	AuthorSurname string
}

type CreateTechnicalRecord struct {
	SPZ                string
	Date               string
	FailureID          []int
	FailureDescription string
	Details            string
	AuthorID           int
	AuthorName         string
	AuthorSurname      string
}

func GetTechnicalRecords() []TechnicalRecord {
	query := `SELECT tz.id, tz.spz_vozidla, tz.datum, tz.popis, t.id, t.name, t.surname FROM tech_zaznamy tz
			  JOIN users t ON tz.autor=t.id;`

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

	var technicalRecords []TechnicalRecord
	var technicalRecord TechnicalRecord
	for rows.Next() {
		var dateStr string

		err := rows.Scan(&technicalRecord.ID, &technicalRecord.SPZ, &dateStr, &technicalRecord.Details, &technicalRecord.AuthorID, &technicalRecord.AuthorName, &technicalRecord.AuthorSurname)
		if err != nil {
			log.Fatal(err)
		}

		split := strings.Split(dateStr, "T")
		technicalRecord.Date = split[0]
		technicalRecord.Failures = GetFailuresForTechRecord(technicalRecord.ID)
		technicalRecords = append(technicalRecords, technicalRecord)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return technicalRecords
}

func CreateNewTechnicalRecord(techRecord CreateTechnicalRecord) CreateTechnicalRecord {
	fmt.Println(techRecord)
	query := `INSERT INTO tech_zaznamy (spz_vozidla, datum, popis, autor) VALUES (?, ?, ?, ?); SELECT LAST_INSERT_ID();
`
	var newID int
	err := db.QueryRow(query, techRecord.SPZ, techRecord.Date, techRecord.Details, techRecord.AuthorID).Scan(&newID)

	fmt.Println(newID)
	// TODO
	//AssignFailuresToTechRecord()

	if err != nil {
		log.Fatal("CreateNewTechnicalRecord: " + err.Error())
	}
	return techRecord
}
