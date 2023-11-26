package db

import (
	"database/sql"
	"fmt"
	"log"
)

type Failure struct {
	FailureID         int
	SPZ               string
	Description       string
	TechnicianID      *int
	TechnicianName    *string
	TechnicianSurname *string
	State             string
	AuthorId          int
	AuthorName        string
	AuthorSurname     string
}

type CreateFailure struct {
	SPZ          string
	AuthorID     int
	TechnicianID *int
	Description  string
	State        int
}

func GetFailures() []Failure {
	query := `SELECT z.id, z.SPZ, z.popis, t.id AS technician_id, t.name AS technician_name, t.surname AS technician_surname, sz.stav, s.id AS author_id, s.name AS author_name, s.surname AS author_surname
				FROM zavady z
				JOIN users s ON z.autor = s.id
				LEFT JOIN users t ON z.technik = t.id
				JOIN stav_zavady sz ON z.stav = sz.id
				ORDER BY z.id DESC;`

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	var failures []Failure
	var fail Failure

	for rows.Next() {
		err := rows.Scan(&fail.FailureID, &fail.SPZ, &fail.Description, &fail.TechnicianID, &fail.TechnicianName, &fail.TechnicianSurname, &fail.State, &fail.AuthorId, &fail.AuthorName, &fail.AuthorSurname)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(fail.FailureID, fail.SPZ, fail.Description, fail.TechnicianID, fail.TechnicianName, fail.TechnicianSurname, fail.State, fail.AuthorId, fail.AuthorName, fail.AuthorSurname)
		failures = append(failures, fail)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return failures
}

func GetFailureById(id int) Failure {
	var fail Failure

	query := `SELECT z.id, z.SPZ, z.popis, t.id, t.name, t.surname, sz.stav, s.id, s.name, s.surname FROM zavady z
				JOIN users s ON z.autor=s.id
			  	JOIN users t ON z.technik=t.id
			  	JOIN stav_zavady sz ON z.stav=sz.id
			  WHERE z.id = ?;`

	err := db.QueryRow(query, id).Scan(&fail.FailureID, &fail.SPZ, &fail.Description, &fail.TechnicianID, &fail.TechnicianName, &fail.TechnicianSurname, &fail.State, &fail.AuthorId, &fail.AuthorName, &fail.AuthorSurname)

	if err != nil {
		log.Fatal("GetFailureById: " + err.Error())
	}

	return fail
}

func GetFailuresForSpecificSPZWithSpecificState(SPZ string, state int) []Failure {
	query := `SELECT z.id, z.SPZ, z.popis, t.id, t.name, t.surname, sz.stav, s.id, s.name, s.surname FROM zavady z
			  JOIN users s ON z.autor=s.id
			  JOIN users t ON z.technik=t.id
			  JOIN stav_zavady sz ON z.stav=sz.id
            WHERE sz.id = ? AND z.SPZ = ?;`

	rows, err := db.Query(query, state, SPZ)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	var failures []Failure
	var fail Failure

	for rows.Next() {
		err := rows.Scan(&fail.FailureID, &fail.SPZ, &fail.Description, &fail.TechnicianID, &fail.TechnicianName, &fail.TechnicianSurname, &fail.State, &fail.AuthorId, &fail.AuthorName, &fail.AuthorSurname)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(fail.FailureID, fail.SPZ, fail.Description, fail.TechnicianID, fail.TechnicianName, fail.TechnicianSurname, fail.State, fail.AuthorId, fail.AuthorName, fail.AuthorSurname)
		failures = append(failures, fail)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return failures
}

func CreateNewFailure(failure CreateFailure) []Failure {
	var technicianID sql.NullInt32

	if failure.TechnicianID == nil {
		technicianID = sql.NullInt32{Valid: false}
	} else {
		technicianID = sql.NullInt32{Int32: int32(*failure.TechnicianID), Valid: true}
	}

	fmt.Println(failure)
	query := `INSERT INTO zavady (spz, autor, technik, popis, stav) VALUES
                                                                        (?, ?, ?, ?, ?);`

	_, err := db.Exec(query, failure.SPZ, failure.AuthorID, technicianID, failure.Description, failure.State)

	if err != nil {
		log.Fatal("CreateNewFailure: " + err.Error())
	}

	failures := GetFailures()

	return failures
}
