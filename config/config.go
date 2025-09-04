package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

func ConnectDB(dialect, dsn string) (*gorm.DB, error) {
	if dialect == "sqlite" {
		return gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	}
	return nil, nil
}
