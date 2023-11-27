
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
                                                              (2, 32, 1),
                                                              (4, 46, 1),
                                                              (5, 14, 1),
                                                              (6, 38, 1),
                                                              (8, 50, 1),
                                                              (10, 42, 1),
                                                              (13, 14, 1),
                                                              (14, 46, 1),
                                                              (17, 49, 1),

                                                              (0, 2, 2),
                                                              (2, 20, 2),
                                                              (3, 15, 2),
                                                              (4, 47, 2),
                                                              (6, 18, 2),
                                                              (7, 9, 2),
                                                              (9, 13, 2),
                                                              (11, 48, 2),
                                                              (14, 13, 2),
                                                              (15, 9, 2),
                                                              (16, 23, 2),
                                                              (17, 43, 2),
                                                              (18, 32, 2),
                                                              (19, 49, 2),

                                                              (0, 2, 3),
                                                              (2, 33, 3),
                                                              (4, 4, 3),
                                                              (7, 1, 3),
                                                              (8, 30, 3),
                                                              (9, 21, 3),
                                                              (11, 41, 3),
                                                              (13, 27, 3),
                                                              (14, 17, 3),
                                                              (15, 19, 3),
                                                              (16, 36, 3),
                                                              (18, 1, 3),
                                                              (20, 4, 3),
                                                              (23, 24, 3);


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
                                                                       (1, 49, '5:53', 2),
                                                                       (1, 49, '6:15', 2),
                                                                       (1, 49, '7:10', 2),
                                                                       (1, 49, '7:32', 2),
                                                                       (1, 49, '8:15', 2),
                                                                       (1, 49, '9:15', 2),
                                                                       (1, 49, '12:50', 2),
                                                                       (1, 49, '13:15', 2),
                                                                       (1, 49, '13:45', 2),
                                                                       (1, 49, '14:10', 1),

                                                                       (1, 2, '5:59', 2),
                                                                       (1, 2, '6:19', 2),
                                                                       (1, 2, '7:13', 2),
                                                                       (1, 2, '7:38', 2),
                                                                       (1, 2, '8:17', 2),
                                                                       (1, 2, '9:19', 2),
                                                                       (1, 2, '12:53', 2),
                                                                       (1, 2, '13:15', 2),
                                                                       (1, 2, '13:43', 2),
                                                                       (1, 2, '14:17', 1),

                                                                       (2, 10, '9:00', 3),
                                                                       (3, 1, '10:00', 2),
                                                                       (4, 5, '11:00', 5),
                                                                       (5, 14, '12:00', 4);

INSERT INTO jizda (spz, spoj, ridic) VALUES
                                       ('MA018DE', 1, 1),
                                       ('MA807PQ', 2, 1),
                                       ('MA103QQ', 3, 4),
                                       ('MA382DB', 4, 1),
                                       ('MA369BA', 5, 1);


COMMIT;