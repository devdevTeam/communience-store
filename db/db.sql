CREATE SCHEMA commusto;
SET search_path = commusto;
CREATE TABLE users (
    uid VARCHAR(32) NOT NULL, 
    mail VARCHAR(32) NOT NULL,
    password VARCHAR(32) NOT NULL,
    name VARCHAR(32) NOT NULL,
    PRIMARY KEY (uid)
);

CREATE TABLE rooms (
    rid VARCHAR(64) NOT NULL, 
    name VARCHAR(32) NOT NULL,
    password VARCHAR(32) NOT NULL,
    PRIMARY KEY (rid)
);

CREATE TABLE user_room_relation (
    admin BOOLEAN NOT NULL,
    uid VARCHAR(32) NOT NULL,
    rid VARCHAR(64) NOT NULL,
    FOREIGN KEY (uid) REFERENCES users(uid) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (rid) REFERENCES rooms(rid) ON DELETE CASCADE ON UPDATE CASCADE
)
