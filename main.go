package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	db_user   = "root"
	db_passwd = "password"
	db_addr   = "localhost"
	db_port   = "3306"
	db_db     = "mygodb"
)

type Person struct {
	Id     int
	Nama   string
	Usia   int
	Lokasi string
}

func main() {
	fmt.Println("Koneksi MySQL ke Go app")
	s := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_user, db_passwd, db_addr, db_port, db_db)
	fmt.Println(s)
	db, err := sql.Open("mysql", s)
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	// err = insertData(db)
	// if err != nil {
	// 	panic(err.Error())
	// }

	people, err := getAllData(db)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(people)

	err = deleteCondionalData(db, 25)
	if err != nil {
		panic(err.Error())
	}
	people, err = getAllData(db)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(people)

	err = updateUsia(db, "Afif A. Iskandar", 30)
	if err != nil {
		panic(err.Error())
	}
	people, err = getAllData(db)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(people)

}

func insertData(db *sql.DB) error {
	people := GetData()
	for _, person := range people {
		q := "INSERT INTO biodata (nama, usia, lokasi) VALUES (?, ?, ?)"
		insert, err := db.Prepare(q)
		defer insert.Close()

		if err != nil {
			panic(err.Error())
		}

		_, err = insert.Exec(person.Nama, person.Usia, person.Lokasi)

		if err != nil {
			panic(err.Error())
		}
	}
	return nil
}

func getAllData(db *sql.DB) (people []Person, err error) {
	resp, err := db.Query("SELECT * FROM biodata")
	defer resp.Close()
	if err != nil {
		panic(err.Error())
	}
	for resp.Next() {
		var pPerson Person
		err = resp.Scan(&pPerson.Id, &pPerson.Nama, &pPerson.Usia, &pPerson.Lokasi)
		if err != nil {
			panic(err.Error())
		}
		people = append(people, pPerson)
	}
	return people, nil
}

func deleteCondionalData(db *sql.DB, usia int) error {
	q := "DELETE FROM biodata WHERE usia > ?"
	drop, err := db.Prepare(q)
	defer drop.Close()
	if err != nil {
		panic(err.Error())
	}

	_, err = drop.Exec(usia)
	if err != nil {
		panic(err.Error())
	}
	return nil
}

func updateUsia(db *sql.DB, nama string, usia int) error {
	q := "UPDATE biodata SET usia = ? WHERE nama like ?"
	update, err := db.Prepare(q)
	defer update.Close()
	if err != nil {
		panic(err.Error())
	}

	_, err = update.Exec(usia, nama)
	if err != nil {
		panic(err.Error())
	}
	return nil
}
