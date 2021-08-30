package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const (
	db_user   = "root"
	db_passwd = "password"
	db_addr   = "localhost"
	db_port   = "3306"
	db_db     = "gosql_handson"
)

type Employee struct {
	Id      int
	Name    string
	Email   string
	Address string
	Age     int
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

	var pilihan, input_int int
	var input_mod string
	fmt.Println("Selamat datang di program CRUD sederhana menggunakan GO")
	fmt.Println("Menu : ")
	fmt.Println("1. Masukkan Data\n2. Lihat Data\n3. Update Data\n4. Hapus Data")
	fmt.Print("Masukkan pilihan Anda : ")
	reader := bufio.NewReader(os.Stdin)
	fmt.Scanf("%d", &pilihan)
	if pilihan == 1 {
		name, email, address, age := getInputData()
		new_emp := Employee{Name: name, Email: email, Address: address, Age: age}
		err = insertData(db, new_emp)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Data Berhasil Masuk")
	} else if pilihan == 2 {
		employee, err := getAllData(db)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(employee)
	} else if pilihan == 3 {
		fmt.Print("Masukkan email pegawai yang ingin diupdate: ")
		email, _ := reader.ReadString('\n')
		email = strings.TrimSpace(email)
		selected_emp, err := getSelectedEmail(db, email)
		fmt.Println(selected_emp)
		fmt.Println("Data yang ingin diubah")
		fmt.Println("1. Nama\n2. Email\n3. Usia\n4. Alamat")
		fmt.Print("Masukkan pilihan Anda : ")
		fmt.Scanf("%d", &pilihan)
		if pilihan == 1 {
			fmt.Print("Masukkan Nama Baru : ")
			input_mod, _ = reader.ReadString('\n')
			input_mod = strings.TrimSpace(input_mod)
			selected_emp.Name = input_mod
		} else if pilihan == 2 {
			fmt.Print("Masukkan Email Baru : ")
			input_mod, _ = reader.ReadString('\n')
			input_mod = strings.TrimSpace(input_mod)
			selected_emp.Email = input_mod
		} else if pilihan == 3 {
			fmt.Print("Masukkan Usia Baru : ")
			fmt.Scanf("%d", &input_int)
			selected_emp.Age = input_int
		} else if pilihan == 4 {
			fmt.Print("Masukkan Alamat Baru : ")
			input_mod, _ = reader.ReadString('\n')
			input_mod = strings.TrimSpace(input_mod)
			selected_emp.Address = input_mod
		}
		fmt.Println(selected_emp)
		err = updateData(db, selected_emp)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Data berhasil di update")
	} else if pilihan == 4 {
		fmt.Print("Masukkan email pegawai yang ingin diupdate: ")
		email, _ := reader.ReadString('\n')
		email = strings.TrimSpace(email)
		err = deleteData(db, email)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Data berhasil dihapus")
	}

}

func getInputData() (name, email, address string, age int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan Nama : ")
	name, _ = reader.ReadString('\n')
	fmt.Print("Masukkan Email : ")
	email, _ = reader.ReadString('\n')
	fmt.Print("Masukkan Alamat : ")
	address, _ = reader.ReadString('\n')
	fmt.Print("Masukkan Usia : ")
	fmt.Scanf("%d", &age)
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	address = strings.TrimSpace(address)
	return name, email, address, age
}

func insertData(db *sql.DB, emp Employee) error {
	q := "INSERT INTO employee (name, email, address, age) VALUES (?, ?, ?, ?)"
	insert, err := db.Prepare(q)
	defer insert.Close()

	if err != nil {
		panic(err.Error())
	}

	_, err = insert.Exec(emp.Name, emp.Email, emp.Address, emp.Age)

	if err != nil {
		panic(err.Error())
	}
	return nil
}

func getAllData(db *sql.DB) (emp []Employee, err error) {
	resp, err := db.Query("SELECT id, name, email, address, age FROM employee")
	defer resp.Close()
	if err != nil {
		panic(err.Error())
	}
	for resp.Next() {
		var pPerson Employee
		err = resp.Scan(&pPerson.Id, &pPerson.Name, &pPerson.Email, &pPerson.Address, &pPerson.Age)
		if err != nil {
			panic(err.Error())
		}
		emp = append(emp, pPerson)
	}
	return emp, nil
}

func getSelectedEmail(db *sql.DB, email string) (emp Employee, err error) {
	q := fmt.Sprintf(`SELECT id, name, email, address, age FROM employee WHERE email like "%s" LIMIT 1;`, email)
	select_db, err := db.Query(q)
	defer select_db.Close()
	if err != nil {
		panic(err.Error())
	}

	for select_db.Next() {
		err = select_db.Scan(&emp.Id, &emp.Name, &emp.Email, &emp.Address, &emp.Age)
		if err != nil {
			panic(err.Error())
		}
	}
	return emp, nil
}

func deleteData(db *sql.DB, email string) error {
	q := fmt.Sprintf(`DELETE FROM employee WHERE email like "%s";`, email)
	fmt.Println(q)
	drop, err := db.Prepare(q)
	defer drop.Close()
	if err != nil {
		panic(err.Error())
	}

	_, err = drop.Exec()
	if err != nil {
		panic(err.Error())
	}
	return nil
}

func updateData(db *sql.DB, emp Employee) error {
	q := "UPDATE employee SET name = ?, address = ?, age = ?, email = ?  WHERE email like ?"
	update, err := db.Prepare(q)
	defer update.Close()
	if err != nil {
		panic(err.Error())
	}

	_, err = update.Exec(emp.Name, emp.Address, emp.Age, emp.Email, emp.Email)
	if err != nil {
		panic(err.Error())
	}
	return nil
}
