package db

import (
	"fmt"
	"log"
	"time"
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
func GetLine(id int) Line_t {
	var lines Line_t
	rows := db.QueryRow("SELECT id, nazev from linky where id = ?", id)
	err := rows.Scan(&lines.Id, &lines.Name)

	if err != nil {
		log.Fatal(err)
	}
	return lines
}

func DeleteLine(id int) error {
	_, err := db.Exec("DELETE FROM linky where id=?", id)
	return err
}

func CreateLine(linename string) error {
	_, err := db.Exec("INSERT INTO linky (nazev) values (?)", linename)
	return err
}

func DeleteStopLine(id int) error {
	_, err := db.Exec("DELETE FROM linka_zastavka where id=?", id)
	return err
}

type Stop_line_t struct {
	Stop_id      int
	Stop_name    string
	Time         string
	Line_id      int
	Stop_line_id int
}

func GetLineStops(lineId int) []Stop_line_t {
	var time, name string
	var id, slID int
	var stops []Stop_line_t
	rows, err := db.Query("SELECT lz.zastavka, z.nazov_zastavky, lz.cas, lz.id from linka_zastavka lz join zastavky z on z.id=lz.zastavka where lz.linka = ? order by lz.cas asc ", lineId)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name, &time, &slID)

		if err != nil {
			log.Fatal(err)
		}
		stops = append(stops, Stop_line_t{
			Stop_id:      id,
			Stop_name:    name,
			Time:         time,
			Line_id:      lineId,
			Stop_line_id: slID,
		})
	}
	return stops
}

func AddLineStops(data Stop_line_t) error {
	_, err := db.Exec("INSERT INTO linka_zastavka (cas, zastavka, linka) values (?, ?, ?)", data.Time, data.Stop_id, data.Line_id)
	return err
}

func UpdateLineStops(data Stop_line_t) error {
	_, err := db.Exec("UPDATE `linka_zastavka` SET `cas` = ?, `zastavka` = ?, `linka` = ? WHERE `linka_zastavka`.`id` = ?", data.Time, data.Stop_id, data.Line_id, data.Stop_line_id)
	return err
}

type Stop_t struct {
	Id   int
	Name string
}

func GetStop(id int) Stop_t {
	var spoj_temp Stop_t
	rows := db.QueryRow("Select id, nazov_zastavky from zastavky where id = ?", id)

	err := rows.Scan(&(spoj_temp.Id), &spoj_temp.Name)

	if err != nil {
		log.Fatal(err)
	}
	return spoj_temp
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

type LineFromStop struct {
	LineName int
	Time     string
	NextStop string
	Day      string
}

func GetLinesFromStop(stopName string) []LineFromStop {
	fmt.Println(stopName)
	var startTime string
	var stopTime int

	query := `SELECT l.nazev, lz.cas, s.cas_odjezdu, sz.nazov_zastavky, dj.den_jizdy  from linky l
			  join linka_zastavka lz ON lz.linka = l.id
			  join zastavky z ON z.id = lz.zastavka
			  join spoje s ON s.linka = l.id
			  join zastavky sz ON s.smer_jizdy = sz.id
			  join dny_jizdy dj ON s.dny_jizdy= dj.id
			  WHERE z.nazov_zastavky = ?;`

	rows, err := db.Query(query, stopName)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	var linesFromStop []LineFromStop
	var lineFS LineFromStop

	for rows.Next() {
		err := rows.Scan(&lineFS.LineName, &stopTime, &startTime, &lineFS.NextStop, &lineFS.Day)
		if err != nil {
			log.Fatal(err)
		}

		startTimeT, err := time.Parse(time.TimeOnly, startTime)
		if err != nil {
			log.Fatal("Error parsing start time:", err)
		}

		duration, err := time.Parse(time.TimeOnly, CalculateStopTime(startTime, stopTime))

		if err != nil {
			log.Fatal("Error parsing start time:", err)
		}

		endTime := startTimeT.Add(duration.Sub(time.Time{}))
		endTimeStr := endTime.Format(time.TimeOnly)

		lineFS.Time = endTimeStr
		linesFromStop = append(linesFromStop, lineFS)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return linesFromStop
}

func StopExists(stopName string) bool {
	query := `SELECT COUNT(*) FROM zastavky WHERE nazov_zastavky=?;`
	var cnt int

	err := db.QueryRow(query, stopName).Scan(&cnt)

	if err != nil {
		fmt.Printf("ERROR") //TODO
	}

	if cnt == 0 {
		return false
	} else {
		return true
	}
}
