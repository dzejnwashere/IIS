package db

import (
	"fmt"
	"log"
)

func GetSPZs() []string {
	var spzs []string
	var spz string

	query := `SELECT spz FROM vozy;`

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&spz)

		if err != nil {
			log.Fatal(err)
		}
		spzs = append(spzs, spz)
	}
	return spzs
}

func SPZexists(input string) bool {
	query := `SELECT COUNT(*) FROM vozy WHERE spz=?;`
	var cnt int

	err := db.QueryRow(query, input).Scan(&cnt)

	if err != nil {
		fmt.Printf("ERROR") //TODO
	}

	if cnt == 0 {
		return false
	} else {
		return true
	}
}
