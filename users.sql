CREATE TABLE users (
    id int unsigned NOT NULL AUTO_INCREMENT,
    nama text,
    usia int,
    email VARCHAR(30) UNIQUE NOT NULL,
    PRIMARY KEY (id)
);