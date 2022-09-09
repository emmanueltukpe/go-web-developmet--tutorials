package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db * gorm.DB
)

func Connect(){
	d, err := gorm.Open("sqlite3", "./gorm.db")
	 if err != nil {
		panic(err) 
	 }
	 db = d
	 //defer d.Close()
}

func GetDB() *gorm.DB {
	return db
}