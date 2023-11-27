package db

import (
	"log"
	"strconv"
	"strings"
)

type Jizda_t struct {
	Id     int
	Driver int
	Spoj   int
	Vuz    string
	// optional params
	StartStop Stop_t
	EndStop   Stop_t
	StartTime string
	EndTime   string
	LineName  string
}

func CalculateStopTime(Departure string, duration int) string {
	split := strings.Split(Departure, ":")
	hours, _ := strconv.Atoi(split[0])
	mins, _ := strconv.Atoi(split[1])
	if ((hours*60 + mins + duration) % 60) < 10 {
		return strconv.Itoa((hours*60+mins+duration)/60) + ":0" + strconv.Itoa((hours*60+mins+duration)%60) + ":00"
	}
	return strconv.Itoa((hours*60+mins+duration)/60) + ":" + strconv.Itoa((hours*60+mins+duration)%60) + ":00"

}

func GetMyRides(driver int) []Jizda_t {
	var jizdaTemp Jizda_t
	var ret []Jizda_t

	rows, err := db.Query("SELECT id, spz, spoj, ridic from jizda where ridic = ?", driver)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&jizdaTemp.Id, &jizdaTemp.Vuz, &jizdaTemp.Spoj, &jizdaTemp.Driver)

		if err != nil {
			log.Fatal(err)
		}

		spoj := GetSpoj(jizdaTemp.Spoj)
		jizdaTemp.StartTime = spoj.CasOdjezdu
		stops := GetLineStops(spoj.Linka)
		line := GetLine(spoj.Linka)
		jizdaTemp.LineName = line.Name
		atoi, err := strconv.Atoi(stops[len(stops)-1].Time)
		jizdaTemp.EndTime = CalculateStopTime(spoj.CasOdjezdu, atoi)
		if spoj.PrimarniSmer {
			jizdaTemp.StartStop = Stop_t{
				Id:   stops[0].Stop_id,
				Name: stops[0].Stop_name,
			}
			jizdaTemp.EndStop = Stop_t{
				Id:   stops[len(stops)-1].Stop_id,
				Name: stops[len(stops)-1].Stop_name,
			}
		} else {
			jizdaTemp.EndStop = Stop_t{
				Id:   stops[0].Stop_id,
				Name: stops[0].Stop_name,
			}
			jizdaTemp.StartStop = Stop_t{
				Id:   stops[len(stops)-1].Stop_id,
				Name: stops[len(stops)-1].Stop_name,
			}
		}
		ret = append(ret, jizdaTemp)
	}
	return ret
}
