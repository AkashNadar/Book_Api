package config

import (
	"fmt"

	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDSNString() string {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "3308"
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "root"
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "password"
	}

	dbName := os.Getenv("DB_DBNAME")
	if dbName == "" {
		dbName = "dbname"
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName)
}

func Connect() {
	// for getting info from the env
	// dsn := GetDSNString()
	// dsn := "newUser:password@tcp(mysql_db:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "newUser:password@tcp(mysql_db)/dbname"
	// dsn := "user:password@tcp(db)/dbname?charset=utf8"
	fmt.Print(dsn)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connection Successful")
	db = d
}

func GetDB() *gorm.DB {
	return db
}
