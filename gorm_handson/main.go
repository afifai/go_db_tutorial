package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	db_user   = "root"
	db_passwd = "password"
	db_addr   = "localhost"
	db_port   = "3306"
	db_db     = "gorm_handson"
)

type Employee struct {
	gorm.Model
	Id      int
	Name    string
	Email   string
	Address string
	Age     int
}

func main() {
	s := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_passwd, db_addr, db_port, db_db)
	fmt.Println(s)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&Employee{})

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
		db.Create(&new_emp)
		fmt.Println("Data Berhasil Masuk")
	} else if pilihan == 2 {
		var employees []Employee
		db.Find(&employees)
		for _, e := range employees {
			fmt.Println("Name : ", e.Name, "E-Mail", e.Email, "Address : ", e.Address, "Age : ", e.Age)
		}
	} else if pilihan == 3 {
		var emp Employee
		fmt.Print("Masukkan email pegawai yang ingin diupdate: ")
		email, _ := reader.ReadString('\n')
		email = strings.TrimSpace(email)
		fmt.Println("Data yang ingin diubah")
		fmt.Println("1. Nama\n2. Email\n3. Usia\n4. Alamat")
		fmt.Print("Masukkan pilihan Anda : ")
		fmt.Scanf("%d", &pilihan)
		if pilihan == 1 {
			fmt.Print("Masukkan Nama Baru : ")
			input_mod, _ = reader.ReadString('\n')
			input_mod = strings.TrimSpace(input_mod)
			db.Model(&emp).Where("email=?", email).Update("name", input_mod)
		} else if pilihan == 2 {
			fmt.Print("Masukkan Email Baru : ")
			input_mod, _ = reader.ReadString('\n')
			input_mod = strings.TrimSpace(input_mod)
			db.Model(&emp).Where("email=?", email).Update("email", input_mod)
		} else if pilihan == 3 {
			fmt.Print("Masukkan Usia Baru : ")
			fmt.Scanf("%d", &input_int)
			db.Model(&emp).Where("email=?", email).Update("age", input_int)
		} else if pilihan == 4 {
			fmt.Print("Masukkan Alamat Baru : ")
			input_mod, _ = reader.ReadString('\n')
			input_mod = strings.TrimSpace(input_mod)
			db.Model(&emp).Where("email=?", email).Update("address", input_mod)
		}
		fmt.Println("Data berhasil di update")
	} else if pilihan == 4 {
		var emp Employee
		fmt.Print("Masukkan email pegawai yang ingin diupdate: ")
		email, _ := reader.ReadString('\n')
		email = strings.TrimSpace(email)
		db.Model(&emp).Where("email=?", email).Delete(&emp)
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
