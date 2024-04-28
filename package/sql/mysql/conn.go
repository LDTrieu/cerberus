package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ldtrieu/cerberus/config"
	"github.com/ldtrieu/cerberus/package/gzlog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	maxOpenConns    = 60
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

func InitConnection(cfg *config.MysqlConfig) (*gorm.DB, *sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.MySqlUser, cfg.MySqlPassword, cfg.MySqlHost, cfg.MySqlPort, cfg.MySqlDbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger(cfg.MySqlLogLevel),
	})
	if err != nil {
		return nil, nil, err
	}

	dbConfig, err := db.DB()
	if err != nil {
		return nil, nil, err
	}
	dbConfig.SetMaxOpenConns(maxOpenConns)
	dbConfig.SetConnMaxLifetime(connMaxLifetime * time.Second)
	dbConfig.SetMaxIdleConns(maxIdleConns)
	dbConfig.SetConnMaxIdleTime(connMaxIdleTime * time.Second)
	if err = dbConfig.Ping(); err != nil {
		return nil, nil, err
	}

	return db, dbConfig, nil
}

func newLogger(log string) gzlog.Logger {
	var level logger.LogLevel
	switch log {
	case "error":
		level = logger.Error
	case "warn":
		level = logger.Warn
	default:
		level = logger.Info
	}

	return gzlog.New(level)
}
