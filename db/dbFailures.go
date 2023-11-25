package db

import "log"

type Failure struct {
	FailureID         int
	SPZ               string
	Description       string
	TechnicianID      int
	TechnicianName    string
	TechnicianSurname string
	State             string
	AuthorId          int
	AuthorName        string
	AuthorSurname     string
}

func GetFailures() []Failure {
	query := `SELECT z.id, z.SPZ, z.popis, t.user, t.jmeno, t.prijmeni, sz.stav, s.id, s.jmeno, s.prijmeni FROM zavady z
			  JOIN spravci s ON z.autor=s.user
			  JOIN technici t ON z.technik=t.user
			  JOIN stav_zavady sz ON z.stav=sz.id;`

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

	query := `SELECT z.id, z.SPZ, z.popis, t.user, t.jmeno, t.prijmeni, sz.stav, s.id, s.jmeno, s.prijmeni FROM zavady z
				JOIN spravci s ON z.autor=s.user
			  	JOIN technici t ON z.technik=t.user
			  	JOIN stav_zavady sz ON z.stav=sz.id
			  WHERE z.id = ?;`

	err := db.QueryRow(query, id).Scan(&fail.FailureID, &fail.SPZ, &fail.Description, &fail.TechnicianID, &fail.TechnicianName, &fail.TechnicianSurname, &fail.State, &fail.AuthorId, &fail.AuthorName, &fail.AuthorSurname)

	if err != nil {
		log.Fatal("GetFailureById: " + err.Error())
	}

	return fail
}

func GetFailuresForSpecificSPZWithSpecificState(SPZ string, state int) []Failure {
	query := `SELECT z.id, z.SPZ, z.popis, t.user, t.jmeno, t.prijmeni, sz.stav, s.id, s.jmeno, s.prijmeni FROM zavady z
			  JOIN spravci s ON z.autor=s.user
			  JOIN technici t ON z.technik=t.user
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
