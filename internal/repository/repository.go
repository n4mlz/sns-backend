package repository

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// TODO: 環境変数から読み込むようにする
const (
	user     = "root"
	password = "password"
	host     = "db"
	port     = "3306"
	dbname   = "sns"
)

func NewRepository() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("failed to connect database: %v\n", err)
		return nil, err
	}

	return db, nil
}
