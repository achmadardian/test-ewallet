package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type database interface {
	Write() *gorm.DB
	Read() *gorm.DB
}

type dbInstance struct {
	writeDB *gorm.DB
	readDB  *gorm.DB
}

func (d *dbInstance) Write() *gorm.DB {
	return d.writeDB
}

func (d *dbInstance) Read() *gorm.DB {
	return d.readDB
}

func InitDB() database {
	// init master DB
	masterDB, err := initMasterDB()
	if err != nil {
		log.Fatalf("database: failed initialize master database: %v", err)
	}

	return &dbInstance{
		writeDB: masterDB,
		readDB:  masterDB,
	}
}

func initMasterDB() (*gorm.DB, error) {
	dbUser := os.Getenv("DB_MASTER_USER")
	dbPass := os.Getenv("DB_MASTER_PASSWORD")
	dbName := os.Getenv("DB_MASTER_NAME")
	dbHost := os.Getenv("DB_MASTER_HOST")
	dbPort := os.Getenv("DB_MASTER_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName, dbPort)
	masterDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open DB: %w", err)
	}

	sqlDB, err := masterDB.DB()
	if err != nil {
		return nil, fmt.Errorf("get raw DB: %w", err)
	}

	// set conn
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(1 * time.Hour)

	log.Println("database: success initialize database")
	return masterDB, nil
}
