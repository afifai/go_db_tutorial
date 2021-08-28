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
	db_db     = "gorm_tutorial_db"
)

type User struct {
	gorm.Model
	Nama  string
	Email string `gorm:"not null;unique;index"`
}

func main() {
	fmt.Println("Main Program GORM")
	s := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_passwd, db_addr, db_port, db_db)
	fmt.Println(s)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&User{})

	// db.Create(&User{Nama: "Afif", Email: "surel.afifai@gmail.com"})
	nama, email := getInfo()
	u := User{
		Nama:  nama,
		Email: email,
	}

	db.Create(&u)

	// var u []User

	// // db.Model(&u).Where("email=?", "keren@gmail.com").Update("email", "jago@gmail.com")

	// // db.First(&u)
	// // fmt.Println(u)

	// res := db.Find(&u)
	// fmt.Println(res.RowsAffected)
}

func getInfo() (nama, email string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan Nama : ")
	nama, _ = reader.ReadString('\n')
	fmt.Print("Masukkan Email : ")
	email, _ = reader.ReadString('\n')
	nama = strings.TrimSpace(nama)
	email = strings.TrimSpace(email)
	return nama, email
}
