package db

import (
	"admin/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DataB *gorm.DB

func GetDbConnection() *gorm.DB {
	_, dbConf ,err := config.GetConfig("db")
	conString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbConf.DbConnection.Host,  dbConf.DbConnection.Port, dbConf.DbConnection.User, dbConf.DbConnection.Password, dbConf.DbConnection.Dbname)
	db, err := gorm.Open(postgres.Open(conString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})					 								//TODO: тут нужно использовать данные с конфига а не статично
	if err != nil {										//Done
		panic("не удалось подключиться к базе данных")
	}
	DataB = db
	return db
}
