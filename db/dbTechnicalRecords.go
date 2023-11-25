package db

import (
	"database/sql"
	"fmt"
	"log"
)

type TechnicalRecord struct {
	SPZ           string
	Date          string
	Failure       Failure
	Details       string
	AuthorID      int
	AuthorName    string
	AuthorSurname string
}

type CreateTechnicalRecord struct {
	SPZ       string
	Date      string
	FailureID int
	Details   string
	AuthorID  int
}

func GetTechnicalRecords() []TechnicalRecord {
	query := `SELECT tz.spz_vozidla, tz.datum, tz.zavada, tz.popis, t.user, t.jmeno, t.prijmeni FROM tech_zaznamy tz
			  JOIN technici t ON tz.autor=t.user;`

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
	var zavadaIDNull *int
	var zavadaID sql.NullInt64

	for rows.Next() {
		err := rows.Scan(&technicalRecord.SPZ, &technicalRecord.Date, &zavadaIDNull, &technicalRecord.Details, &technicalRecord.AuthorID, &technicalRecord.AuthorName, &technicalRecord.AuthorSurname)
		if err != nil {
			log.Fatal(err)
		}

		if zavadaIDNull != nil {
			zavadaID = sql.NullInt64{int64(*zavadaIDNull), true}
		} else {
			zavadaID = sql.NullInt64{Valid: false}
		}

		if zavadaID.Valid {
			technicalRecord.Failure = GetFailureById(int(zavadaID.Int64))
		}

		technicalRecords = append(technicalRecords, technicalRecord)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return technicalRecords
}

func CreateNewTechnicalRecord(techRecord CreateTechnicalRecord) {
	fmt.Println(techRecord)
	query := `INSERT INTO tech_zaznamy (spz_vozidla, datum, zavada, popis, autor) VALUES
                                                                        (?, ?, ?, ?, ?);`

	_, err := db.Exec(query, techRecord.SPZ, techRecord.Date, techRecord.FailureID, techRecord.Details, techRecord.AuthorID)

	if err != nil {
		log.Fatal("CreateNewTechnicalRecord: " + err.Error())
	}
}
