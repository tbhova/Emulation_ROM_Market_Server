INSERT INTO USERS (ID, FIRST_NAME, LAST_NAME, EMAIL, USERNAME, PASSWORD) VALUES (
    '26fe1f9e-25e6-11e7-93ae-92361f002671', 'First', 'Last', 'games@horganes.com', 'download', 'password'
);

INSERT INTO GAME (ID, TITLE, CONSOLE, RELEASE_DATE, DEVELOPER, PUBLISHER, GENRE, PLAYERS, PRICE, S3LINK) VALUES
    ('4e16ad2a-25e7-11e7-93ae-92361f002671', 'The Legend of Zelda', 'nes', TO_DATE('2016-04-01', 'YYYY-MM-DD'), '', '', '', 1, 4.99, 'Legend of Zelda, The.nes'),
    ('5c9e39d0-25e7-11e7-93ae-92361f002671', 'Mario Bros.', 'nes', TO_DATE('2016-04-01', 'YYYY-MM-DD'), '', '', '', 1, 4.99, 'Mario Bros..nes'),
    ('6ac00d0e-25e7-11e7-93ae-92361f002671', 'Donkey Kong', 'nes', TO_DATE('2016-04-01', 'YYYY-MM-DD'), '', '', '', 1, 4.99, 'Donkey Kong.nes'),
    ('74bfab5c-25e7-11e7-93ae-92361f002671', 'Donkey Kong Country', 'snes', TO_DATE('2016-04-01', 'YYYY-MM-DD'), '', '', '', 1, 4.99, 'Donkey Kong Country.smc'),
    ('2450bdda-25ee-11e7-93ae-92361f002671', 'The Legend of Zelda - A Link to the Past', 'snes', TO_DATE('2016-04-01', 'YYYY-MM-DD'), '', '', '', 1, 4.99, 'Legend of Zelda, The - A Link to the Past.smc'),
    ('2f949860-25ee-11e7-93ae-92361f002671', 'Super Mario Kart', 'snes', TO_DATE('2016-04-01', 'YYYY-MM-DD'), '', '', '', 1, 4.99, 'Super Mario Kart.smc'),
    ('340b9416-25ee-11e7-93ae-92361f002671', '', '', TO_DATE('2016-04-01', 'YYYY-MM-DD'), '', '', '', 1, 4.99, ''),
    ('40e5821e-25ee-11e7-93ae-92361f002671', '', '', TO_DATE('2016-04-01', 'YYYY-MM-DD'), '', '', '', 1, 4.99, ''),
    ('4b96d7c6-25ee-11e7-93ae-92361f002671', '', '', TO_DATE('2016-04-01', 'YYYY-MM-DD'), '', '', '', 1, 4.99, ''),
    ('52666ff8-25ee-11e7-93ae-92361f002671', '', '', TO_DATE('2016-04-01', 'YYYY-MM-DD'), '', '', '', 1, 4.99, ''),
    ('67d2781e-25ee-11e7-93ae-92361f002671', '', '', TO_DATE('2016-04-01', 'YYYY-MM-DD'), '', '', '', 1, 4.99, ''),
    ('7033551e-25ee-11e7-93ae-92361f002671', '', '', TO_DATE('2016-04-01', 'YYYY-MM-DD'), '', '', '', 1, 4.99, ''),
    ('738f20d0-25ee-11e7-93ae-92361f002671', 'Spyro the Dragon', 'psx', TO_DATE('2016-04-01', 'YYYY-MM-DD'), '', '', '', 1, 4.99, 'Spyro the Dragon.bin/Spyron the Dragon.cue'),
    ('7db45ed6-25ee-11e7-93ae-92361f002671', 'Asteroids.a78', 'atari7800', TO_DATE('2016-04-01', 'YYYY-MM-DD'), '', '', '', 1, 4.99, 'Asteroids.a78'),
    ('85fb915e-25ee-11e7-93ae-92361f002671', 'Pole Position II', 'atari7800', TO_DATE('2016-04-01', 'YYYY-MM-DD'), '', '', '', 1, 4.99, 'Pole Position II.a78');

INSERT INTO PURCHASE (USER_ID, GAME_ID, PURCHASE_DATE) VALUES
    ('26fe1f9e-25e6-11e7-93ae-92361f002671', '4e16ad2a-25e7-11e7-93ae-92361f002671', TO_DATE('2017-03-20', 'YYYY-MM-DD')),
    ('26fe1f9e-25e6-11e7-93ae-92361f002671', '5c9e39d0-25e7-11e7-93ae-92361f002671', TO_DATE('2017-03-20', 'YYYY-MM-DD')),
    ('26fe1f9e-25e6-11e7-93ae-92361f002671', '6ac00d0e-25e7-11e7-93ae-92361f002671', TO_DATE('2017-03-20', 'YYYY-MM-DD')),
    ('26fe1f9e-25e6-11e7-93ae-92361f002671', '74bfab5c-25e7-11e7-93ae-92361f002671', TO_DATE('2017-03-20', 'YYYY-MM-DD')),
    ('26fe1f9e-25e6-11e7-93ae-92361f002671', '2450bdda-25ee-11e7-93ae-92361f002671', TO_DATE('2017-03-20', 'YYYY-MM-DD')),
    ('26fe1f9e-25e6-11e7-93ae-92361f002671', '2f949860-25ee-11e7-93ae-92361f002671', TO_DATE('2017-03-20', 'YYYY-MM-DD')),
    ('26fe1f9e-25e6-11e7-93ae-92361f002671', '340b9416-25ee-11e7-93ae-92361f002671', TO_DATE('2017-03-20', 'YYYY-MM-DD')),
    ('26fe1f9e-25e6-11e7-93ae-92361f002671', '40e5821e-25ee-11e7-93ae-92361f002671', TO_DATE('2017-03-20', 'YYYY-MM-DD')),
    ('26fe1f9e-25e6-11e7-93ae-92361f002671', '4b96d7c6-25ee-11e7-93ae-92361f002671', TO_DATE('2017-03-20', 'YYYY-MM-DD')),
    ('26fe1f9e-25e6-11e7-93ae-92361f002671', '52666ff8-25ee-11e7-93ae-92361f002671', TO_DATE('2017-03-20', 'YYYY-MM-DD')),
    ('26fe1f9e-25e6-11e7-93ae-92361f002671', '67d2781e-25ee-11e7-93ae-92361f002671', TO_DATE('2017-03-20', 'YYYY-MM-DD')),
    ('26fe1f9e-25e6-11e7-93ae-92361f002671', '738f20d0-25ee-11e7-93ae-92361f002671', TO_DATE('2017-03-20', 'YYYY-MM-DD')),
    ('26fe1f9e-25e6-11e7-93ae-92361f002671', '7033551e-25ee-11e7-93ae-92361f002671', TO_DATE('2017-03-20', 'YYYY-MM-DD')),
    ('26fe1f9e-25e6-11e7-93ae-92361f002671', '7db45ed6-25ee-11e7-93ae-92361f002671', TO_DATE('2017-03-20', 'YYYY-MM-DD')),
    ('26fe1f9e-25e6-11e7-93ae-92361f002671', '85fb915e-25ee-11e7-93ae-92361f002671', TO_DATE('2017-03-20', 'YYYY-MM-DD'));
