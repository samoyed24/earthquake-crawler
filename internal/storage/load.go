package storage

import (
	"database/sql"
	"earthquake-crawler/internal/config"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func initDB() error {
	err := InitJapanEarthquakeDB()
	if err != nil {
		return err
	}
	err = InitJapanEEWDB()
	if err != nil {
		return err
	}
	return nil
}

func LoadDB() error {
	_db, err := sql.Open("sqlite", config.Cfg.DB.DBPath)
	if err != nil {
		return err
	}
	DB = _db

	if err := initDB(); err != nil {
		return fmt.Errorf("在初始化数据库的过程中发生错误: %v", err)
	}
	return nil
}
