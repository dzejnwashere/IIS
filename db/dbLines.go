package db

import (
	"log"
)

type Line_t struct {
	Id   int
	Name string
}

func GetAllLines() []Line_t {
	var line string
	var id int
	var lines []Line_t
	rows, err := db.Query("SELECT id, nazev from linky")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &line)

		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, Line_t{
			Id:   id,
			Name: line,
		})
	}
	return lines
}

type Stop_line_t struct {
	Stop_id   int
	Stop_name string
	Time      string
	Line_id   int
}

func GetLineStops(lineId int) []Stop_line_t {
	var time, name string
	var id int
	var stops []Stop_line_t
	rows, err := db.Query("SELECT lz.zastavka, z.nazov_zastavky, lz.cas from linka_zastavka lz join zastavky z on z.id=lz.zastavka where lz.linka = ? order by lz.cas asc ", lineId)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name, &time)

		if err != nil {
			log.Fatal(err)
		}
		stops = append(stops, Stop_line_t{
			Stop_id:   id,
			Stop_name: name,
			Time:      time,
			Line_id:   lineId,
		})
	}
	return stops
}

func AddLineStops(data Stop_line_t) error {
	_, err := db.Exec("INSERT INTO linka_zastavka (cas, zastavka, linka) values (?, ?, ?)", data.Time, data.Stop_id, data.Line_id)
	return err
}

type Stop_t struct {
	Id   int
	Name string
}

func GetStops2() []Stop_t {
	var name string
	var id int
	var stops []Stop_t
	rows, err := db.Query("Select id, nazov_zastavky from zastavky")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		stops = append(stops, Stop_t{
			Id:   id,
			Name: name,
		})
	}
	return stops
}
