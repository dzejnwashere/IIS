
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
                                                              ('553', 2, 1),
                                                              ('555', 32, 1),
                                                              ('557', 46, 1),
                                                              ('558', 14, 1),
                                                              ('559', 38, 1),
                                                              ('601', 50, 1),
                                                              ('603', 42, 1),
                                                              ('606', 14, 1),
                                                              ('607', 46, 1),
                                                              ('610', 49, 1),
                                                              ('613', 2, 1),

                                                              ('557', 2, 2),
                                                              ('559', 20, 2),
                                                              ('600', 15, 2),
                                                              ('601', 47, 2),
                                                              ('603', 18, 2),
                                                              ('604', 9, 2),
                                                              ('606', 13, 2),
                                                              ('608', 48, 2),
                                                              ('611', 13, 2),
                                                              ('612', 9, 2),
                                                              ('613', 23, 2),
                                                              ('614', 43, 2),
                                                              ('615', 32, 2),
                                                              ('616', 49, 2),
                                                              ('618', 2, 2),

                                                              ('545', 2, 3),
                                                              ('547', 33, 3),
                                                              ('549', 4, 3),
                                                              ('552', 1, 3),
                                                              ('553', 30, 3),
                                                              ('554', 21, 3),
                                                              ('556', 41, 3),
                                                              ('557', 27, 3),
                                                              ('558', 17, 3),
                                                              ('559', 19, 3),
                                                              ('600', 36, 3),
                                                              ('602', 1, 3),
                                                              ('604', 4, 3),
                                                              ('607', 24, 3),
                                                              ('610', 2, 3);


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
                                                                       (1, 1, '8:00', 1),
                                                                       (2, 10, '9:00', 3),
                                                                       (3, 1, '10:00', 2),
                                                                       (4, 5, '11:00', 5),
                                                                       (5, 14, '12:00', 4);

COMMIT;