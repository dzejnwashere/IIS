package db

import "log"

type Spoj_t struct {
	Id           int
	Linka        int
	CasOdjezdu   string
	PrimarniSmer bool
	DenJizdy     int
}

func GetAllSpoje() []Spoj_t {
	var spoj_temp Spoj_t
	var smj int
	var spoje []Spoj_t
	rows, err := db.Query("SELECT id, linka, cas_odjezdu, smer_jizdy, dny_jizdy from spoje")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&(spoj_temp.Id), &spoj_temp.Linka, &spoj_temp.CasOdjezdu, &smj, &spoj_temp.DenJizdy)

		if err != nil {
			log.Fatal(err)
		}
		spoj_temp.PrimarniSmer = smj == 0
		spoje = append(spoje, spoj_temp)
	}
	return spoje
}

func GetSpoj(spojId int) Spoj_t {
	var spoj_temp Spoj_t
	var smj int
	rows := db.QueryRow("SELECT id, linka, cas_odjezdu, smer_jizdy, dny_jizdy from spoje where id = ?", spojId)

	err := rows.Scan(&(spoj_temp.Id), &spoj_temp.Linka, &spoj_temp.CasOdjezdu, &smj, &spoj_temp.DenJizdy)

	if err != nil {
		log.Fatal(err)
	}
	spoj_temp.PrimarniSmer = smj == 0
	return spoj_temp
}

func GetSpojeByLine(lineId int) []Spoj_t {
	var spoj_temp Spoj_t
	var smj int
	var spoje []Spoj_t
	rows, err := db.Query("SELECT id, linka, cas_odjezdu, smer_jizdy, dny_jizdy from spoje where linka = ?", lineId)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&(spoj_temp.Id), &spoj_temp.Linka, &spoj_temp.CasOdjezdu, &smj, &spoj_temp.DenJizdy)

		if err != nil {
			log.Fatal(err)
		}
		spoj_temp.PrimarniSmer = smj == 0
		spoje = append(spoje, spoj_temp)
	}
	return spoje
}

type DenJizdy_t struct {
	Id    int
	Popis string
}

func GetAllDnyJizdy() []DenJizdy_t {
	var denJizdyTemp DenJizdy_t
	var ret []DenJizdy_t
	rows, err := db.Query("SELECT id, den_jizdy from dny_jizdy")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&denJizdyTemp.Id, &denJizdyTemp.Popis)

		if err != nil {
			log.Fatal(err)
		}
		ret = append(ret, denJizdyTemp)
	}
	return ret
}

func DeleteSpoj(id int) error {
	_, err := db.Exec("DELETE FROM spoje where id=?", id)
	return err
}

func CreateSpoj(spoj Spoj_t) error {
	var smj int
	if spoj.PrimarniSmer {
		smj = 0
	} else {
		smj = 1
	}
	_, err := db.Exec("INSERT INTO spoje (linka, cas_odjezdu, smer_jizdy, dny_jizdy) values (?, ?, ?, ?)", spoj.Linka, spoj.CasOdjezdu, smj, spoj.DenJizdy)
	return err
}

func UpdateSpoj(spoj Spoj_t) error {
	var smj int
	if spoj.PrimarniSmer {
		smj = 0
	} else {
		smj = 1
	}
	_, err := db.Exec("UPDATE spoje SET linka = ?, cas_odjezdu = ?, smer_jizdy = ?, dny_jizdy = ? where id = ?", spoj.Linka, spoj.CasOdjezdu, smj, spoj.DenJizdy, spoj.Id)
	return err
}
