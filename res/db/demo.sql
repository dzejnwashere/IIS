
START TRANSACTION ;

INSERT INTO `users` (`id`, `username`, `password`, `permissions`) VALUES
    (2, 'spravce', '$2a$10$/wDKWSx1/ukpHb7x2QWkWez9924v42yNpEq2Ro6TPTNULSr9a90A2', 16),
    (3, 'technik', '$2a$10$9lQD3rr3WOfQxvRhwtIs9uXUuusOuvmaOrXjoWx/z14ZRNM.1IYdq', 8),
    (4, 'ridic', '$2a$10$WTtg3lJpFbiqXAQd.v20r.MbNNIfRctpppct3V.BIVBQqO1fXxchu', 2),
    (5, 'dispecer', '$2a$10$03lureaEspzJWWw6dx6PtOYQH.SI4kaUC1tyVJyLC4KTLlHAAv0TK', 4);

INSERT INTO spravci (jmeno, prijmeni, user) VALUES ("Jožko", "Mrkvička", 2);

INSERT INTO technici (jmeno, prijmeni, user) VALUES ("Janko", "Hraško", 3);

INSERT INTO ridici (jmeno, prijmeni, user) VALUES ("Ja", "Neviemuš", 4);

INSERT INTO dispeceri (jmeno, prijmeni, user) VALUES ("Disp", "Ečer", 5);

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
                                          ('M. R. Štefánika');

INSERT INTO linky (nazev) VALUES
                              (1),
                              (2),
                              (3),
                              (10),
                              (11);

INSERT INTO linka_zastavka (poradi, cas, zastavka, linka) VALUES
                                                              (1, '1:00', 2, 1),
                                                              (2, '1:05', 3, 1),
                                                              (3, '1:07', 1, 1),
                                                              (1, '1:20', 4, 2),
                                                              (2, '1:30', 19, 2),
                                                              (3, '1:40', 22, 2),
                                                              (4, '1:50', 23, 2);

INSERT INTO vozy (spz, druh, znacka, kapacita) VALUES
                                                   ('MA018DE', 'električka', 'Škoda', 118),
                                                   ('MA369BA', 'električka', 'Škoda', 111),
                                                   ('MA807PQ', 'autobus', 'Iveco', 45),
                                                   ('MA382DB', 'autobus', 'Iveco', 57),
                                                   ('MA103QQ', 'trolejbus', 'SOR', 68);

INSERT INTO stav_zavady (stav) VALUES
                                   ('Vyřešeno'),
                                   ('V řešení'),
                                   ('Pozastaveno');

INSERT INTO zavady (spz, autor, technik, popis, stav) VALUES
                                                          ('MA018DE', 2, 3, 'defekt', 1),
                                                          ('MA807PQ', 2, 3, 'porouchané přední dveře – nedovírají', 3),
                                                          ('MA103QQ', 2, 3, 'nesvítí levé přední světlo', 2),
                                                          ('MA103QQ', 2, 3, 'prasklé čelní sklo', 2);

INSERT INTO tech_zaznamy (spz_vozidla, datum, zavada) VALUES
                                                          ('MA018DE', '2023-11-18', NULL),
                                                          ('MA103QQ', '2023-09-10', 4);

INSERT INTO dny_jizdy (den_jizdy) VALUES
                                      ('Všetky dni'),
                                      ('Pracovné dni'),
                                      ('Víkendy'),
                                      ('Sviatky'),
                                      ('Prázdniny');

INSERT INTO spoje (linka, vuz, smer_jizdy, cas_odjezdu, dny_jizdy) VALUES
                                                                       (1, 'MA018DE', 1, '8:00', 1),
                                                                       (2, 'MA369BA', 10, '9:00', 3),
                                                                       (3, 'MA807PQ', 1, '10:00', 2),
                                                                       (4, 'MA382DB', 5, '11:00', 5),
                                                                       (5, 'MA103QQ', 14, '12:00', 4);

COMMIT;