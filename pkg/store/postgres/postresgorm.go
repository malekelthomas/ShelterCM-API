package postgres

import (
	"log"

	"github.com/malekelthomas/ShelterCM-API/pkg/config"
	"github.com/malekelthomas/ShelterCM-API/pkg/store"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn store.GormDBConn

var gdb *gorm.DB

func init() {
	dsn := config.Config.GetDSN()
	log.Printf("Opening db connection: %s\n", dsn)
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbConn.DB = gormDB
	gdb = gormDB
}
