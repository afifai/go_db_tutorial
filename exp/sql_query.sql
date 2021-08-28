CREATE TABLE transaksi (
    id int unsigned NOT NULL AUTO_INCREMENT,
    user_id int,
    credit int,
    PRIMARY KEY (id)
);

CREATE TABLE user_details (
    id int unsigned NOT NULL AUTO_INCREMENT,
    nama text,
    PRIMARY KEY (id)
);

INSERT INTO transaksi 
    (user_id, credit)
VALUES
    (1, 1000000),
    (1, 500000),
    (1, 3000000),
    (1, 170000),
    (1, 125000),
    (2, 2100000),
    (2, 175000),
    (2, 225000),
    (3, 2150000),
    (3, 185000),
    (3, 525000),
    (5, 100000),
    (6, 700000);

INSERT INTO user_details
    (nama)
VALUES
    ("Afif A. Iskandar"),
    ("Shafa Siregar"),
    ("Budi Nugroho"),
    ("Lisa Blekping");

SELECT t1.nama, t2.credit FROM user_details AS t1 INNER JOIN transaksi AS t2 ON t1.id = t2.user_id;
SELECT t1.nama, t2.credit FROM user_details AS t1 LEFT JOIN transaksi AS t2 ON t1.id = t2.user_id;
SELECT t1.nama, t2.credit FROM user_details AS t1 RIGHT JOIN transaksi AS t2 ON t1.id = t2.user_id;
