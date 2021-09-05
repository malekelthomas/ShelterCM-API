package postgres

import (
	"github.com/malekelthomas/ShelterCM-API/pkg/store"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn store.GormDBConn

var gdb *gorm.DB

func init() {
	dsn := ""
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbConn.DB = gormDB
	gdb = gormDB
}
