CREATE SCHEMA commusto;
SET search_path = commusto;
CREATE TABLE users (
    uid VARCHAR(32) NOT NULL, 
    mail VARCHAR(32) NOT NULL,
    password VARCHAR(32) NOT NULL,
    name VARCHAR(32) NOT NULL,
    hash VARCHAR(64) NOT NULL,
    PRIMARY KEY (uid)
);

CREATE TABLE user_friend_relation (
    uid VARCHAR(32) NOT NULL,
    fid VARCHAR(32) NOT NULL,
    FOREIGN KEY (uid) REFERENCES users(uid) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (fid) REFERENCES users(uid) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE default_cards (
  name VARCHAR(64),
  hurigana VARCHAR(64),
  birthday VARCHAR(64),
  instagram VARCHAR(64),
  twitter VARCHAR(64),
  facebook VARCHAR(64),
  free text,
  hobby text,
  uid VARCHAR(32) NOT NULL,
  FOREIGN KEY (uid) REFERENCES users(uid) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE rooms (
    rid VARCHAR(64) NOT NULL, 
    name VARCHAR(32) NOT NULL,
    password VARCHAR(32) NOT NULL,
    have_form BOOLEAN,
    hash VARCHAR(64) NOT NULL,
    PRIMARY KEY (rid)
);

CREATE TABLE user_room_relation (
    admin BOOLEAN NOT NULL,
    uid VARCHAR(32) NOT NULL,
    rid VARCHAR(64) NOT NULL,
    FOREIGN KEY (uid) REFERENCES users(uid) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (rid) REFERENCES rooms(rid) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE forms (
    rid VARCHAR(64) NOT NULL,
    col_name VARCHAR(32) NOT NULL,
    col_idx INT NOT NULL,
    display BOOLEAN NOT NULL,
    FOREIGN KEY (rid) REFERENCES rooms(rid) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE events (
    eid VARCHAR(64) NOT NULL,
    org_uid VARCHAR(32) NOT NULL,
    password VARCHAR(32) NOT NULL,
    rid VARCHAR(64) NOT NULL,
    PRIMARY KEY (eid),
    FOREIGN KEY (rid) REFERENCES rooms(rid) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE event_col (
    eid VARCHAR(64) NOT NULL,
    col_name VARCHAR(32) NOT NULL,
    col_idx INT NOT NULL,
    hidden BOOLEAN NOT NULL,
    FOREIGN KEY (eid) REFERENCES events(eid) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE participants (
    eid VARCHAR(64) NOT NULL,
    uid VARCHAR(32) NOT NULL,
    FOREIGN KEY (eid) REFERENCES events(eid) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (uid) REFERENCES users(uid) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE card_value (
    uid VARCHAR(32) NOT NULL,
    rid VARCHAR(64) NOT NULL,
    value VARCHAR(128),
    col_idx INT NOT NULL,
    FOREIGN KEY (uid) REFERENCES users(uid) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (rid) REFERENCES rooms(rid) ON DELETE CASCADE ON UPDATE CASCADE
)