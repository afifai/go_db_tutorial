package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	NamaProduk string
	Harga      int
	Stok       int
}

type User struct {
	gorm.Model
	Nama   string
	Credit int
	Bucket int
}

const (
	db_user   = "root"
	db_passwd = "password"
	db_addr   = "localhost"
	db_port   = "3306"
	db_db     = "gorm_transaction"
)

func main() {
	fmt.Println("Main Program GORM")
	s := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_passwd, db_addr, db_port, db_db)
	fmt.Println(s)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&Product{}, &User{})

	user := User{Nama: "afif", Credit: 1000000, Bucket: 0}
	produk := Product{NamaProduk: "SIMCard", Harga: 50000, Stok: 50}
	// db.Create(&user)
	// db.Create(&produk)
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&user).Where("nama=?", "afif").Update("credit", 700000).Error; err != nil {
		tx.Rollback()
	}
	if err := tx.Model(&produk).Where("nama_produk=?", "SIMCard").Update("stok", 47).Error; err != nil {
		tx.Rollback()
	}
	if err := tx.Model(&user).Where("nama=afif").Update("bucket", 3).Error; err != nil {
		tx.Rollback()
	}
	tx.Commit()
}
