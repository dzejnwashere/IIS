package db

import (
	"IIS/typedef"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func feed_zastavky() {
	query := `
	INSERT INTO zastavky (nazov_zastavky) VALUES
		('Štadión'),
		('Žel. stanica'),
		('ZŠ Dr. J. Dérera'),
		('Tesco'),
		('Nad výhonom'),
		('ZŠ Záhorácka'),
		('Gymnázium'),
		('Na majeri'),
		('Ludwiga Angerera'),
		('Plavecký Štvrtok'),
		('Vampil'),
		('Kozánek'),
		('Nový cintorín'),
		('Novomestského'),
		('Cintorín'),
		('Družstevná'),
		('Hurbanova ul.'),
		('Pri pekárni'),
		('Lesná'),
		('Pribinova'),
		('Rakárenská'),
		('Oslobodenia'),
		('Cesta mládeže'),
		('Pomník padlých'),
		('Tower'),
		('Záhorácka'),
		('Džbánkareň'),
		('Pasienky'),
		('Fritz'),
		('Nemocnica'),
		('Kúpalisko'),
		('Okresný súd'),
		('Sasinkova'),
		('Polesie'),
		('IKEA Boards'),
		('Ul. R. Dilonga'),
		('Kompostáreň'),
		('Zdravotné stredisko'),
		('IKEA Components'),
		('HSF'),
		('Písniky'),
		('Sídlisko Juh'),
		('Okresný úrad'),
		('Basso'),
		('Červený kríž'),
		('Jánošíkova'),
		('Kozia'),
		('Vinohrádok'),
		('Nám. SNP'),
		('M. R. Štefánika')
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

}

func feed_vozy() {
	query := `INSERT INTO vozy (spz, druh, znacka, kapacita) VALUES
              ('MA018DE', 'električka', 'Škoda', 118),
              ('MA369BA', 'električka', 'Škoda', 111),
              ('MA807PQ', 'autobus', 'Iveco', 45),
              ('MA382DB', 'autobus', 'Iveco', 57),
              ('MA103QQ', 'trolejbus', 'SOR', 68)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func feed_linky() {
	query := `INSERT INTO linky (nazev) VALUES
                                          (1),
                                          (2),
                                          (3),
                                          (10),
                                          (11)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func feed_linka_zastavka() {
	query := `INSERT INTO linka_zastavka (poradi, cas, zastavka, linka) VALUES
                                          (1, '1:00', 2, 1),
                                          (2, '1:05', 3, 1),
                                          (3, '1:07', 1, 1),
                                          (1, '1:20', 4, 2),
                                          (2, '1:30', 19, 2),
                                          (3, '1:40', 22, 2),
                                          (4, '1:50', 23, 2)
                                          `

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func create_users() {
	passHash, _ := bcrypt.GenerateFromPassword([]byte("spravce"), 10)
	id, _ := CreateOrUpdateUser(-1, "spravce", string(passHash), typedef.SpravcePerm)

	_, err := db.Exec(`INSERT INTO spravci (jmeno, prijmeni, user) VALUES ("Jožko", "Mrkvička", ?);`, id)
	if err != nil {
		log.Fatal(err)
	}

	passHash, _ = bcrypt.GenerateFromPassword([]byte("technik"), 10)
	id, _ = CreateOrUpdateUser(-1, "technik", string(passHash), typedef.TechnikPerm)

	_, err = db.Exec(`INSERT INTO technici (jmeno, prijmeni, user) VALUES ("Janko", "Hraško", ?);`, id)
	if err != nil {
		log.Fatal(err)
	}

	passHash, _ = bcrypt.GenerateFromPassword([]byte("ridic"), 10)
	id, _ = CreateOrUpdateUser(-1, "ridic", string(passHash), typedef.RidicPerm)

	_, err = db.Exec(`INSERT INTO ridici (jmeno, prijmeni, user) VALUES ("Ja", "Neviemuš", ?);`, id)
	if err != nil {
		log.Fatal(err)
	}

	passHash, _ = bcrypt.GenerateFromPassword([]byte("dispecer"), 10)
	id, _ = CreateOrUpdateUser(-1, "dispecer", string(passHash), typedef.DispecerPerm)

	_, err = db.Exec(`INSERT INTO dispeceri (jmeno, prijmeni, user) VALUES ("Disp", "Ečer", ?);`, id)
	if err != nil {
		log.Fatal(err)
	}
}

func feed_stav_zavady() {
	query := `INSERT INTO stav_zavady (stav) VALUES
                                          ('Vyřešeno'),
                                          ('V řešení'),
                                          ('Pozastaveno')
                                          `

	//1 vyřešeno
	//2 v řešení
	//3 pozastaveno

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func feed_zavady() {
	query := `INSERT INTO zavady (spz, autor, technik, popis, stav) VALUES
                                          ('MA018DE', 1, 1, 'defekt', 1),
                                          ('MA807PQ', 1, 1, 'porouchané přední dveře – nedovírají', 3),
                                          ('MA103QQ', 1, 1, 'nesvítí levé přední světlo', 2),
                                          ('MA103QQ', 1, 1, 'prasklé čelní sklo', 2)
                                          `

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func feed_tech_zaznamy() {
	query := `INSERT INTO tech_zaznamy (spz_vozidla, datum, zavada) VALUES
                                          ('MA018DE', '2023-11-18', NULL),
                                          ('MA103QQ', '2023-09-10', 4)
    `

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func feed_dny_jizdy() {
	query := `INSERT INTO dny_jizdy (den_jizdy) VALUES
                                          ('Všetky dni'),
                                          ('Pracovné dni'),
                                          ('Víkendy'),
                                          ('Sviatky'),
                                          ('Prázdniny')
                                          `

	//1 všetky dni
	//2 pracovné dni
	//3 víkendy
	//4 sviatky
	//5 prázdiny

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func feed_spoje() {
	query := `INSERT INTO spoje (linka, vuz, smer_jizdy, cas_odjezdu, dny_jizdy) VALUES
                                          (1, 'MA018DE', 1, '8:00', 1),
                                          (2, 'MA369BA', 10, '9:00', 3),
                                          (3, 'MA807PQ', 1, '10:00', 2),
                                          (4, 'MA382DB', 5, '11:00', 5),
                                          (5, 'MA103QQ', 14, '12:00', 4)
    `

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
