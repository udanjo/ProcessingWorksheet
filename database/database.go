package database

import (
	"log"
	"time"

	"github.com/ueverson/ProcessingWorksheetGO/database/migrations"
	"github.com/ueverson/ProcessingWorksheetGO/middleware"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	connString := "sqlserver://sa:udanjo2011@@localhost:5434?database=ProjectSend"

	database, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})

	if err != nil {
		log.Fatal("error database: ", err)
	}

	db = database
	config, err := db.DB()

	middleware.Handler(err)

	config.SetConnMaxIdleTime(10)
	config.SetMaxOpenConns(50)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)

}

func GetDatabase() *gorm.DB {
	return db
}
