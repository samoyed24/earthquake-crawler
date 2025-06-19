package repo

import (
	"earthquake-crawler/internal/config"
	"earthquake-crawler/internal/model"
	"earthquake-crawler/internal/storage"
	"encoding/json"
	"fmt"
)

func AddJapanEEWRecord(eewData *model.JapanEEWData) error {
	if storage.DB == nil {
		return fmt.Errorf("数据库未初始化")
	}
	insertSQL := "INSERT INTO japan_eew_record (report_time, eew_detail) VALUES (?, ?)"
	jsonBytes, err := json.Marshal(eewData)
	if err != nil {
		return err
	}

	_, err = storage.DB.Exec(insertSQL, eewData.ReportTime, string(jsonBytes))
	if err != nil {
		return err
	}
	return nil
}

func RPushJapanEEWRecord(eewData *model.JapanEEWData) error {
	jsonData, err := json.Marshal(eewData)
	if err != nil {
		return err
	}
	return storage.RPushRedis(config.Cfg.JPEEW.RedisKey, jsonData)
}
