
START TRANSACTION ;


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
                                          ('ZŠ Štúrova'),
                                          ('M. R. Štefánika');

INSERT INTO linky (nazev) VALUES
                              (1),
                              (2),
                              (3),
                              (10),
                              (11);

INSERT INTO linka_zastavka (cas, zastavka, linka) VALUES
                                                              (0, 2, 1),
                                                              (200, 32, 1),
                                                              (400, 46, 1),
                                                              (500, 14, 1),
                                                              (600, 38, 1),
                                                              (800, 50, 1),
                                                              (1000, 42, 1),
                                                              (1300, 14, 1),
                                                              (1400, 46, 1),
                                                              (1700, 49, 1),
                                                              (2000, 2, 1),

                                                              (0, 2, 2),
                                                              (200, 20, 2),
                                                              (300, 15, 2),
                                                              (400, 47, 2),
                                                              (600, 18, 2),
                                                              (700, 9, 2),
                                                              (900, 13, 2),
                                                              (1100, 48, 2),
                                                              (1400, 13, 2),
                                                              (1500, 9, 2),
                                                              (1600, 23, 2),
                                                              (1700, 43, 2),
                                                              (1800, 32, 2),
                                                              (1900, 49, 2),
                                                              (2100, 2, 2),

                                                              (0, 2, 3),
                                                              (200, 33, 3),
                                                              (400, 4, 3),
                                                              (700, 1, 3),
                                                              (800, 30, 3),
                                                              (900, 21, 3),
                                                              (1100, 41, 3),
                                                              (1300, 27, 3),
                                                              (1400, 17, 3),
                                                              (1500, 19, 3),
                                                              (1600, 36, 3),
                                                              (1800, 1, 3),
                                                              (2000, 4, 3),
                                                              (2300, 24, 3),
                                                              (2600, 2, 3);


INSERT INTO vozy (spz, druh, znacka, kapacita) VALUES
                                                   ('MA018DE', 'električka', 'Škoda', 118),
                                                   ('MA369BA', 'električka', 'Škoda', 111),
                                                   ('MA807PQ', 'autobus', 'Iveco', 45),
                                                   ('MA382DB', 'autobus', 'Iveco', 57),
                                                   ('MA103QQ', 'trolejbus', 'SOR', 68);

INSERT INTO stav_zavady (stav) VALUES
                                   ('Nepřideleno'),
                                   ('Přideleno'),
                                   ('V řešení'),
                                   ('Pozastaveno'),
                                   ('Vyřešeno');

INSERT INTO zavady (spz, autor, technik, popis, stav) VALUES
                                                          ('MA018DE', 2, 5, 'defekt', 1),
                                                          ('MA807PQ', 2, 5, 'porouchané přední dveře – nedovírají', 3),
                                                          ('MA103QQ', 2, 5, 'nesvítí levé přední světlo', 2),
                                                          ('MA103QQ', 2, 5, 'prasklé čelní sklo', 2);

INSERT INTO tech_zaznamy (spz_vozidla, datum, zavada, popis, autor) VALUES
                                                          ('MA018DE', '2023-11-18', NULL, 'testi sakj as', 5),
                                                          ('MA103QQ', '2023-09-10', 4, 'bla vla sldkslkaed', 5);

INSERT INTO dny_jizdy (den_jizdy) VALUES
                                      ('Všetky dni'),
                                      ('Pracovné dni'),
                                      ('Víkendy'),
                                      ('Sviatky'),
                                      ('Prázdniny');

INSERT INTO spoje (linka, smer_jizdy, cas_odjezdu, dny_jizdy) VALUES
                                                                       (1, 1, '5:53', 1),
                                                                       (1, 1, '6:15', 2),
                                                                       (1, 1, '7:10', 1),
                                                                       (1, 1, '7:32', 2),
                                                                       (1, 1, '8:15', 1),
                                                                       (1, 1, '9:15', 2),
                                                                       (1, 1, '12:50', 1),
                                                                       (1, 1, '13:15', 2),
                                                                       (1, 1, '13:45', 1),
                                                                       (1, 1, '14:10', 2),

                                                                       (2, 10, '9:00', 3),
                                                                       (3, 1, '10:00', 2),
                                                                       (4, 5, '11:00', 5),
                                                                       (5, 14, '12:00', 4);

COMMIT;