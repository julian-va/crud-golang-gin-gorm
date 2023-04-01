package config

import (
	"crud-golang-gin-gorm/helper"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	HOST     = "localHost"
	PORT     = 5432
	USER     = "TEST"
	PASSWORD = "TEST"
	DBNAME   = "TEST"
)

func DataBaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("")
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)
	return db
}
