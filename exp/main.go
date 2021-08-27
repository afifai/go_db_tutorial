package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Koneksi database MySQL")

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/mygodb")
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	// _, err = db.Exec(`
	// 			INSERT INTO users (usia, email, nama) VALUES (30, 'test@noreply.com', 'Rendi Darko');
	// 			`)
	// if err != nil {
	// 	panic(err.Error())
	// }

	var id int
	var nama, email string
	row := db.QueryRow(`
		SELECT id, nama, email
		FROM users
		WHERE id=?;`, 2)
	err = row.Scan(&id, &nama, &email)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Id", id, "Nama", nama, "email", email)

}
