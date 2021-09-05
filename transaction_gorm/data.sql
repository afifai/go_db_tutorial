CREATE TABLE produk (
    id int unsigned NOT NULL AUTO_INCREMENT,
    nama_produk text,
    harga int,
    stok int,
    PRIMARY KEY (id)
);

START TRANSACTION;
INSERT INTO produk
    (nama_produk, harga, stok)
VALUES
    ('madu', 100000, 100);

INSERT INTO produk
    (nama_produk, harga, stok)
VALUES
    ('garam', 500, 'n');



DELIMITER $
CREATE PROCEDURE tambahDataTiga(IN in_harga INT)
BEGIN
    DECLARE EXIT HANDLER FOR SQLEXCEPTION, SQLWARNING ROLLBACK;
    START TRANSACTION;
        INSERT INTO produk(nama_produk, harga, stok) VALUES ('roti', in_harga, 100);
        SELECT * FROM produk;
        INSERT INTO produk(nama_produk, harga, stok) VALUES ('biskuit', in_harga, 'n');
        SELECT * FROM produk;
    COMMIT;
END $
DELIMITER ;