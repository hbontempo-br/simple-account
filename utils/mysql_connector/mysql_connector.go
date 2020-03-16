package mysql_connector

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"simple-account/configuration"
)

var db *gorm.DB //database

func init() {

	dbConfig := configuration.GetDBConfig()
	dbUri := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", dbConfig.User, dbConfig.Password, dbConfig.Address, dbConfig.Name) //Build connection string

	conn, err := gorm.Open("mysql", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	log.Print("Connected to database")
	db = conn
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
