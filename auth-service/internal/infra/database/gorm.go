package database

import (
	"auth_service/internal/entities"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewSQLDB(connectString string, dbName string) *gorm.DB {

	var entities = []interface{}{
		&entities.Site{},
		&entities.User{},
		&entities.UserDetail{},
	}

	db, err := gorm.Open(mysql.Open(connectString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(fmt.Errorf("failed to connect to sql: %v", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("failed to connect to sql: %v", err))

	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := db.AutoMigrate(entities); err != nil {
		panic(fmt.Errorf("failed to auto migrate: %v", err))
	}
	return db
}
