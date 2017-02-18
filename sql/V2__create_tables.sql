CREATE TABLE IF NOT EXISTS USERS
(
  ID UUID PRIMARY KEY DEFAULT (uuid_generate_v1()),
  FIRST_NAME VARCHAR(255) NOT NULL,
  LAST_NAME VARCHAR(255) NOT NULL,
  EMAIL VARCHAR(255) NOT NULL UNIQUE,
  USERNAME VARCHAR(255) NOT NULL UNIQUE,
  PASSWORD VARCHAR(255) NOT NULL,
  DATE_CREATED TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  LAST_UPDATED TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS GAME
(
  ID UUID PRIMARY KEY DEFAULT (uuid_generate_v1()),
  TITLE VARCHAR(255) NOT NULL,
  CONSOLE VARCHAR(255) NOT NULL,
  RELEASE_DATE DATE NOT NULL,
  DEVELOPER VARCHAR(255) NOT NULL,
  PUBLISHER VARCHAR(255) NOT NULL,
  GENRE VARCHAR(255) NOT NULL,
  PLAYERS SMALLINT NOT NULL,
  PRICE DECIMAL(6,2) NOT NULL,
  DATE_CREATED TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  LAST_UPDATED TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS PURCHASE
(
  ID UUID PRIMARY KEY DEFAULT (uuid_generate_v1()),
  USER_ID UUID NOT NULL,
  GAME_ID UUID NOT NULL,
  PURCHASE_DATE DATE NOT NULL,
  DATE_CREATED TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  LAST_UPDATED TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);