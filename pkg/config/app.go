package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	db *gorm.DB
)

func Connect() {
	errEnv := godotenv.Load("../../.env")
	if errEnv != nil {
		log.Fatalln(errEnv)
		panic("Error loading .env file")
	}
	
	d, err := gorm.Open("mssql", os.Getenv("SQL_CONNECTION_STRING"))
	if err != nil {
		panic(err)
	}

	db = d
	//defer d.Close()
}
func GetDB() *gorm.DB {
	return db
}
