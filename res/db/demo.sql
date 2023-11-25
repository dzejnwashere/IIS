
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
                                          ('M. R. Štefánika');

INSERT INTO linky (nazev) VALUES
                              (1),
                              (2),
                              (3),
                              (10),
                              (11);

INSERT INTO linka_zastavka (cas, zastavka, linka) VALUES
                                                              ('100', 2, 1),
                                                              ('105', 3, 1),
                                                              ('107', 1, 1),
                                                              ('120', 4, 2),
                                                              ('130', 19, 2),
                                                              ('140', 22, 2),
                                                              ('150', 23, 2);

INSERT INTO vozy (spz, druh, znacka, kapacita) VALUES
                                                   ('MA018DE', 'električka', 'Škoda', 118),
                                                   ('MA369BA', 'električka', 'Škoda', 111),
                                                   ('MA807PQ', 'autobus', 'Iveco', 45),
                                                   ('MA382DB', 'autobus', 'Iveco', 57),
                                                   ('MA103QQ', 'trolejbus', 'SOR', 68);

INSERT INTO stav_zavady (stav) VALUES
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

INSERT INTO spoje (linka, vuz, smer_jizdy, cas_odjezdu, dny_jizdy) VALUES
                                                                       (1, 'MA018DE', 1, '8:00', 1),
                                                                       (2, 'MA369BA', 10, '9:00', 3),
                                                                       (3, 'MA807PQ', 1, '10:00', 2),
                                                                       (4, 'MA382DB', 5, '11:00', 5),
                                                                       (5, 'MA103QQ', 14, '12:00', 4);

COMMIT;