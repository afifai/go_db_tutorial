package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	db_user   = "root"
	db_passwd = "password"
	db_addr   = "localhost"
	db_port   = "3306"
	db_db     = "gorm_migration"
)

type Product struct {
	gorm.Model
	NamaProduk string `gorm:"size:255"`
	Harga      int
	Deskripsi  string
	Stok       int
}

type User struct {
	gorm.Model
	Nama  string `gorm:"size:255"`
	Email string `gorm:"unique;not null;index"`
	Bio   string
	Saldo int
}

type spam struct {
	gorm.Model
	UserId   int
	Username string `gorm:"size:100"`
}

func main() {
	s := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db_user, db_passwd, db_addr, db_port, db_db)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	var prod Product
	db.AutoMigrate(&prod, &User{}, &spam{})

	prod = Product{NamaProduk: "roti",
		Harga: 10000}
	db.Create(&prod)

}
