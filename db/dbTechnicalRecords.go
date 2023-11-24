package db

import (
	"database/sql"
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

func GetTechnicalRecords() []TechnicalRecord {
	query := `SELECT tz.spz_vozidla, tz.datum, tz.zavada, tz.popis, t.user, t.jmeno, t.prijmeni FROM tech_zaznamy tz
			  JOIN technici t ON tz.autor=t.user;`

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

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
