package db

import "log"

type Technician struct {
	ID      int
	Name    string
	Surname string
}

func GetTechnicians() []Technician {
	var technician Technician
	var technicians []Technician

	query := `SELECT user, jmeno, prijmeni FROM technici;`

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
