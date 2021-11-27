CREATE SCHEMA commusto;
SET search_path = commusto;
CREATE TABLE users (
    uid VARCHAR(32) NOT NULL, 
    mail VARCHAR(32) NOT NULL,
    password VARCHAR(32) NOT NULL,
    name VARCHAR(32) NOT NULL,
    PRIMARY KEY (uid)
);
