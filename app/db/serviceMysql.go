package db

import (
	repositoryMysql "UserApi/app/gormMysql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func Initialize() *gorm.DB {
	dsnPattern := "%s:%s@tcp(%s:%s)/%s?parseTime=true"
	print(os.Getenv("MYSQL_HOST"))
	dsn := fmt.Sprintf(
		dsnPattern,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	print("mysql connect ok\n")
	err = db.AutoMigrate(&repositoryMysql.User{})
	if err != nil {
		panic("Failed user migration: " + err.Error())
	}
	print("AutoMigrate ok\n")

	return db
}
